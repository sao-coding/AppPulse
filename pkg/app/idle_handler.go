package app

import (
	"time"

	"apppulse-easy/pkg/models"
)

// resetWindowIdleTimer 重置視窗閒置計時器。
func (c *AppPulseClient) resetWindowIdleTimer(event *models.WindowEvent) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.lastWindowEvent = event

	if c.windowIdleTimer != nil {
		c.windowIdleTimer.Stop()
	}

	c.logger.Printf("重置視窗閒置計時器，將在 %d 秒後標記為閒置", c.config.Idle.WindowSeconds)
	c.windowIdleTimer = time.AfterFunc(time.Duration(c.config.Idle.WindowSeconds)*time.Second, c.handleWindowIdle)
}

// handleWindowIdle 在視窗閒置超時後觸發。
func (c *AppPulseClient) handleWindowIdle() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.lastWindowEvent == nil {
		return
	}

	c.logger.Println("視窗閒置逾時，傳送閒置事件")

	idleEvent := &models.WindowEvent{
		Time:    time.Now().Format(time.RFC3339),
		Type:    "window",
		Status:  "idle",
		Process: c.lastWindowEvent.Process,
		App:     c.lastWindowEvent.App,
		Title:   c.lastWindowEvent.Title,
	}

	if err := c.sendToWebhook(idleEvent); err != nil {
		c.logger.Printf("傳送視窗閒置事件失敗: %v", err)
	}

	// 清空最後事件，避免重複傳送
	c.lastWindowEvent = nil
}

// newMusicIdleTimer 建立一個新的音樂閒置計時器。
func (c *AppPulseClient) newMusicIdleTimer() *time.Timer {
	return time.AfterFunc(time.Duration(c.config.Idle.MusicSeconds)*time.Second, c.handleMusicIdle)
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