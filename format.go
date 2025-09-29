package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

// GenerateNewName 根據格式字串生成新檔名
// format: 使用者定義的格式，例如 "<YYYY>-<MM>-<DD>_image_<#>"
// t: 照片時間
// originalPath: 原始檔案路徑
// index: 當前處理的檔案序號
func GenerateNewName(format string, t time.Time, originalPath string, index int) string {
	originalName := filepath.Base(originalPath)
	ext := filepath.Ext(originalName)
	nameWithoutExt := strings.TrimSuffix(originalName, ext)

	// 定義替換對映，所有關鍵字都加上了 <>
	// 建議順序：長標籤放在短標籤前面（例如 <YYYY> 在 <YY> 前），防止誤匹配
	replacements := []struct {
		old string
		new string
	}{
		{"<YYYY>", fmt.Sprintf("%04d", t.Year())},
		{"<YY>", fmt.Sprintf("%02d", t.Year()%100)},
		{"<MM>", fmt.Sprintf("%02d", t.Month())},
		{"<M>", fmt.Sprintf("%d", t.Month())},
		{"<DD>", fmt.Sprintf("%02d", t.Day())},
		{"<D>", fmt.Sprintf("%d", t.Day())},
		{"<HH>", fmt.Sprintf("%02d", t.Hour())},
		{"<H>", fmt.Sprintf("%d", t.Hour())},
		{"<mm>", fmt.Sprintf("%02d", t.Minute())},
		{"<m>", fmt.Sprintf("%d", t.Minute())},
		{"<ss>", fmt.Sprintf("%02d", t.Second())},
		{"<s>", fmt.Sprintf("%d", t.Second())},
		{"<##>", fmt.Sprintf("%02d", index)},
		{"<#>", fmt.Sprintf("%d", index)},
		{"<*>", nameWithoutExt},
	}

	result := format

	// 執行替換
	for _, r := range replacements {
		result = strings.ReplaceAll(result, r.old, r.new)
	}

	// 拼接副檔名
	return result + ext
}
