package main

import (
	"golang.org/x/text/language"
)

func Language_ja() {
	RegisterTranslations(language.Japanese, map[string]string{
		"介绍1":    "瞬きに永遠を、常夜に温もりを。",
		"介绍2":    "は、写真の撮影日時を抽出し、ファイル名を自動でリネームするツールです。",
		"cancel": "Cancel",
	})
}
