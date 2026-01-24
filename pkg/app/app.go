package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"apppulse/pkg/config"
	"apppulse/pkg/logger"
	"apppulse/pkg/models"
	"apppulse/pkg/tray"

	"golang.org/x/sys/windows"
)

// AppPulseClient 是應用程式的核心客戶端結構。
type AppPulseClient struct {
	config        *models.Config
	pipeHandle    windows.Handle
	logger        *log.Logger
	httpClient    *http.Client
	csharpProcess *os.Process
	stopOnce      sync.Once
	shutdownChan  chan struct{}
	restartChan   chan struct{}

	// 閒置處理相關欄位
	mu              sync.Mutex
	lastWindowEvent *models.WindowEvent
	windowIdleTimer *time.Timer
	lastMusicEvent  *models.MusicEvent
	musicIdleTimer  *time.Timer
}

// NewAppPulseClient 建立並返回一個新的 AppPulseClient 實例。
func NewAppPulseClient(configPath string) (*AppPulseClient, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("載入配置失敗: %v", err)
	}

	// 設定日誌，傳入 debug 模式參數
	lg := logger.New(cfg.Logging.Level, cfg.Logging.FilePath, cfg.Debug)

	// 如果開啟了除錯模式，則列印完整的配置
	if cfg.Debug {
		configJSON, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			lg.Printf("無法序列化配置以進行除錯: %v", err)
		} else {
			lg.Printf("DEBUG: 載入的配置:\n%s", string(configJSON))
		}
	}

	// 檢查並設定預設閒置時間
	if cfg.Idle.WindowSeconds == 0 {
		lg.Println("警告: 視窗閒置時間 (window_seconds) 未在 config.yaml 中配置或為0。使用預設值 60 秒。")
		cfg.Idle.WindowSeconds = 60
	}
	if cfg.Idle.MusicSeconds == 0 {
		lg.Println("警告: 音樂閒置時間 (music_seconds) 未在 config.yaml 中配置或為0。使用預設值 300 秒。")
		cfg.Idle.MusicSeconds = 300
	}

	// 建立HTTP客戶端
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &AppPulseClient{
		config:       cfg,
		logger:       lg,
		httpClient:   httpClient,
		shutdownChan: make(chan struct{}),
		restartChan:  make(chan struct{}),
	}, nil
}

// InitTray 初始化系統托盤
func (c *AppPulseClient) InitTray() {
	tray.Initialize(c.logger, c.config.Logging.FilePath)
}

// GetLogger 取得 logger 實例
func (c *AppPulseClient) GetLogger() *log.Logger {
	return c.logger
}

// Shutdown 關閉客戶端
func (c *AppPulseClient) Shutdown() {
	c.logger.Println("正在關閉客戶端...")
	select {
	case <-c.shutdownChan:
		// 已經關閉
	default:
		close(c.shutdownChan)
	}
	c.stopCSharpApp()
	if c.pipeHandle != 0 {
		windows.CloseHandle(c.pipeHandle)
	}
}

// Restart 重新啟動客戶端
func (c *AppPulseClient) Restart() {
	c.logger.Println("正在重新啟動客戶端...")
	select {
	case <-c.restartChan:
		// 已經發送重啟信號
	default:
		close(c.restartChan)
	}
}

// Run 啟動並執行 AppPulse 客戶端。
func (c *AppPulseClient) Run() error {
	c.logger.Println("啟動 AppPulse 客戶端...")

	// 啟動C#應用程式
	if err := c.startCSharpApp(); err != nil {
		return fmt.Errorf("啟動C#應用程式失敗: %v", err)
	}
	defer c.stopCSharpApp()

	// 連接到管道
	if err := c.connectToPipe(); err != nil {
		// 即使連線失敗，也要確保 C# 程序被關閉
		return fmt.Errorf("連線失敗: %v", err)
	}
	defer windows.CloseHandle(c.pipeHandle)

	// 設定訊號處理
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 啟動讀取協程
	go func() {
		if err := c.readFromPipe(); err != nil {
			// 在 readFromPipe 內部已經記錄了錯誤，這裡可能只需要簡單退出或標記錯誤狀態
			c.logger.Printf("讀取管道時發生無法恢復的錯誤: %v", err)
			// 發送一個關閉信號來終止應用程式
			sigChan <- syscall.SIGINT
		}
	}()

	// 等待訊號或關閉通知
	select {
	case <-sigChan:
		c.logger.Println("收到退出訊號，正在關閉...")
	case <-c.shutdownChan:
		c.logger.Println("收到關閉通知，正在關閉...")
	case <-c.restartChan:
		c.logger.Println("收到重啟通知，正在重新啟動...")
		// 停止 C# 應用程式
		c.stopCSharpApp()
		// 關閉管道
		if c.pipeHandle != 0 {
			windows.CloseHandle(c.pipeHandle)
			c.pipeHandle = 0
		}
		// 返回特殊錯誤表示需要重啟
		return fmt.Errorf("restart")
	}

	// 停止C#應用程式 (defer 會再次調用，但 stopOnce 會防止重複執行)
	c.stopCSharpApp()

	return nil
}