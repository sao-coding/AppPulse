package app

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"apppulse/pkg/models"
)

func newTestClient(inprivate bool, apps ...string) *AppPulseClient {
	cfg := &models.Config{}
	cfg.Filter.MsedgeInprivate = inprivate
	cfg.Filter.Apps = apps
	return &AppPulseClient{config: cfg, logger: log.New(io.Discard, "", 0)}
}

func TestIsEventFiltered(t *testing.T) {
	cases := []struct {
		name      string
		inprivate bool
		apps      []string
		process   string
		title     string
		want      bool
	}{
		{"InPrivate 英文標題", true, nil, "msedge", "New tab - [InPrivate] - Microsoft​ Edge", true},
		{"InPrivate 大小寫不同", true, nil, "msedge", "新索引標籤 - [inprivate] - Microsoft​ Edge", true},
		{"process 帶 .exe 與大寫", true, nil, "MSEdge.exe", "x - InPrivate", true},
		{"一般 Edge 視窗", true, nil, "msedge", "GitHub - Microsoft​ Edge", false},
		{"關閉過濾時放行", false, nil, "msedge", "New tab - [InPrivate]", false},
		{"非 Edge 不受影響", true, nil, "chrome", "[InPrivate]", false},
		{"應用程式清單", false, []string{"explorer", "dwm"}, "explorer", "x", true},
		{"清單不分大小寫與 .exe", false, []string{"Explorer.exe"}, "explorer", "x", true},
		{"空白項目不誤殺", false, []string{""}, "", "x", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(tc.inprivate, tc.apps...)
			got := c.isEventFiltered(&models.WindowEvent{Process: tc.process, Title: tc.title})
			if got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMaskTitle(t *testing.T) {
	c := newTestClient(false)
	c.config.Filter.MaskTitleApps = []string{"KeePass.exe"}
	c.config.Filter.MaskTitleText = "（已隱藏）"
	c.config.Idle.WindowSeconds = 300

	ev := &models.WindowEvent{Process: "keepass", App: "KeePass", Title: "secret.kdbx"}
	if err := c.handleWindowEvent(ev); err != nil {
		t.Fatal(err)
	}
	if c.curWindow.Title != "（已隱藏）" {
		t.Fatalf("title not masked: %q", c.curWindow.Title)
	}

	ev2 := &models.WindowEvent{Process: "code", App: "VS Code", Title: "main.go"}
	_ = c.handleWindowEvent(ev2)
	if c.curWindow.Title != "main.go" {
		t.Fatalf("unrelated app masked: %q", c.curWindow.Title)
	}
}

func TestInputIdleAndResume(t *testing.T) {
	c := newTestClient(false)
	c.config.Idle.WindowSeconds = 300
	_ = c.handleWindowEvent(&models.WindowEvent{Process: "code", App: "VS Code", Title: "a"})
	if c.windowIdle {
		t.Fatal("should start active")
	}
	_ = c.handleInputEvent(&models.InputEvent{IdleSeconds: 10})
	if c.windowIdle {
		t.Fatal("recent input must stay active")
	}
	_ = c.handleInputEvent(&models.InputEvent{IdleSeconds: 301})
	if !c.windowIdle {
		t.Fatal("no input for 301s must be idle")
	}
	_ = c.handleInputEvent(&models.InputEvent{IdleSeconds: 2})
	if c.windowIdle {
		t.Fatal("input after idle must resume active")
	}
}

func TestLockAndPause(t *testing.T) {
	c := newTestClient(false)
	c.config.Idle.WindowSeconds = 300
	_ = c.handleWindowEvent(&models.WindowEvent{Process: "code", App: "VS Code", Title: "a"})

	_ = c.handleSessionEvent(&models.SessionEvent{State: "locked"})
	if !c.windowIdle || !c.locked {
		t.Fatal("lock must mark idle")
	}
	// 鎖定期間的視窗事件不得讓狀態回到 active
	_ = c.handleWindowEvent(&models.WindowEvent{Process: "LockApp", App: "x", Title: "lock"})
	if !c.windowIdle {
		t.Fatal("window event while locked must not reactivate")
	}
	// 沒有輸入回報時，解鎖直接恢復
	_ = c.handleSessionEvent(&models.SessionEvent{State: "unlocked"})
	if c.windowIdle || c.locked {
		t.Fatal("unlock without input flow must resume")
	}

	c.SetPaused(true)
	if !c.windowIdle {
		t.Fatal("pause must mark idle")
	}
	_ = c.handleInputEvent(&models.InputEvent{IdleSeconds: 0})
	if !c.windowIdle {
		t.Fatal("input while paused must not resume")
	}
	c.SetPaused(false)
	_ = c.handleInputEvent(&models.InputEvent{IdleSeconds: 0})
	if c.windowIdle {
		t.Fatal("resume + input must be active")
	}
}

func TestReloadAndInPrivateToggle(t *testing.T) {
	p := filepath.Join(t.TempDir(), "c.yaml")
	body := "idle:\n  window_seconds: 10\nfilter:\n  msedge_inprivate: true\n"
	if err := os.WriteFile(p, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
	c, err := NewAppPulseClient(p)
	if err != nil {
		t.Fatal(err)
	}
	if !c.cfg().Filter.MsedgeInprivate {
		t.Fatal("initial true")
	}
	c.SetInPrivateFilter(false)
	if c.cfg().Filter.MsedgeInprivate {
		t.Fatal("toggle must apply")
	}
	data, _ := os.ReadFile(p)
	if !strings.Contains(string(data), "msedge_inprivate: false") {
		t.Fatalf("file not updated: %s", data)
	}
	// 壞掉的設定檔不得覆蓋舊設定
	_ = os.WriteFile(p, []byte("filter: [oops"), 0644)
	if err := c.ReloadConfig(); err == nil {
		t.Fatal("expected error")
	}
	if c.cfg().Idle.WindowSeconds != 10 {
		t.Fatal("old config must be kept")
	}
}
