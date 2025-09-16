package main

import (
	"fmt"
	"log"
	"os"

	"apppulse-easy/pkg/app"
	"apppulse-easy/pkg/config"
)

func main() {
	const configPath = "config.yaml"

	// 檢查設定檔案是否存在，如果不存在則建立一個預設的
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if err := config.CreateDefault(configPath); err != nil {
			log.Fatalf("建立預設設定檔案失敗: %v", err)
		}
		fmt.Printf("已建立預設設定檔案: %s\n", configPath)
		fmt.Println("請編輯設定檔案後重新執行程式。")
		return
	}

	// 建立 AppPulse 客戶端
	client, err := app.NewAppPulseClient(configPath)
	if err != nil {
		// 由於 logger 可能尚未初始化，我們在這裡使用標準的 log.Fatalf
		log.Fatalf("建立 AppPulse 客戶端失敗: %v", err)
	}

	// 執行客戶端
	if err := client.Run(); err != nil {
		// 這裡 client 應該已經有 logger 了，但為了以防萬一，
		// 且為了與上面的錯誤處理保持一致，我們仍可使用 log
		log.Fatalf("執行客戶端時發生錯誤: %v", err)
	}
}
