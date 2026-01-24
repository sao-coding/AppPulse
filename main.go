package main

import (
	"os"
	"strings"

	"apppulse/pkg/app"
	"apppulse/pkg/config"
	"apppulse/pkg/logger"
	"apppulse/pkg/tray"
)

func main() {
	const configPath = "config.yaml"

	// 檢查設定檔案是否存在，如果不存在則建立一個預設的
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := config.CreateDefault(configPath); err != nil {
			// 建立設定檔案失敗時，寫入日誌檔案而非控制台
			// 這裡使用臨時 logger 避免開啟控制台視窗
			tempLogger := logger.New("error", "apppulse_error.log", false)
			tempLogger.Fatalf("建立預設設定檔案失敗: %v", err)
		}
		// 設定檔案建立成功，寫入日誌檔案而非控制台
		tempLogger := logger.New("info", "apppulse.log", false)
		tempLogger.Printf("已建立預設設定檔案: %s", configPath)
		tempLogger.Println("請編輯設定檔案後重新執行程式。")
		return
	}

	// 建立 AppPulse 客戶端
	client, err := app.NewAppPulseClient(configPath)
	if err != nil {
		// 由於 logger 可能尚未初始化，建立臨時 logger 寫入錯誤日誌
		tempLogger := logger.New("error", "apppulse_error.log", false)
		tempLogger.Fatalf("建立 AppPulse 客戶端失敗: %v", err)
	}

	// 初始化系統托盤
	client.InitTray()

	// 設定重啟回調函式
	tray.SetRestartFunc(func() {
		client.Restart()
	})

	// 啟動系統托盤（這會阻塞直到退出）
	tray.Run(func() {
		// 托盤就緒後的回調
		tray.Setup()
		// 在背景執行客戶端（支援重啟循環）
		go runClientLoop(client)
	}, func() {
		// 托盤退出時的回調
		client.Shutdown()
	})
}

// runClientLoop 執行客戶端循環，支援重啟
func runClientLoop(client *app.AppPulseClient) {
	for {
		err := client.Run()
		if err != nil {
			// 檢查是否是重啟請求
			if strings.Contains(err.Error(), "restart") {
				client.GetLogger().Println("正在重新啟動...")
				// 重新建立客戶端
				newClient, err := app.NewAppPulseClient("config.yaml")
				if err != nil {
					client.GetLogger().Printf("重啟失敗: %v", err)
					return
				}
				// 更新客戶端引用
				*client = *newClient
				client.GetLogger().Println("重啟完成，繼續執行...")
				continue
			}
			client.GetLogger().Printf("執行客戶端時發生錯誤: %v", err)
		}
		// 正常退出，結束循環
		break
	}
}
