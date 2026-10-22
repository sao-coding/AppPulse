package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// 實機整合測試：需要真正的 AppPulse.exe。
// APPPULSE_CSHARP_PATH=D:\...\plugins\AppPulse.exe go test ./pkg/app -run Integration -v
func TestIntegrationSupervisor(t *testing.T) {
	if os.Getenv("APPPULSE_CSHARP_PATH") == "" {
		t.Skip("未設定 APPPULSE_CSHARP_PATH")
	}

	var mu sync.Mutex
	var got []map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		var m map[string]any
		_ = json.Unmarshal(b, &m)
		mu.Lock()
		got = append(got, m)
		mu.Unlock()
	}))
	defer srv.Close()

	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	cfg := fmt.Sprintf("webhook:\n  url: %q\nidle:\n  window_seconds: 600\n  music_seconds: 60\nfilter:\n  msedge_inprivate: true\nlogging:\n  file_path: %q\n",
		srv.URL, filepath.Join(dir, "log.txt"))
	if err := os.WriteFile(cfgPath, []byte(cfg), 0644); err != nil {
		t.Fatal(err)
	}
	c, err := NewAppPulseClient(cfgPath)
	if err != nil {
		t.Fatal(err)
	}

	runDone := make(chan error, 1)
	go func() { runDone <- c.Run() }()

	waitFor := func(what string, d time.Duration, cond func() bool) {
		t.Helper()
		end := time.Now().Add(d)
		for time.Now().Before(end) {
			if cond() {
				return
			}
			time.Sleep(200 * time.Millisecond)
		}
		logb, _ := os.ReadFile(filepath.Join(dir, "log.txt"))
		t.Fatalf("逾時等待 %s\n日誌:\n%s", what, logb)
	}
	count := func() int { mu.Lock(); defer mu.Unlock(); return len(got) }
	csharpPIDs := func() string {
		out, _ := exec.Command("tasklist", "/FI", "IMAGENAME eq AppPulse.exe", "/NH").Output()
		return string(out)
	}

	// 1. 啟動後要收到至少一個視窗事件（C# 啟動時會回報目前前景視窗）
	waitFor("第一個事件", 20*time.Second, func() bool { return count() > 0 })

	// 2. 殺掉 C# 程序 → 要自動重啟並再次回報
	before := count()
	_ = exec.Command("taskkill", "/F", "/IM", "AppPulse.exe").Run()
	waitFor("重啟後的事件", 40*time.Second, func() bool { return count() > before+1 })
	if !strings.Contains(csharpPIDs(), "AppPulse.exe") {
		t.Fatal("C# 沒有被重新啟動")
	}

	// 3. 關閉 → C# 要被優雅結束
	c.Shutdown()
	select {
	case <-runDone:
	case <-time.After(10 * time.Second):
		t.Fatal("Run 沒有在關閉後返回")
	}
	waitFor("C# 結束", 8*time.Second, func() bool { return !strings.Contains(csharpPIDs(), "AppPulse.exe") })
}
