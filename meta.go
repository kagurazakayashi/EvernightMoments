package main

import (
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// GetPhotoTime 返回：时间, 来源描述, 错误
func GetPhotoTime(filePath string) (time.Time, string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return time.Time{}, "", err
	}
	defer f.Close()

	// 1. 嘗試 EXIF
	x, err := exif.Decode(f)
	if err == nil {
		tm, err := x.DateTime()
		if err == nil {
			return tm, "EXIF", nil
		}
	}

	// 2. 回退到修改日期
	info, err := f.Stat()
	if err != nil {
		return time.Time{}, "", err
	}
	return info.ModTime(), i18n.T("修改日期"), nil
}
