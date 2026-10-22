package tray

import (
	_ "embed"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"fyne.io/systray"
)

//go:embed icon.ico
var iconData []byte

// Callbacks 是托盤選單觸發的動作，由應用程式核心提供。
type Callbacks struct {
	Restart          func()
	Reload           func()
	SetPaused        func(paused bool)
	SetInPrivate     func(enabled bool)
	ConfigPath       string
	InitialInPrivate bool
}

var (
	logger  *log.Logger
	logPath string
	cb      Callbacks

	mu         sync.Mutex
	mStatus    *systray.MenuItem
	mPause     *systray.MenuItem
	mInPrivate *systray.MenuItem
	pendingIP  *bool // 托盤就緒前收到的 InPrivate 狀態
	pendingTxt string
	paused     bool
)

// Initialize 初始化系統托盤
func Initialize(lg *log.Logger, logFilePath string) {
	logger = lg
	logPath = logFilePath
}

// SetCallbacks 設定選單動作。
func SetCallbacks(c Callbacks) {
	cb = c
	v := c.InitialInPrivate
	pendingIP = &v
}

// Run 啟動系統托盤
func Run(onReady func(), onExit func()) {
	systray.Run(onReady, onExit)
}

// SetStatus 更新選單最上方的狀態文字。
func SetStatus(text string) {
	mu.Lock()
	defer mu.Unlock()
	pendingTxt = text
	if mStatus != nil {
		mStatus.SetTitle("狀態：" + text)
		systray.SetTooltip("AppPulse - " + text)
	}
}

// SetInPrivate 同步「過濾 Edge InPrivate」勾選狀態（設定檔重載後呼叫）。
func SetInPrivate(enabled bool) {
	mu.Lock()
	defer mu.Unlock()
	pendingIP = &enabled
	if mInPrivate != nil {
		if enabled {
			mInPrivate.Check()
		} else {
			mInPrivate.Uncheck()
		}
	}
}

// Setup 設定系統托盤的圖示和選單
func Setup() {
	systray.SetIcon(iconData)
	systray.SetTitle("AppPulse")
	systray.SetTooltip("AppPulse - 應用程式活動監控")

	mu.Lock()
	mStatus = systray.AddMenuItem("狀態：啟動中", "目前的追蹤狀態")
	mStatus.Disable()
	if pendingTxt != "" {
		mStatus.SetTitle("狀態：" + pendingTxt)
	}
	systray.AddSeparator()
	mPause = systray.AddMenuItem("暫停追蹤", "暫停回報視窗與音樂")
	checked := pendingIP != nil && *pendingIP
	mInPrivate = systray.AddMenuItemCheckbox("過濾 Edge InPrivate", "不回報 Edge InPrivate 視窗", checked)
	systray.AddSeparator()
	mReload := systray.AddMenuItem("重新載入設定", "重新讀取 config.yaml")
	mOpenConfig := systray.AddMenuItem("開啟設定檔", "用記事本開啟 config.yaml")
	mShowLog := systray.AddMenuItem("顯示 Log", "開啟 apppulse.log 檔案")
	mRestart := systray.AddMenuItem("重新啟動", "重新啟動 AppPulse")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出 AppPulse")
	mu.Unlock()

	go func() {
		for {
			select {
			case <-mPause.ClickedCh:
				paused = !paused
				if paused {
					mPause.SetTitle("恢復追蹤")
				} else {
					mPause.SetTitle("暫停追蹤")
				}
				if cb.SetPaused != nil {
					cb.SetPaused(paused)
				}
			case <-mInPrivate.ClickedCh:
				enable := !mInPrivate.Checked()
				if enable {
					mInPrivate.Check()
				} else {
					mInPrivate.Uncheck()
				}
				if cb.SetInPrivate != nil {
					cb.SetInPrivate(enable)
				}
			case <-mReload.ClickedCh:
				if cb.Reload != nil {
					cb.Reload()
				}
			case <-mOpenConfig.ClickedCh:
				openFile(cb.ConfigPath)
			case <-mShowLog.ClickedCh:
				openFile(logPath)
			case <-mRestart.ClickedCh:
				restart()
			case <-mQuit.ClickedCh:
				logf("使用者從托盤選單請求退出")
				systray.Quit()
				return
			}
		}
	}()
}

func logf(format string, args ...any) {
	if logger != nil {
		logger.Printf(format, args...)
	}
}

// openFile 用記事本開啟檔案
func openFile(path string) {
	if path == "" {
		logf("警告: 未設定檔案路徑")
		return
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		logf("取得檔案絕對路徑失敗: %v", err)
		return
	}
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		logf("檔案不存在: %s", absPath)
		return
	}
	logf("開啟檔案: %s", absPath)
	if err := exec.Command("notepad", absPath).Start(); err != nil {
		logf("開啟檔案失敗: %v", err)
	}
}

// restart 重新啟動應用程式
func restart() {
	logf("使用者從托盤選單請求重新啟動")
	if cb.Restart != nil {
		cb.Restart()
	} else {
		logf("警告: 重新啟動功能未設定")
	}
}

// Quit 退出系統托盤
func Quit() {
	systray.Quit()
}
