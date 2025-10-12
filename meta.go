package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// ExiftoolPath 儲存已偵測到的 Exiftool 執行檔路徑，用於全域快取以避免重複搜尋。
var ExiftoolPath = ""

// GetExiftoolPath 嘗試從系統環境變數中尋找 Exiftool 執行檔。
// 會依序搜尋 "exiftool" 與 Windows 常見的 "exiftool(-k)"，並回傳找到的完整路徑。
func GetExiftoolPath() string {
	// 如果已經有快取路徑，則直接回傳
	if ExiftoolPath != "" {
		return ExiftoolPath
	}
	
	// 定義可能的執行檔名稱
	commands := []string{"exiftool", "exiftool(-k)"}
	for _, cmd := range commands {
		// 在系統 PATH 中搜尋執行檔
		path, err := exec.LookPath(cmd)
		if err == nil {
			ExiftoolPath = path
			return path
		}
	}
	
	// 若皆未找到則確保變數為空
	ExiftoolPath = ""
	return ""
}

// GetExiftoolPathI18n 回傳目前使用的 EXIF 擷取器描述字串。
// 包含擷取器類型與其路徑，並透過 i18n 進行本地化處理。
func GetExiftoolPathI18n() string {
	var path string = i18n.T("内置")
	if ExiftoolPath != "" {
		path = ExiftoolPath
	}
	// 組合輸出字串，例如 "EXIF 獲取器: /usr/bin/exiftool"
	return fmt.Sprintf("EXIF %s: %s", i18n.T("获取器"), path)
}

// GetPhotoTime 嘗試從照片檔案中取得拍攝日期。
// 優先權順序：
// 1. Exiftool 外部指令（支援度最廣）
// 2. goexif 函式庫（原生 Go 實作）
// 3. 檔案系統最後修改時間（最終備案）
func GetPhotoTime(filePath string) (time.Time, string, error) {
	// --- 第一階段：優先嘗試使用 Exiftool ---
	if ExiftoolPath != "" {
		// 執行外部指令獲取原始拍攝日期
		// -s3: 簡短模式，僅輸出標籤值而不含標籤名稱與空格
		// -DateTimeOriginal: 指定讀取原始拍攝時間標籤
		out, err := exec.Command(ExiftoolPath, "-s3", "-DateTimeOriginal", filePath).Output()
		if err == nil {
			dateStr := strings.TrimSpace(string(out))
			if dateStr != "" {
				// Exiftool 預設的日期格式通常為 "YYYY:MM:DD HH:MM:SS"
				// 注意：日期部分是以冒號 ":" 分隔
				tm, err := time.Parse("2006:01:02 15:04:05", dateStr)
				if err == nil {
					return tm, "ExifTool", nil
				}
			}
		}
	}

	// --- 第二階段：回退至原生 goexif 函式庫 ---
	// 以唯讀模式開啟檔案
	f, err := os.Open(filePath)
	if err != nil {
		return time.Time{}, "", err
	}
	// 確保函式結束時關閉檔案資源
	defer f.Close()

	// 嘗試解碼 EXIF 中繼資料
	x, err := exif.Decode(f)
	if err == nil {
		// 呼叫 DateTime 輔助函式解析時間標籤
		tm, err := x.DateTime()
		if err == nil {
			return tm, "Exif", nil
		}
	}

	// --- 第三階段：最終回退至檔案系統修改日期 ---
	// 當檔案不含 EXIF 資訊或讀取失敗時，讀取檔案本身的狀態
	info, err := f.Stat()
	if err != nil {
		return time.Time{}, "", err
	}
	
	// 回傳檔案最後修改時間，並標註來源文字
	return info.ModTime(), i18n.T("修改日期"), nil
}
