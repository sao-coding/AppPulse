package logger

import (
	"io"
	"log"
	"os"
)

// New 函式根據提供的日誌等級和檔案路徑設定並返回一個新的 *log.Logger 實例。
func New(level, filePath string) *log.Logger {
	var output io.Writer = os.Stdout

	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Printf("無法開啟日誌檔案 %s: %v", filePath, err)
		} else {
			// 注意：這裡我們應該在應用程式關閉時關閉檔案，
			// 在當前的簡單設計中，我們暫時忽略它，但在生產級應用中需要處理。
			output = file
		}
	}

	// 這裡可以根據 level 參數進一步擴充，例如設定不同的日誌輸出或格式。
	// 目前，我們保持與原始程式碼相同的行為。

	return log.New(output, "[AppPulse] ", log.LstdFlags|log.Lshortfile)
}