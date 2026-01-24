package tray

import (
	_ "embed"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"fyne.io/systray"
)

//go:embed icon.ico
var iconData []byte

var (
	logger       *log.Logger
	logPath      string
	restartFunc  func()
)

// Initialize 初始化系統托盤
func Initialize(lg *log.Logger, logFilePath string) {
	logger = lg
	logPath = logFilePath
}

// SetRestartFunc 設定重新啟動回調函式
func SetRestartFunc(fn func()) {
	restartFunc = fn
}

// Run 啟動系統托盤
func Run(onReady func(), onExit func()) {
	systray.Run(onReady, onExit)
}

// Setup 設定系統托盤的圖示和選單
func Setup() {
	// 設定托盤圖示
	systray.SetIcon(iconData)
	systray.SetTitle("AppPulse")
	systray.SetTooltip("AppPulse - 應用程式活動監控")

	// 建立選單項目
	mShowLog := systray.AddMenuItem("顯示 Log", "開啟 apppulse.log 檔案")
	mRestart := systray.AddMenuItem("重新啟動", "重新啟動 AppPulse")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出 AppPulse")

	// 處理選單點擊事件
	go func() {
		for {
			select {
			case <-mShowLog.ClickedCh:
				openLogFile()
			case <-mRestart.ClickedCh:
				restart()
			case <-mQuit.ClickedCh:
				logger.Println("使用者從托盤選單請求退出")
				systray.Quit()
				return
			}
		}
	}()
}

// openLogFile 開啟 log 檔案
func openLogFile() {
	if logPath == "" {
		if logger != nil {
			logger.Println("警告: 未設定 log 檔案路徑")
		}
		return
	}

	// 取得絕對路徑
	absPath, err := filepath.Abs(logPath)
	if err != nil {
		if logger != nil {
			logger.Printf("取得 log 檔案絕對路徑失敗: %v", err)
		}
		return
	}

	// 檢查檔案是否存在
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		if logger != nil {
			logger.Printf("Log 檔案不存在: %s", absPath)
		}
		return
	}

	if logger != nil {
		logger.Printf("開啟 log 檔案: %s", absPath)
	}

	// 在 Windows 上使用 notepad 開啟
	cmd := exec.Command("notepad", absPath)
	if err := cmd.Start(); err != nil {
		if logger != nil {
			logger.Printf("開啟 log 檔案失敗: %v", err)
		}
	}
}

// restart 重新啟動應用程式
func restart() {
	if logger != nil {
		logger.Println("使用者從托盤選單請求重新啟動")
	}

	if restartFunc != nil {
		restartFunc()
	} else {
		if logger != nil {
			logger.Println("警告: 重新啟動功能未設定")
		}
	}
}

// Quit 退出系統托盤
func Quit() {
	systray.Quit()
}
