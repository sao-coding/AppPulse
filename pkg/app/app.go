package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"sync"
	"syscall"
	"time"

	"apppulse/pkg/config"
	"apppulse/pkg/logger"
	"apppulse/pkg/models"
	"apppulse/pkg/tray"
)

const defaultMaskTitle = "（已隱藏）"

var (
	errShutdown = errors.New("shutdown")
	errRestart  = errors.New("restart")
)

// AppPulseClient 是應用程式的核心客戶端結構。
type AppPulseClient struct {
	configPath string
	logger     *log.Logger
	httpClient *http.Client

	// 設定可在執行中被熱重載，一律透過 cfg() 讀取。
	cfgMu  sync.RWMutex
	config *models.Config

	ctlMu        sync.Mutex // 保護 shutdownChan / restartChan
	shutdownChan chan struct{}
	restartChan  chan struct{}
	watchOnce    sync.Once

	// 閒置、暫停與工作階段狀態，皆由 mu 保護
	mu              sync.Mutex
	paused          bool
	locked          bool
	curWindow       *models.WindowEvent // 最近一次實際的視窗事件（閒置後仍保留，供恢復使用）
	windowIdle      bool                // 目前是否已對外標記為閒置
	windowIdleTimer *time.Timer
	lastInputAt     time.Time // 最近一次收到輸入閒置回報的時間
	lastMusicEvent  *models.MusicEvent
	musicIdleTimer  *time.Timer
}

// applyDefaults 補上設定的預設值。
func applyDefaults(cfg *models.Config, lg *log.Logger) {
	if cfg.Idle.WindowSeconds == 0 {
		lg.Println("警告: 視窗閒置時間 (window_seconds) 未在 config.yaml 中配置或為0。使用預設值 60 秒。")
		cfg.Idle.WindowSeconds = 60
	}
	if cfg.Idle.MusicSeconds == 0 {
		lg.Println("警告: 音樂閒置時間 (music_seconds) 未在 config.yaml 中配置或為0。使用預設值 300 秒。")
		cfg.Idle.MusicSeconds = 300
	}
	if cfg.Filter.MaskTitleText == "" {
		cfg.Filter.MaskTitleText = defaultMaskTitle
	}
}

// NewAppPulseClient 建立並返回一個新的 AppPulseClient 實例。
func NewAppPulseClient(configPath string) (*AppPulseClient, error) {
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("載入配置失敗: %v", err)
	}

	// 設定日誌，傳入 debug 模式參數
	lg := logger.New(cfg.Logging.Level, cfg.Logging.FilePath, cfg.Debug)

	for _, w := range cfg.Warnings {
		lg.Printf("警告: %s", w)
	}
	lg.Printf("Edge InPrivate 過濾: %v，過濾應用程式: %v，隱藏標題的應用程式: %v",
		cfg.Filter.MsedgeInprivate, cfg.Filter.Apps, cfg.Filter.MaskTitleApps)

	// 如果開啟了除錯模式，則列印完整的配置
	if cfg.Debug {
		configJSON, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			lg.Printf("無法序列化配置以進行除錯: %v", err)
		} else {
			lg.Printf("DEBUG: 載入的配置:\n%s", string(configJSON))
		}
	}

	applyDefaults(cfg, lg)

	return &AppPulseClient{
		configPath:   configPath,
		config:       cfg,
		logger:       lg,
		httpClient:   &http.Client{Timeout: 30 * time.Second},
		shutdownChan: make(chan struct{}),
		restartChan:  make(chan struct{}),
	}, nil
}

// cfg 回傳目前生效的設定（熱重載後會被整份替換，不會原地修改）。
func (c *AppPulseClient) cfg() *models.Config {
	c.cfgMu.RLock()
	defer c.cfgMu.RUnlock()
	return c.config
}

