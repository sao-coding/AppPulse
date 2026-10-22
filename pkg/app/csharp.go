package app

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"
)

// errCSharpMissing 表示找不到 C# 監聽程式，重試也沒有意義。
var errCSharpMissing = errors.New("C# 應用程式不存在")

// csharpProc 代表一個執行中的 C# 監聽子程序。
type csharpProc struct {
	cmd   *exec.Cmd
	stdin io.WriteCloser
	job   windows.Handle // 設有 KILL_ON_JOB_CLOSE：Go 端無論如何結束，子程序都會被帶走
	done  chan struct{}  // 子程序結束後關閉
}

// lineLogger 把子程序輸出逐行寫進日誌。
type lineLogger struct {
	c      *AppPulseClient
	prefix string
	buf    []byte
}

func (w *lineLogger) Write(p []byte) (int, error) {
	w.buf = append(w.buf, p...)
	for {
		i := bytes.IndexByte(w.buf, '\n')
		if i < 0 {
			break
		}
		if line := string(bytes.TrimSpace(w.buf[:i])); line != "" {
			w.c.logger.Printf("%s%s", w.prefix, line)
		}
		w.buf = w.buf[i+1:]
	}
	return len(p), nil
}

// startCSharpApp 負責啟動 C# 子應用程式。
func (c *AppPulseClient) startCSharpApp() (*csharpProc, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("取得可執行檔案路徑失敗: %v", err)
	}
	csharpExePath := filepath.Join(filepath.Dir(exePath), "plugins", "AppPulse.exe")
	if override := os.Getenv("APPPULSE_CSHARP_PATH"); override != "" {
		csharpExePath = override // 測試用
	}

	if _, err := os.Stat(csharpExePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%w: %s", errCSharpMissing, csharpExePath)
	}

	c.logger.Printf("啟動 C# 應用程式: %s", csharpExePath)

	cmd := exec.Command(csharpExePath)
	// 錯誤輸出一律寫進日誌，方便診斷子程序為何結束；一般輸出只在除錯模式記錄。
	cmd.Stderr = &lineLogger{c: c, prefix: "[C#] "}
	if c.cfg().Debug {
		cmd.Stdout = &lineLogger{c: c, prefix: "[C#] "}
	} else {
		cmd.Stdout = io.Discard
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}

	// stdin 是優雅關閉的通道：關閉它，C# 端讀到 EOF 就會自行結束。
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("建立 C# stdin 失敗: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("啟動 C# 應用程式失敗: %v", err)
	}

	p := &csharpProc{cmd: cmd, stdin: stdin, done: make(chan struct{})}
	p.job = c.attachKillOnCloseJob(cmd.Process.Pid)
	go func() {
		err := cmd.Wait()
		if err != nil {
			c.logger.Printf("C# 應用程式已結束: %v", err)
		} else {
			c.logger.Println("C# 應用程式已結束")
		}
		close(p.done)
	}()

	c.logger.Println("C# 應用程式已啟動")
	return p, nil
}

// attachKillOnCloseJob 把子程序放進 Job Object，Job 關閉（含 Go 程式崩潰）時一併終止子程序。
func (c *AppPulseClient) attachKillOnCloseJob(pid int) windows.Handle {
	job, err := windows.CreateJobObject(nil, nil)
	if err != nil {
		c.logger.Printf("建立 Job Object 失敗（略過）: %v", err)
		return 0
	}
	info := windows.JOBOBJECT_EXTENDED_LIMIT_INFORMATION{}
	info.BasicLimitInformation.LimitFlags = windows.JOB_OBJECT_LIMIT_KILL_ON_JOB_CLOSE
	if _, err := windows.SetInformationJobObject(job, windows.JobObjectExtendedLimitInformation,
		uintptr(unsafe.Pointer(&info)), uint32(unsafe.Sizeof(info))); err != nil {
		c.logger.Printf("設定 Job Object 失敗（略過）: %v", err)
		windows.CloseHandle(job)
		return 0
	}
	proc, err := windows.OpenProcess(windows.PROCESS_SET_QUOTA|windows.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		c.logger.Printf("開啟子程序失敗（略過）: %v", err)
		windows.CloseHandle(job)
		return 0
	}
	defer windows.CloseHandle(proc)
	if err := windows.AssignProcessToJobObject(job, proc); err != nil {
		c.logger.Printf("加入 Job Object 失敗（略過）: %v", err)
		windows.CloseHandle(job)
		return 0
	}
	return job
}

// stopCSharpApp 先請 C# 優雅結束（關閉 stdin），逾時才強制終止。
func (c *AppPulseClient) stopCSharpApp(p *csharpProc) {
	if p == nil {
		return
	}
	select {
	case <-p.done:
	default:
		c.logger.Println("正在停止 C# 應用程式...")
		p.stdin.Close()
		select {
		case <-p.done:
			c.logger.Println("C# 應用程式已優雅結束")
		case <-time.After(3 * time.Second):
			c.logger.Println("C# 應用程式未在 3 秒內結束，強制終止...")
			if p.job != 0 {
				windows.CloseHandle(p.job) // KILL_ON_JOB_CLOSE
				p.job = 0
			} else {
				_ = p.cmd.Process.Kill()
			}
			select {
			case <-p.done:
			case <-time.After(3 * time.Second):
				c.logger.Println("C# 應用程式仍未結束，放棄等待")
			}
		}
	}
	if p.job != 0 {
		windows.CloseHandle(p.job)
		p.job = 0
	}
}
