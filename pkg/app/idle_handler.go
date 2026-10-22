package app

import (
	"time"

	"apppulse/pkg/models"
	"apppulse/pkg/tray"
)

// inputFlowWindow 內有收到 C# 的輸入閒置回報，視窗閒置就完全以輸入為準；否則退回計時器。
const inputFlowWindow = 20 * time.Second

func (c *AppPulseClient) suppressedLocked() bool { return c.paused || c.locked }

func (c *AppPulseClient) inputFlowingLocked() bool {
	return !c.lastInputAt.IsZero() && time.Since(c.lastInputAt) < inputFlowWindow
}

func (c *AppPulseClient) stopTimersLocked() {
	if c.windowIdleTimer != nil {
		c.windowIdleTimer.Stop()
		c.windowIdleTimer = nil
	}
	if c.musicIdleTimer != nil {
		c.musicIdleTimer.Stop()
		c.musicIdleTimer = nil
	}
}

func (c *AppPulseClient) armWindowTimerLocked() {
	if c.windowIdleTimer != nil {
		c.windowIdleTimer.Stop()
	}
	c.windowIdleTimer = time.AfterFunc(time.Duration(c.cfg().Idle.WindowSeconds)*time.Second, c.handleWindowIdle)
}

// resetWindowIdleTimer 記錄最新的視窗事件並重置視窗閒置計時器。
func (c *AppPulseClient) resetWindowIdleTimer(event *models.WindowEvent) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.curWindow = event
	c.windowIdle = false
	c.armWindowTimerLocked()
	c.refreshTrayLocked()
}

// handleWindowIdle 在視窗閒置超時後觸發（沒有輸入回報時的後備機制）。
func (c *AppPulseClient) handleWindowIdle() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.suppressedLocked() || c.inputFlowingLocked() {
		return
	}
	c.sendWindowIdleLocked("視窗閒置逾時")
}

// sendWindowIdleLocked 對外標記目前視窗為閒置（已閒置則不重複傳送）。
func (c *AppPulseClient) sendWindowIdleLocked(reason string) {
	if c.curWindow == nil || c.windowIdle {
		return
	}
	c.logger.Printf("%s，傳送閒置事件", reason)
	c.windowIdle = true
	c.refreshTrayLocked()

	idle := *c.curWindow
	idle.Time = time.Now().Format(time.RFC3339)
	idle.Status = "idle"
	if err := c.sendToWebhook(&idle); err != nil {
		c.logger.Printf("傳送視窗閒置事件失敗: %v", err)
	}
}

// sendWindowActiveLocked 讓目前視窗從閒置恢復為活動。
func (c *AppPulseClient) sendWindowActiveLocked(reason string) {
	if c.curWindow == nil || !c.windowIdle {
		return
	}
	c.logger.Printf("%s，恢復為活動狀態", reason)
	c.windowIdle = false
	c.armWindowTimerLocked()
	c.refreshTrayLocked()

	active := *c.curWindow
	active.Time = time.Now().Format("15:04:05")
	active.Status = "active"
	if err := c.sendToWebhook(&active); err != nil {
		c.logger.Printf("傳送視窗活動事件失敗: %v", err)
	}
}

// handleInputEvent 處理 C# 回報的「距離上次鍵鼠輸入的秒數」。
func (c *AppPulseClient) handleInputEvent(e *models.InputEvent) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.lastInputAt = time.Now()
	if c.suppressedLocked() || c.curWindow == nil {
		return nil
	}
	if e.IdleSeconds >= c.cfg().Idle.WindowSeconds {
		c.sendWindowIdleLocked("使用者已無輸入")
	} else {
		c.sendWindowActiveLocked("偵測到使用者輸入")
	}
	return nil
}

// handleSessionEvent 處理鎖定/解鎖與睡眠/喚醒。
func (c *AppPulseClient) handleSessionEvent(e *models.SessionEvent) error {
	c.logger.Printf("工作階段狀態: %s", e.State)
	switch e.State {
	case "locked", "suspend":
		c.setSuppressed(func() { c.locked = true })
	case "unlocked", "resume":
		c.setSuppressed(func() { c.locked = false })
	}
	return nil
}

// SetPaused 暫停或恢復追蹤（托盤選單使用）。
func (c *AppPulseClient) SetPaused(paused bool) {
	if paused {
		c.logger.Println("追蹤已暫停")
	} else {
		c.logger.Println("追蹤已恢復")
	}
	c.setSuppressed(func() { c.paused = paused })
}

// setSuppressed 改變暫停/鎖定旗標，並依結果切換視窗狀態。
func (c *AppPulseClient) setSuppressed(change func()) {
	c.mu.Lock()
	defer c.mu.Unlock()

	was := c.suppressedLocked()
	change()
	now := c.suppressedLocked()

	switch {
	case !was && now:
		c.stopTimersLocked()
		c.sendWindowIdleLocked("進入暫停/鎖定")
	case was && !now:
		// 有輸入回報時，下一次回報（約 5 秒內）會自動恢復；否則直接恢復。
		if !c.inputFlowingLocked() {
			c.sendWindowActiveLocked("結束暫停/鎖定")
		}
	}
	c.refreshTrayLocked()
}

// markDisconnected 在 C# 監聽程式中斷時，把目前視窗標記為閒置，避免網站一直顯示過期的活動。
func (c *AppPulseClient) markDisconnected() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.lastInputAt = time.Time{}
	c.sendWindowIdleLocked("監聽程式中斷")
}

// refreshTrayLocked 更新托盤上的狀態文字。
func (c *AppPulseClient) refreshTrayLocked() {
	switch {
	case c.paused:
		tray.SetStatus("已暫停")
	case c.locked:
		tray.SetStatus("已鎖定")
	case c.curWindow == nil:
		tray.SetStatus("等待事件")
	case c.windowIdle:
		tray.SetStatus("閒置")
	default:
		tray.SetStatus("追蹤中：" + c.curWindow.Process)
	}
}

// newMusicIdleTimer 建立一個新的音樂閒置計時器。
func (c *AppPulseClient) newMusicIdleTimer() *time.Timer {
	return time.AfterFunc(time.Duration(c.cfg().Idle.MusicSeconds)*time.Second, c.handleMusicIdle)
}

// handleMusicIdle 在音樂閒置超時後觸發。
func (c *AppPulseClient) handleMusicIdle() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lastMusicEvent == nil {
		return
	}

	c.logger.Println("音樂閒置逾時，傳送閒置事件")

	idleEvent := &models.MusicEvent{
		Time:   time.Now().Format(time.RFC3339),
		Type:   "music",
		Status: "idle",
		Title:  c.lastMusicEvent.Title,
		Artist: c.lastMusicEvent.Artist,
	}

	if err := c.sendToWebhook(idleEvent); err != nil {
		c.logger.Printf("傳送音樂閒置事件失敗: %v", err)
	}

	// 清空最後事件，避免重複傳送
	c.lastMusicEvent = nil
}
