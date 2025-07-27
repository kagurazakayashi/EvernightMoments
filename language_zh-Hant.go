package main

import (
	"golang.org/x/text/language"
)

func Language_zhHant() {
	RegisterTranslations(language.TraditionalChinese, map[string]string{
		"介绍1":    "予瞬息以永恆，於長夜留餘溫。",
		"介绍2":    "是一款透過提取照片原始拍攝時間，為您自動重新命名的工具。",
		"cancel": "Cancel",
	})
}
