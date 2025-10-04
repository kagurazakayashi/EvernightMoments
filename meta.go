package main

import (
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// GetPhotoTime 嘗試獲取照片的日期資訊。
// 優先從 EXIF 中繼資料提取拍攝時間，若失敗則回傳檔案系統的最後修改時間。
//
// 參數:
//
//	filePath: 檔案的完整路徑
//
// 回傳值:
//
//	time.Time: 獲取到的時間物件
//	string: 時間來源描述（例如 "EXIF" 或 "修改日期"）
//	error: 執行過程中的錯誤資訊
func GetPhotoTime(filePath string) (time.Time, string, error) {
	// 以唯讀模式開啟檔案
	f, err := os.Open(filePath)
	if err != nil {
		return time.Time{}, "", err
	}
	// 確保函式執行完畢後，檔案資源會被正確關閉
	defer f.Close()

	// 1. 嘗試讀取 EXIF 中繼資料
	// 許多數位相機與手機會在照片中寫入 EXIF 資訊，包含拍攝時的詳細參數
	x, err := exif.Decode(f)
	if err == nil {
		// 嘗試從 EXIF 中尋找拍攝日期與時間
		tm, err := x.DateTime()
		if err == nil {
			return tm, "EXIF", nil
		}
	}

	// 2. 回退到檔案系統的修改日期
	// 若檔案不包含 EXIF 資訊（例如經過通訊軟體壓縮）或讀取失敗，則讀取檔案狀態
	info, err := f.Stat()
	if err != nil {
		// 若連檔案狀態都無法讀取，則回傳錯誤
		return time.Time{}, "", err
	}

	// 回傳檔案的最後修改時間，並透過 i18n 轉換來源描述文字
	return info.ModTime(), i18n.T("修改日期"), nil
}
