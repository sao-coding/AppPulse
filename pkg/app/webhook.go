package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"apppulse/pkg/models"
)

// sendToWebhook 負責將事件序列化為 JSON 並發送到設定中指定的 Webhook URL。
func (c *AppPulseClient) sendToWebhook(event models.Event) error {
	if c.config.Webhook.URL == "" {
		return nil // 如果沒有設定 Webhook URL，則直接返回
	}

	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("序列化事件失敗: %v", err)
	}

	req, err := http.NewRequest("POST", c.config.Webhook.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("建立 HTTP 請求失敗: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.config.Webhook.AuthKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.Webhook.AuthKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("傳送 Webhook 請求失敗: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Webhook 請求失敗，狀態碼: %d, 回應: %s", resp.StatusCode, string(body))
	}

	c.logger.Printf("成功傳送事件到 Webhook: %s", event.GetType())
	return nil
}