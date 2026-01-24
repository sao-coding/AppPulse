# 系統托盤圖示說明

## 當前圖示

目前使用的是一個簡單的藍色背景白色圓點圖示（16x16 像素）。

## 如何自訂圖示

### 方式 1: 替換現有圖示檔案

1. 準備一個 ICO 格式的圖示檔案（建議大小：16x16 或 32x32 像素）
2. 將檔案命名為 `icon.ico`
3. 替換 `pkg/tray/icon.ico` 檔案
4. 重新編譯程式

### 方式 2: 使用線上工具建立圖示

推薦的線上 ICO 產生器：
- https://www.favicon-generator.org/
- https://convertio.co/zh/png-ico/
- https://www.icoconverter.com/

步驟：
1. 設計或選擇一個圖片（PNG、JPG 等格式）
2. 使用線上工具轉換為 ICO 格式
3. 下載 ICO 檔案
4. 替換 `pkg/tray/icon.ico`
5. 重新編譯

### 方式 3: 使用 Photoshop 或 GIMP

1. 在 Photoshop 或 GIMP 中設計圖示
2. 調整尺寸為 16x16 或 32x32 像素
3. 匯出為 ICO 格式
4. 替換 `pkg/tray/icon.ico`
5. 重新編譯

## 圖示設計建議

- **尺寸**: 建議使用 16x16 或 32x32 像素，Windows 系統托盤通常使用這些尺寸
- **格式**: 必須使用 ICO 格式（Windows 圖示格式）
- **顏色**: 使用對比度高的顏色，確保在淺色和深色背景下都清晰可見
- **簡潔**: 圖示應該簡單明瞭，因為顯示空間很小

## 重新編譯

替換圖示後，執行打包腳本重新編譯：

```cmd
go build -ldflags="-H=windowsgui"
```

## 注意事項

- 圖示檔案會被內嵌到編譯後的執行檔中（使用 Go 的 `//go:embed` 功能）
- 修改圖示後必須重新編譯才會生效
- ICO 檔案不應該太大，建議小於 10KB
