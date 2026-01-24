package config

import (
	"os"

	"apppulse/pkg/models"

	"sigs.k8s.io/yaml"
)

// Load 函式從指定的路徑載入 YAML 設定檔並解析到 models.Config 結構中。
func Load(configPath string) (*models.Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config models.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// CreateDefault 函式會建立一個預設的設定檔。
func CreateDefault(configPath string) error {
	defaultConfig := `webhook:
  url: ""            # Webhook URL，留空則停用
  auth_key: ""       # 認證金鑰

idle:
  window_seconds: 300 # 視窗閒置時間 (秒)
  music_seconds: 60 # 音樂閒置時間 (秒)

filter:
  apps:              # 要過濾的應用程式清單
    - "explorer"
    - "dwm"
 
  msedge_inprivate: true # 是否過濾 Edge InPrivate 模式 (true: 過濾, false: 不過濾)

logging:
  level: "info"      # 日誌等級: debug, info, warn, error
  file_path: "apppulse.log" # 日誌檔案路徑，留空則輸出到控制台

debug: false         # 除錯模式 (true: 輸出到控制台和日誌檔案, false: 只輸出到日誌檔案)
`
	return os.WriteFile(configPath, []byte(defaultConfig), 0644)
}