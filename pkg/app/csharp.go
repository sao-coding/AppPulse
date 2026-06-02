package app

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
)

// startCSharpApp 負責啟動 C# 子應用程式。
func (c *AppPulseClient) startCSharpApp() error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("取得可執行檔案路徑失敗: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	csharpExePath := filepath.Join(exeDir, "plugins", "AppPulse.exe")

	if _, err := os.Stat(csharpExePath); os.IsNotExist(err) {
		return fmt.Errorf("C# 應用程式不存在: %s", csharpExePath)
	}

	c.logger.Printf("啟動 C# 應用程式: %s", csharpExePath)

	cmd := exec.Command(csharpExePath)
	
	// 根據 debug 模式決定子程序的輸出目標
	if c.config.Debug {
		// Debug 模式：輸出到控制台
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		// 非 Debug 模式：丟棄輸出，避免開啟控制台視窗
		cmd.Stdout = io.Discard
		cmd.Stderr = io.Discard
	}

	// Windows 特定：隱藏子程序視窗
	if runtime.GOOS == "windows" && !c.config.Debug {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow: true,
		}
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("啟動 C# 應用程式失敗: %v", err)
	}

	c.csharpProcess = cmd.Process
	c.logger.Println("C# 應用程式已啟動")

	// 等待一段時間讓 C# 應用程式完全啟動
	time.Sleep(2 * time.Second)

	return nil
}

// stopCSharpApp 負責停止 C# 子應用程式。
func (c *AppPulseClient) stopCSharpApp() {
	c.stopOnce.Do(func() {
		if c.csharpProcess == nil {
			return
		}

		c.logger.Println("正在停止 C# 應用程式...")

		// 在 Windows 上，使用優雅關閉 + 重試機制
		if runtime.GOOS == "windows" {
			c.stopCSharpAppWindows()
			return
		}

		// 對於非 Windows 系統的優雅關閉邏輯
		c.logger.Println("非 Windows 平台：傳送 SIGTERM 訊號...")
		if err := c.csharpProcess.Signal(syscall.SIGTERM); err != nil {
			c.logger.Printf("傳送 SIGTERM 訊號失敗: %v，嘗試強制終止...", err)
			_ = c.csharpProcess.Kill()
			return
		}

		// 等待程序結束
		done := make(chan error, 1)
		go func() {
			_, err := c.csharpProcess.Wait()
			done <- err
		}()

		select {
		case <-time.After(5 * time.Second):
			c.logger.Println("C# 應用程式未在 5 秒內回應 SIGTERM，強制終止...")
			if err := c.csharpProcess.Kill(); err != nil {
				c.logger.Printf("強制終止 C# 應用程式失敗: %v", err)
			}
		case err := <-done:
			if err != nil {
				c.logger.Printf("C# 應用程式退出時出錯: %v", err)
			} else {
				c.logger.Println("C# 應用程式已正常退出。")
			}
		}
	})
}

// stopCSharpAppWindows Windows 平台專用的 C# 應用程式停止邏輯，包含重試機制
func (c *AppPulseClient) stopCSharpAppWindows() {
	const maxRetries = 3
	const retryDelay = 2 * time.Second

	for attempt := 1; attempt <= maxRetries; attempt++ {
		c.logger.Printf("Windows 平台：嘗試終止 C# 程序（第 %d/%d 次）", attempt, maxRetries)

		// 嘗試終止程序
		if err := c.csharpProcess.Kill(); err != nil {
			c.logger.Printf("終止程序失敗: %v", err)
			
			// 如果是最後一次嘗試，記錄錯誤並退出
			if attempt == maxRetries {
				c.logger.Printf("已達到最大重試次數（%d次），放棄終止 C# 應用程式", maxRetries)
				return
			}
			
			// 等待後重試
			c.logger.Printf("等待 %v 後重試...", retryDelay)
			time.Sleep(retryDelay)
			continue
		}

		// 終止成功，等待程序退出
		c.logger.Println("已發送終止訊號，等待程序退出...")
		done := make(chan error, 1)
		go func() {
			_, err := c.csharpProcess.Wait()
			done <- err
		}()

		select {
		case <-time.After(3 * time.Second):
			c.logger.Println("程序未在 3 秒內退出")
			if attempt < maxRetries {
				c.logger.Println("準備重試...")
				time.Sleep(retryDelay)
				continue
			}
		case err := <-done:
			if err != nil {
				c.logger.Printf("C# 應用程式退出時出錯: %v", err)
			} else {
				c.logger.Println("C# 應用程式已成功終止")
			}
			return
		}
	}

	c.logger.Println("C# 應用程式終止流程完成")
}