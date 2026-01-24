# AppPulse

一個用 Go 撰寫的程式，用來與 C# AppPulse 應用程式進行 IPC（命名管道）通訊，處理媒體播放與前景視窗事件，並提供將事件轉發到 Webhook 的功能。

## 功能特色

- **IPC 通訊**：透過 Windows 命名管道與 C# 應用程式溝通
- **事件處理**：支援媒體（音樂）播放事件與前景視窗事件
- **Webhook 整合**：自動把事件送到設定的 Webhook URL
- **設定管理**：使用 YAML 格式的 `config.yaml`
- **應用程式過濾**：可設定要過濾的應用程式清單
- **日誌記錄**：支援輸出到檔案或控制台，並支援多種日誌等級
- **穩定性**：自動重連與錯誤復原機制

## 專案概述

本專案在 Windows 上監聽前景視窗與媒體播放資訊，主要包含兩個部分：

- 一個以 Go 寫成的守護程式（主程式），透過 Windows 命名管道與原生外掛程式溝通，處理事件並轉發到 Webhook。
- 一個以 .NET 實作的原生外掛（plugins/win-listener），負責擷取系統的前景視窗與媒體播放狀態，並透過命名管道送給 Go 程式。

此 README 已整合主程式與原生外掛的說明、安裝與執行步驟，以及範例輸出與技術細節。

## 主要功能

- IPC（Windows 命名管道）：Go 與 .NET 原生外掛之間的雙向通訊
- 事件處理：支援媒體播放事件與前景視窗事件
- Webhook：以 JSON 格式將事件送到設定的 Webhook（可設定 Bearer token）
- 設定管理：使用 YAML 的 `config.yaml`
- 應用過濾：可設定要忽略的應用程式名稱（大小寫不敏感）
- 日誌：支援檔案或控制台輸出，並支援不同日誌等級
- 穩定性：自動重連、錯誤記錄，並支援優雅關閉

## 事件範例

媒體播放事件 (type = "music") 範例：

```json
{
  "time": "20:59:11",
  "type": "music",
  "status": "playing",
  "title": "水平線",
  "artist": "back number"
}
```

前景視窗事件 (type = "window") 範例：

```json
{
  "time": "21:24:07",
  "type": "window",
  "process": "explorer",
  "app": "Microsoft® Windows® Operating System",
  "title": "無標題"
}
```

## 安裝與編譯

系統需求：

- Windows 10/11
- Go 1.21+
- .NET 10（若要編譯/執行 `plugins/win-listener`）

1. 下載或更新相依套件（Go）：

```powershell
go mod tidy
```

2. 編譯 Go 主程式：

```powershell
go build -o apppulse.exe main.go
```

3.（選用）編譯原生外掛（.NET）：

在 `plugins/win-listener` 資料夾執行：

```powershell
cd plugins/win-listener; `
dotnet restore; dotnet build; dotnet publish -c Release -r win-x64 --self-contained true -p:PublishSingleFile=true -p:IncludeNativeLibrariesForSelfExtract=true -o ./dist
```

編譯後的可執行檔會放在 `plugins/win-listener/dist`（或 `bin/`）目錄。你也可以直接使用專案中提供的 `AppPulse.exe`（若有）。

## 執行

1. 確保原生監聽器（`plugins/win-listener` 的可執行檔）在系統上執行，或確保系統上有可與命名管道通訊的程式（例如 C# 版 AppPulse）。
2. 編輯 `config.yaml`（若不存在，啟動時會建立預設檔）。
3. 啟動 Go 主程式：

```powershell
.\apppulse.exe
```

或在開發模式下：

```powershell
go run main.go
```

## 範例 `config.yaml`

```yaml
webhook:
  url: 'https://your-webhook-url.com/api/events' # Webhook URL，留空則停用
  auth_key: 'your-auth-key' # Bearer token（可選）

filter:
  apps: # 要過濾的應用程式（大小寫不敏感）
    - 'explorer'
    - 'dwm'
    - 'System'

logging:
  level: 'info' # 日誌等級: debug, info, warn, error
  file_path: 'apppulse.log' # 日誌檔案路徑，留空則輸出到控制台

debug: false # 除錯模式
```

## 原生外掛（plugins/win-listener）說明

此專案包含一個以 .NET 撰寫的輕量監聽程式，負責：

- 使用 Win32 API 取得目前前景視窗（process、app 名稱、視窗標題）
- 使用 WinRT MediaSession API 取得目前媒體播放資訊（曲名、演出者、播放狀態）
- 將擷取到的事件透過 Windows 命名管道傳送給 Go 主程式

技術細節：

- 前景視窗監聽：呼叫 user32.dll（Win32）取得活動視窗與對應 process
- 媒體監聽：使用 Windows Runtime (WinRT) 的 MediaSession API
- 非同步處理：採用 async/await 模式以降低阻塞
- 資源管理：實作 IDisposable 確保正確釋放系統資源

編譯與執行（簡要）：

```powershell
cd plugins/win-listener
# 還原套件
dotnet restore
# 編譯
dotnet build
# 直接執行
dotnet run
```

或發布 Release 單一檔案：

```powershell
dotnet publish -c Release -r win-x64 --self-contained true -p:PublishSingleFile=true -o ./dist
```

範例輸出

```
[14:30:15] 前景視窗: chrome - Google Chrome
[14:30:16] 媒體播放: ▶ 播放中
  曲名: Bohemian Rhapsody
  演出者: Queen
  專輯: A Night at the Opera

[14:30:20] 前景視窗: notepad - 無標題 - 記事本
[14:30:25] 媒體播放: ⏸ 暫停
  曲名: Bohemian Rhapsody
  演出者: Queen
  專輯: A Night at the Opera
```

## 錯誤處理與穩定性

- 自動重連：當命名管道或原生外掛中斷時，Go 程式會嘗試自動重連
- 優雅關閉：支援 Ctrl+C 結束並釋放資源
- 錯誤日誌：所有錯誤會記錄到日誌，依設定輸出到檔案或控制台

## 相依套件

- Go: `golang.org/x/sys`（Windows 系統呼叫）
- Go: `sigs.k8s.io/yaml`（YAML 解析）
- .NET: Windows Runtime（MediaSession API）

## 注意事項

- 目前僅支援 Windows 平台
- 若使用 webhook，請確保系統可以連線到目標且防火牆允許連線
- 若你要在不同電腦上執行原生外掛與 Go 主程式，請改用支援網路的通訊方式（目前為命名管道，僅限本機）

## 貢獻與開發

- 若想擴充事件型別或支援其他平台，可先閱讀 `REFACTORING_PLAN.md` 中的計畫
- 歡迎提交 Pull Request 或 Issue

---

本 README 已整合原生外掛的說明與使用方式。如果你需要，我可以接著：

- 移除或保留 `plugins/win-listener/README.md`（目前建議保留以便提供獨立說明）
- 新增更詳盡的開發教學或測試腳本

若要我繼續（例如：新增測試、更新 config 範例或加入繁簡轉換），請告訴我下一步。
