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
// 子程序提前結束或收到關閉/重啟通知時會立即放棄。
func (c *AppPulseClient) connectToPipe(p *csharpProc, abort <-chan struct{}) (*os.File, error) {
	c.logger.Printf("正在嘗試連接到命名管道: %s", pipeName)
	var err error

	// 增加重試機制，因為 C# 應用程式可能需要一些時間來建立管道
	for i := 0; i < 15; i++ {
		var handle windows.Handle
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
			c.logger.Println("成功連接到命名管道")
			return os.NewFile(uintptr(handle), "pipe"), nil
		}
		select {
		case <-p.done:
			return nil, fmt.Errorf("連線管道前 C# 程序已結束")
		case <-abort:
			return nil, fmt.Errorf("連線管道時收到中止通知")
		case <-time.After(500 * time.Millisecond):
		}
	}

	return nil, fmt.Errorf("在多次嘗試後，仍無法連接到管道 %s: %v", pipeName, err)
}

// readFromPipe 持續從命名管道讀取資料並進行處理，直到管道關閉或發生錯誤。
func (c *AppPulseClient) readFromPipe(f *os.File) error {
	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				c.logger.Println("管道連線已斷開 (EOF)")
				return err
			}
			c.logger.Printf("讀取管道資料時發生錯誤: %v", err)
			return err
		}

		line = string(bytes.TrimSpace([]byte(line)))
		if len(line) == 0 {
			continue
		}

		if err := c.processEvent([]byte(line)); err != nil {
			c.logger.Printf("處理事件失敗: %v", err)
			// 即使處理失敗，我們也繼續讀取下一個事件
		}
	}
}
