package main

import (
	"os"

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

	// 初始化系統托盤（選單動作在此綁定到客戶端）
	client.InitTray()

	// 啟動系統托盤（這會阻塞直到退出）
	tray.Run(func() {
		tray.Setup()
		// 在背景監督 C# 監聽程式（斷線會自動重啟，重啟與熱重載都在 Run 內處理）
		go func() {
			if err := client.Run(); err != nil {
				client.GetLogger().Printf("執行客戶端時發生無法恢復的錯誤: %v", err)
				tray.Quit()
			}
		}()
	}, func() {
		client.GetLogger().Println("系統托盤正在退出，開始清理資源...")
		client.Shutdown()
		client.GetLogger().Println("清理完成，應用程式即將關閉")
	})
}
