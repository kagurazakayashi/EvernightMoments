package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

// GenerateNewName 根據格式字串生成新檔名
// format: 使用者定義的格式
// t: 照片時間
// originalPath: 原始檔案路徑
// index: 當前處理的檔案序號（用於 ## 和 #）
func GenerateNewName(format string, t time.Time, originalPath string, index int) string {
	originalName := filepath.Base(originalPath)
	ext := filepath.Ext(originalName)
	nameWithoutExt := strings.TrimSuffix(originalName, ext)

	// 定義替換對映
	replacements := []struct {
		old string
		new string
	}{
		{"YYYY", fmt.Sprintf("%04d", t.Year())},
		{"YY", fmt.Sprintf("%02d", t.Year()%100)},
		{"MM", fmt.Sprintf("%02d", t.Month())},
		{"M", fmt.Sprintf("%d", t.Month())},
		{"DD", fmt.Sprintf("%02d", t.Day())},
		{"D", fmt.Sprintf("%d", t.Day())},
		{"HH", fmt.Sprintf("%02d", t.Hour())},
		{"H", fmt.Sprintf("%d", t.Hour())},
		{"mm", fmt.Sprintf("%02d", t.Minute())},
		{"m", fmt.Sprintf("%d", t.Minute())},
		{"ss", fmt.Sprintf("%02d", t.Second())},
		{"s", fmt.Sprintf("%d", t.Second())},
		{"##", fmt.Sprintf("%02d", index)},
		{"#", fmt.Sprintf("%d", index)},
		{"*", nameWithoutExt},
	}

	result := format

	// 執行替換
	for _, r := range replacements {
		// 替換可能會有重疊問題（例如 MM 和 M），
		// 但按照上面的順序（先長後短），YYYY 會先於 YY 被替換，這是安全的。
		result = strings.ReplaceAll(result, r.old, r.new)
	}

	// 拼接副檔名
	return result + ext
}
