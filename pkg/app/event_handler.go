package app

import (
	"encoding/json"
	"fmt"

	"apppulse/pkg/models"
	"strings"
)

// processEvent 解析從管道接收到的原始位元組數據，並根據事件類型進行分發。
func (c *AppPulseClient) processEvent(data []byte) error {
	// 移除可能的 UTF-8 BOM
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		data = data[3:]
	}

	var eventData map[string]interface{}
	if err := json.Unmarshal(data, &eventData); err != nil {
		c.logger.Printf("無法解析JSON資料: %s", string(data))
		return err
	}

	eventType, ok := eventData["type"].(string)
	if !ok {
		c.logger.Printf("事件缺少 'type' 欄位: %s", string(data))
		return fmt.Errorf("事件缺少 'type' 欄位")
	}

	switch eventType {
	case "music":
		var musicEvent models.MusicEvent
		if err := json.Unmarshal(data, &musicEvent); err != nil {
			c.logger.Printf("解析音樂事件失敗: %v", err)
			return err
		}
		return c.handleMusicEvent(&musicEvent)
	case "window":
		var windowEvent models.WindowEvent
		if err := json.Unmarshal(data, &windowEvent); err != nil {
			c.logger.Printf("解析視窗事件失敗: %v", err)
			return err
		}
		return c.handleWindowEvent(&windowEvent)
	default:
		c.logger.Printf("未知事件類型: %s", eventType)
		return nil
	}
}

// handleMusicEvent 處理音樂相關的事件。
func (c *AppPulseClient) handleMusicEvent(event *models.MusicEvent) error {
	c.logger.Printf("音樂事件: %s - %s (%s) - %s",
		event.Status, event.Title, event.Artist, event.Time)

	c.mu.Lock()
	defer c.mu.Unlock()

	// 如果正在播放，取消閒置計時器
	if event.Status == "playing" {
		if c.musicIdleTimer != nil {
			c.musicIdleTimer.Stop()
			c.musicIdleTimer = nil
			c.logger.Println("音樂開始播放，取消閒置計時器")
		}
	} else if event.Status == "paused" {
		// 如果是暫停，則開始計算閒置時間
		c.lastMusicEvent = event
		if c.musicIdleTimer != nil {
			c.musicIdleTimer.Stop()
		}
		c.logger.Printf("音樂已暫停，將在 %d 秒後標記為閒置", c.config.Idle.MusicSeconds)
		c.musicIdleTimer = c.newMusicIdleTimer()
	}

	// 傳送到Webhook
	return c.sendToWebhook(event)
}

// handleWindowEvent 處理視窗活動相關的事件。
func (c *AppPulseClient) handleWindowEvent(event *models.WindowEvent) error {
	// C# IPC通道不會傳送status欄位，在此處手動設定為active
	event.Status = "active"

	// 檢查事件是否應該被過濾
	if c.isEventFiltered(event) {
		return nil
	}

	// 如果是 Edge 瀏覽器，移除標題中的 " - Microsoft Edge"
	if event.Process == "msedge" {
		event.Title = strings.TrimSuffix(event.Title, " - Microsoft\u200b Edge")
	}

	c.logger.Printf("視窗事件: %s - %s (%s) - %s",
		event.Process, event.App, event.Title, event.Time)

	// 重置視窗閒置計時器
	c.resetWindowIdleTimer(event)

	// 傳送到Webhook
	return c.sendToWebhook(event)
}

// isEventFiltered 檢查一個視窗事件是否應該被過濾。
func (c *AppPulseClient) isEventFiltered(event *models.WindowEvent) bool {
	// 檢查應用程式名稱是否在過濾清單中
	for _, filteredApp := range c.config.Filter.Apps {
		if event.Process == filteredApp {
			c.logger.Printf("應用程式 %s 在過濾清單中，已過濾", event.Process)
			return true
		}
	}

	// 檢查 Edge InPrivate 模式
	if c.config.Filter.MsedgeInprivate &&
		event.Process == "msedge" &&
		strings.Contains(event.Title, "[InPrivate]") {
		c.logger.Printf("過濾 Edge InPrivate 視窗: %s", event.Title)
		return true
	}

	return false
}