// ReloadConfig 重新讀取設定檔並套用；失敗時保留舊設定。
func (c *AppPulseClient) ReloadConfig() error {
	cfg, err := config.Load(c.configPath)
	if err != nil {
		c.logger.Printf("重新載入設定失敗，沿用舊設定: %v", err)
		return err
	}
	for _, w := range cfg.Warnings {
		c.logger.Printf("警告: %s", w)
	}
	applyDefaults(cfg, c.logger)

	c.cfgMu.Lock()
	old := c.config
	c.config = cfg
	c.cfgMu.Unlock()

	if old.Logging.Level != cfg.Logging.Level || old.Logging.FilePath != cfg.Logging.FilePath || old.Debug != cfg.Debug {
		c.logger.Println("提醒: 日誌與除錯設定的變更需要重新啟動才會生效")
	}
	c.logger.Printf("已重新載入設定。Edge InPrivate 過濾: %v，過濾應用程式: %v，隱藏標題的應用程式: %v",
		cfg.Filter.MsedgeInprivate, cfg.Filter.Apps, cfg.Filter.MaskTitleApps)
	tray.SetInPrivate(cfg.Filter.MsedgeInprivate)
	return nil
}

var inprivateLine = regexp.MustCompile(`(?m)^(\s*msedge_inprivate:\s*)(true|false)`)

// SetInPrivateFilter 修改設定檔中的 msedge_inprivate 並立即套用（設定檔是唯一的真實來源）。
func (c *AppPulseClient) SetInPrivateFilter(enabled bool) {
	data, err := os.ReadFile(c.configPath)
	if err != nil {
		c.logger.Printf("讀取設定檔失敗: %v", err)
		tray.SetInPrivate(c.cfg().Filter.MsedgeInprivate)
		return
	}
	if !inprivateLine.Match(data) {
		c.logger.Println("設定檔中找不到 msedge_inprivate，無法從托盤切換，請手動新增")
		tray.SetInPrivate(c.cfg().Filter.MsedgeInprivate)
		return
	}
	repl := "${1}false"
	if enabled {
		repl = "${1}true"
	}
	out := inprivateLine.ReplaceAll(data, []byte(repl))
	if err := os.WriteFile(c.configPath, out, 0644); err != nil {
		c.logger.Printf("寫入設定檔失敗: %v", err)
		tray.SetInPrivate(c.cfg().Filter.MsedgeInprivate)
		return
	}
	_ = c.ReloadConfig()
}

// watchConfig 輪詢設定檔，內容有變就熱重載。
func (c *AppPulseClient) watchConfig() {
	var lastMod time.Time
	var lastSize int64 = -1
	if st, err := os.Stat(c.configPath); err == nil {
		lastMod, lastSize = st.ModTime(), st.Size()
	}
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-c.shutdownChan:
			return
		case <-ticker.C:
		}
		st, err := os.Stat(c.configPath)
		if err != nil || (st.ModTime().Equal(lastMod) && st.Size() == lastSize) {
			continue
		}
		lastMod, lastSize = st.ModTime(), st.Size()
		// 編輯器儲存時可能分多次寫入，稍等再讀
		time.Sleep(300 * time.Millisecond)
		c.logger.Println("偵測到設定檔變更，重新載入...")
		_ = c.ReloadConfig()
	}
}

// InitTray 初始化系統托盤
func (c *AppPulseClient) InitTray() {
	tray.Initialize(c.logger, c.cfg().Logging.FilePath)
	tray.SetCallbacks(tray.Callbacks{
		Restart:          c.Restart,
		Reload:           func() { _ = c.ReloadConfig() },
		SetPaused:        c.SetPaused,
		SetInPrivate:     c.SetInPrivateFilter,
		ConfigPath:       c.configPath,
		InitialInPrivate: c.cfg().Filter.MsedgeInprivate,
	})
}

// GetLogger 取得 logger 實例
func (c *AppPulseClient) GetLogger() *log.Logger {
	return c.logger
}

