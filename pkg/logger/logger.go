package logger

import (
	"io"
	"log"
	"os"
)

// New 函式根據提供的日誌等級、檔案路徑和除錯模式設定並返回一個新的 *log.Logger 實例。
// 如果 debug 為 true，日誌會同時輸出到檔案和控制台。
// 如果 debug 為 false，日誌只輸出到檔案（如果指定了 filePath）。
func New(level, filePath string, debug bool) *log.Logger {
	var output io.Writer

	// 非除錯模式：只使用檔案輸出
	if !debug {
		if filePath != "" {
			file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				// 非除錯模式下，如果無法開啟日誌檔案，使用 ioutil.Discard 避免開啟控制台
				output = io.Discard
			} else {
				output = file
			}
		} else {
			// 非除錯模式且未指定檔案路徑，丟棄所有日誌
			output = io.Discard
		}
	} else {
		// 除錯模式：同時輸出到檔案和控制台
		if filePath != "" {
			file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				// 如果無法開啟檔案，至少輸出到控制台
				output = os.Stdout
			} else {
				// 同時寫入檔案和控制台
				output = io.MultiWriter(file, os.Stdout)
			}
		} else {
			// 只輸出到控制台
			output = os.Stdout
		}
	}

	return log.New(output, "[AppPulse] ", log.LstdFlags|log.Lshortfile)
}