package app

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/sys/windows"
)

const pipeName = `\\.\pipe\AppPulsePipe`

// connectToPipe 負責連接到由 C# 應用程式建立的命名管道。
func (c *AppPulseClient) connectToPipe() error {
	c.logger.Printf("正在嘗試連接到命名管道: %s", pipeName)
	var handle windows.Handle
	var err error

	// 增加重試機制，因為 C# 應用程式可能需要一些時間來建立管道
	for i := 0; i < 10; i++ {
		handle, err = windows.CreateFile(
			windows.StringToUTF16Ptr(pipeName),
			windows.GENERIC_READ,
			windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
			nil,
			windows.OPEN_EXISTING,
			windows.FILE_ATTRIBUTE_NORMAL,
			0,
		)
		if err == nil {
			c.pipeHandle = handle
			c.logger.Println("成功連接到命名管道")
			return nil
		}
		c.logger.Printf("連接管道失敗 (嘗試 %d/10): %v. 1秒後重試...", i+1, err)
		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("在多次嘗試後，仍無法連接到管道 %s: %v", pipeName, err)
}

// readFromPipe 持續從命名管道讀取資料並進行處理。
func (c *AppPulseClient) readFromPipe() error {
	reader := bufio.NewReader(os.NewFile(uintptr(c.pipeHandle), "pipe"))

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// 如果是 EOF，代表管道已關閉，可能是 C# 程式崩潰或關閉
			if err == io.EOF {
				c.logger.Println("管道連線已斷開 (EOF)。應用程式即將關閉。")
				return err // 返回錯誤，觸發 Run() 中的關閉流程
			}
			c.logger.Printf("讀取管道資料時發生錯誤: %v", err)
			return err // 對於其他錯誤也返回，以中斷循環
		}

		line = string(bytes.TrimSpace([]byte(line)))
		if len(line) == 0 {
			continue
		}

		// 處理接收到的資料
		if err := c.processEvent([]byte(line)); err != nil {
			c.logger.Printf("處理事件失敗: %v", err)
			// 即使處理失敗，我們也繼續讀取下一個事件
		}
	}
}