// Shutdown 關閉客戶端；實際的資源清理由 Run 內的 runSession 完成。
func (c *AppPulseClient) Shutdown() {
	c.logger.Println("正在關閉客戶端...")
	c.ctlMu.Lock()
	select {
	case <-c.shutdownChan:
		c.ctlMu.Unlock()
		c.logger.Println("客戶端已經處於關閉狀態")
		return
	default:
		close(c.shutdownChan)
	}
	c.ctlMu.Unlock()

	c.mu.Lock()
	c.stopTimersLocked()
	c.mu.Unlock()

	// 給 Run 一點時間停掉 C# 程序
	time.Sleep(200 * time.Millisecond)
	tray.Quit()
	c.logger.Println("客戶端關閉完成")
}

// Restart 要求重新載入設定並重啟 C# 監聽程式。
func (c *AppPulseClient) Restart() {
	c.logger.Println("正在重新啟動客戶端...")
	c.ctlMu.Lock()
	defer c.ctlMu.Unlock()
	select {
	case <-c.restartChan:
	default:
		close(c.restartChan)
	}
}

// Run 啟動並監督 C# 監聽程式；程序或管道中斷時自動重啟，直到收到關閉通知。
func (c *AppPulseClient) Run() error {
	c.logger.Println("啟動 AppPulse 客戶端...")
	c.watchOnce.Do(func() { go c.watchConfig() })

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigChan)

	const maxBackoff = 30 * time.Second
	backoff := time.Second

	for {
		started := time.Now()
		err := c.runSession(sigChan)

		switch {
		case errors.Is(err, errShutdown):
			return nil
		case errors.Is(err, errCSharpMissing):
			return err
		case errors.Is(err, errRestart):
			c.logger.Println("收到重啟通知，重新載入設定並重新啟動監聽程式...")
			_ = c.ReloadConfig()
			c.ctlMu.Lock()
			c.restartChan = make(chan struct{})
			c.ctlMu.Unlock()
			backoff = time.Second
			continue
		}

		// 穩定執行過一段時間後才算恢復正常，退避時間歸零
		if time.Since(started) > maxBackoff {
			backoff = time.Second
		}
		c.logger.Printf("監聽程式中斷 (%v)，%v 後自動重新啟動", err, backoff)
		c.markDisconnected()

		select {
		case <-time.After(backoff):
		case <-c.shutdownChan:
			return nil
		case <-sigChan:
			c.Shutdown()
			return nil
		}
		if backoff *= 2; backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
}

// runSession 完成一輪「啟動 C# → 連管道 → 讀取事件」，並保證結束時 C# 與管道都已清理。
func (c *AppPulseClient) runSession(sigChan <-chan os.Signal) error {
	p, err := c.startCSharpApp()
	if err != nil {
		return err
	}
	var pipe *os.File
	defer func() {
		c.stopCSharpApp(p) // 先停子程序，讀取端才會因 EOF 返回
		if pipe != nil {
			go pipe.Close()
		}
	}()

	done := make(chan struct{})
	defer close(done)
	pipe, err = c.connectToPipe(p, c.anyControlSignal(done))
	if err != nil {
		return err
	}

	readErr := make(chan error, 1)
	go func() { readErr <- c.readFromPipe(pipe) }()

	c.ctlMu.Lock()
	restartCh := c.restartChan
	c.ctlMu.Unlock()

	select {
	case err := <-readErr:
		return err
	case <-p.done:
		return fmt.Errorf("C# 程序意外結束")
	case <-c.shutdownChan:
		return errShutdown
	case <-restartCh:
		return errRestart
	case <-sigChan:
		c.logger.Println("收到退出訊號，正在關閉...")
		c.Shutdown()
		return errShutdown
	}
}

// anyControlSignal 回傳一個在關閉/重啟/結束時會被關閉的 channel。
func (c *AppPulseClient) anyControlSignal(done <-chan struct{}) <-chan struct{} {
	out := make(chan struct{})
	c.ctlMu.Lock()
	restartCh := c.restartChan
	c.ctlMu.Unlock()
	go func() {
		defer close(out)
		select {
		case <-c.shutdownChan:
		case <-restartCh:
		case <-done:
		}
	}()
	return out
}
