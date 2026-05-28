package main

import (
	"os"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// I18nManager 負責管理多國語言翻譯與語系匹配
type I18nManager struct {
	printer *message.Printer // 用於執行格式化輸出的列印器
	matcher language.Matcher // 語系匹配器，負責尋找最接近的支援語言
	support []language.Tag   // 目前系統支援的語系標籤清單
}

// RegisterTranslations 為指定語系註冊翻譯對應表
// tag: 目標語系標籤（例如 language.English）
// translations: 鍵值對形式的翻譯對應表
func RegisterTranslations(tag language.Tag, translations map[string]string) {
	for key, value := range translations {
		// 將翻譯字串設定到全域的消息系統中
		message.SetString(tag, key, value)
	}
}

// NewI18nManager 初始化並回傳一個新的 I18n 管理器執行個體
func NewI18nManager() *I18nManager {
	// 定義目前程式支援的語系
	supported := []language.Tag{
		language.English,            // 英文
		language.SimplifiedChinese,  // 簡體中文
		language.TraditionalChinese, // 繁體中文
		language.Japanese,           // 日文
	}

	// 呼叫各語系的初始化函式（通常由外部產生的翻譯資料定義）
	Language_en()
	Language_zhHans()
	Language_zhHant()
	Language_ja()

	mgr := &I18nManager{
		// 根據支援清單建立匹配器
		matcher: language.NewMatcher(supported),
		support: supported,
	}

	// 根據目前作業系統的環境變數設定語系
	mgr.SetLanguage(mgr.GetSystemLanguage())
	return mgr
}

// GetSystemLanguage 嘗試從作業系統環境變數中取得目前的語系設定
func (m *I18nManager) GetSystemLanguage() string {
	// 依序檢查常見的語系環境變數
	for _, env := range []string{"LC_ALL", "LC_MESSAGES", "LANG"} {
		if l := os.Getenv(env); l != "" {
			// 處理格式如 "zh_TW.UTF-8"，僅取 "." 之前的語言標記部分
			return strings.Split(l, ".")[0]
		}
	}
	// 若皆未設定，則預設回傳英文
	return "en"
}

// SetLanguage 根據傳入的語系字串，匹配並切換目前的翻譯語系
func (m *I18nManager) SetLanguage(langStr string) {
	// 將字串轉換為標籤，並透過 matcher 找到最適合的支援語系
	tag, _, _ := m.matcher.Match(language.Make(langStr))
	// 根據匹配結果建立新的列印器
	m.printer = message.NewPrinter(tag)
}

// MatchIndex 依據傳入的語系字串，回傳其在 support 清單中最接近的索引
// 用於 TUI 下拉選單初始化時定位目前語系所對應的選項位置
func (m *I18nManager) MatchIndex(langStr string) int {
	// matcher.Match 的第二個回傳值即為命中支援清單中的索引
	_, idx, _ := m.matcher.Match(language.Make(langStr))
	return idx
}

// T 根據指定的 Key 執行翻譯，並支援帶入參數進行格式化（類似 Sprintf）
func (m *I18nManager) T(key string, args ...interface{}) string {
	// 使用內部 printer 執行翻譯與字串替換
	text := m.printer.Sprintf(key, args...)
	return text
}
