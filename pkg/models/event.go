package models

// Config 定義了應用程式的所有配置選項。
type Config struct {
	Webhook struct {
		URL     string `yaml:"url" json:"url"`
		AuthKey string `yaml:"auth_key" json:"auth_key"`
	} `yaml:"webhook" json:"webhook"`
	Idle struct {
		WindowSeconds int `yaml:"window_seconds" json:"window_seconds"`
		MusicSeconds  int `yaml:"music_seconds" json:"music_seconds"`
	} `yaml:"idle" json:"idle"`
	Filter struct {
		Apps            []string `yaml:"apps" json:"apps"`
		MsedgeInprivate bool     `yaml:"msedge_inprivate" json:"msedge_inprivate"`
	} `yaml:"filter" json:"filter"`
	Logging struct {
		Level    string `yaml:"level" json:"level"`
		FilePath string `yaml:"file_path" json:"file_path"`
	} `yaml:"logging" json:"logging"`
	Debug bool `yaml:"debug" json:"debug"` // 除錯模式：true 會輸出到控制台，false 只寫日誌檔案
}

// MusicEvent 代表一個音樂播放事件。
type MusicEvent struct {
	Time   string `json:"time"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// WindowEvent 代表一個視窗活動事件。
type WindowEvent struct {
	Time    string `json:"time"`
	Type    string `json:"type"`
	Status  string `json:"status,omitempty"`
	Process string `json:"process"`
	App     string `json:"app"`
	Title   string `json:"title"`
}

// Event 是一個通用事件介面，所有事件類型都應實現它。
type Event interface {
	GetType() string
	GetTime() string
}

// GetType 實現了 MusicEvent 的 Event 介面。
func (m MusicEvent) GetType() string { return m.Type }

// GetTime 實現了 MusicEvent 的 Event 介面。
func (m MusicEvent) GetTime() string { return m.Time }

// GetType 實現了 WindowEvent 的 Event 介面。
func (w WindowEvent) GetType() string { return w.Type }

// GetTime 實現了 WindowEvent 的 Event 介面。
func (w WindowEvent) GetTime() string { return w.Time }