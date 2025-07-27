package main

import (
	"golang.org/x/text/language"
)

func Language_en() {
	RegisterTranslations(language.English, map[string]string{
		"介绍1":    "Bestowing eternity upon the fleeting, and warmth to the everlasting night.",
		"介绍2":    "is a utility that automatically renames your visual archives by extracting original capture timestamps.",
		"cancel": "Cancel",
	})
}
