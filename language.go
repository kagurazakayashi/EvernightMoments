package main

import (
	"os"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type I18nManager struct {
	printer *message.Printer
	matcher language.Matcher
	support []language.Tag
}

func RegisterTranslations(tag language.Tag, translations map[string]string) {
	for key, value := range translations {
		message.SetString(tag, key, value)
	}
}

func NewI18nManager() *I18nManager {
	supported := []language.Tag{
		language.English,
		language.SimplifiedChinese,
		language.TraditionalChinese,
		language.Japanese,
	}

	Language_en()
	Language_zhHans()
	Language_zhHant()
	Language_ja()

	mgr := &I18nManager{
		matcher: language.NewMatcher(supported),
		support: supported,
	}

	mgr.SetLanguage(mgr.GetSystemLanguage())
	return mgr
}

func (m *I18nManager) GetSystemLanguage() string {
	for _, env := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if l := os.Getenv(env); l != "" {
			return strings.Split(l, ".")[0]
		}
	}
	return "en"
}

func (m *I18nManager) SetLanguage(langStr string) {
	tag, _, _ := m.matcher.Match(language.Make(langStr))
	m.printer = message.NewPrinter(tag)
}

func (m *I18nManager) T(key string, args ...interface{}) string {
	text := m.printer.Sprintf(key, args...)
	return text
}
