package main

import (
	"golang.org/x/text/language"
)

func Language_zhHans() {
	RegisterTranslations(language.SimplifiedChinese, map[string]string{
		"介绍1":    "予瞬息以永恒，于长夜留余温。",
		"介绍2":    "是一款通过提取照片原始拍摄时间，为您自动重命名的工具。",
		"cancel": "Cancel",
	})
}
