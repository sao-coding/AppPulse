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

		// 在 Windows 上，我們可以直接嘗試 Kill
		if runtime.GOOS == "windows" {
			c.logger.Println("Windows 平台：嘗試直接終止程序。")
			if err := c.csharpProcess.Kill(); err != nil {
				c.logger.Printf("強制終止 C# 應用程式失敗: %v", err)
			} else {
				c.logger.Println("C# 應用程式已被要求終止。")
			}
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