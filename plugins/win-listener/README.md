# AppPulse - 前台窗口與音樂播放監聽服務

基於 .NET 10.0 構建的控制台應用程序，用於監聽前台窗口變化和音樂播放狀態。

## 功能特性

### 1. 前台窗口監聽

- 使用 Win32 API 捕獲當前活動的應用程式
- 顯示應用程式名稱、進程路徑和窗口標題
- 實時監控窗口切換

### 2. 音樂播放資訊監聽

- 使用 WinRT MediaSession API 獲取音樂播放資訊
- 監聽曲名、演出者、專輯資訊
- 實時檢測播放狀態變化（播放/暫停/停止）

## 系統要求

- Windows 10/11
- .NET 8.0 Runtime
- 支援 WinRT 的 Windows 版本

## 編譯與運行

```bash
# 還原 NuGet 包
dotnet restore

# 編譯項目
dotnet build

# 運行程序
dotnet run

#發布
dotnet publish -c Release -r win-x64 --self-contained true -p:PublishSingleFile=true -p:IncludeNativeLibrariesForSelfExtract=true -o ./dist
```

## 使用說明

1. 運行程序後，會自動開始監聽前台窗口和音樂播放狀態
2. 程序會實時輸出當前活動的應用程式和音樂播放資訊
3. 按 `Ctrl+C` 退出程序

## 輸出範例

```
[14:30:15] 前台窗口: chrome - Google Chrome
[14:30:16] 音樂播放: ▶ 播放中
  曲名: Bohemian Rhapsody
  演出者: Queen
  專輯: A Night at the Opera

[14:30:20] 前台窗口: notepad - 無標題 - 記事本
[14:30:25] 音樂播放: ⏸ 暫停
  曲名: Bohemian Rhapsody
  演出者: Queen
  專輯: A Night at the Opera
```

## 技術實現

- **前台窗口監聽**: 使用 `user32.dll` Win32 API
- **音樂播放監聽**: 使用 Windows Runtime (WinRT) MediaSession API
- **異步處理**: 使用 async/await 模式
- **資源管理**: 實現 IDisposable 接口確保資源正確釋放
