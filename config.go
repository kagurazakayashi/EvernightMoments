package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Config 定義了應用程式的設定結構
type Config struct {
	Language string   `json:"language"`          // 介面語言設定
	Format   string   `json:"format"`            // 輸出格式
	Exclude  []string `json:"exclude,omitempty"` // 排除的副檔名樣式，例如 ["*.xml", "*.txt"]
	Sync     []string `json:"sync,omitempty"`    // 同步更名的副檔名樣式，例如 ["*.txt", "*.xmp"]
	Confirm  bool     `json:"confirm"`           // 是否在執行前顯示確認預覽
	EndPause bool     `json:"endpause"`          // 程式結束後是否暫停（避免視窗直接關閉）
	// ExiftoolPath 指定 ExifTool 可執行檔的路徑，用以讀取更精確的拍攝時間。
	// 使用指標以區分三種狀態：
	//   nil      ：設定檔尚未記錄此項，執行時自動偵測（向後相容舊設定檔）
	//   指向 ""  ：使用者明確留空，代表僅使用內建解析、不使用 ExifTool
	//   指向路徑 ：使用指定路徑的 ExifTool
	ExiftoolPath *string `json:"exiftoolpath,omitempty"`
}

// getConfigDir 根據作業系統回傳標準的設定檔目錄路徑
func getConfigDir() string {
	var base string
	switch runtime.GOOS {
	case "windows":
		base = os.Getenv("APPDATA")
		if base == "" {
			home, _ := os.UserHomeDir()
			base = filepath.Join(home, "AppData", "Roaming")
		}
	case "darwin":
		home, _ := os.UserHomeDir()
		base = filepath.Join(home, "Library", "Application Support")
	default: // linux, freebsd 等遵循 XDG 規範的系統
		base = os.Getenv("XDG_CONFIG_HOME")
		if base == "" {
			home, _ := os.UserHomeDir()
			base = filepath.Join(home, ".config")
		}
	}
	return filepath.Join(base, evernightMoments)
}

// getLegacyConfigPath 回傳舊版（與執行檔同目錄）的設定檔路徑
func getLegacyConfigPath() string {
	exePath, err := os.Executable()
	if err != nil {
		return "config.json"
	}
	dir := filepath.Dir(exePath)
	base := filepath.Base(exePath)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	return filepath.Join(dir, name+".json")
}

// getConfigPath 回傳標準位置的設定檔完整路徑
func getConfigPath() string {
	configDir := getConfigDir()
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "config.json"
	}
	return filepath.Join(configDir, "config.json")
}

// migrateConfig 若標準位置尚無設定檔，但舊版位置存在，則自動遷移
func migrateConfig() {
	newPath := getConfigPath()
	if _, err := os.Stat(newPath); err == nil {
		return
	}

	legacyPath := getLegacyConfigPath()
	data, err := os.ReadFile(legacyPath)
	if err != nil {
		return
	}

	configDir := filepath.Dir(newPath)
	os.MkdirAll(configDir, 0755)
	os.WriteFile(newPath, data, 0644)
}

// LoadConfig 從硬碟讀取設定檔，若檔案不存在或解析出錯則回傳預設值
func LoadConfig() Config {
	// 若舊版設定檔存在而新版尚無，自動遷移
	migrateConfig()

	configPath := getConfigPath()

	// 讀取檔案內容
	data, err := os.ReadFile(configPath)

	// 初始化預設配置：開啟確認預覽與結束暫停
	defaultConf := Config{Format: defaultFormat, Confirm: true, EndPause: true}

	// 若讀取失敗（例如檔案不存在），直接回傳預設設定
	if err != nil {
		return defaultConf
	}

	var conf Config
	// 解析 JSON 資料
	if err := json.Unmarshal(data, &conf); err != nil {
		// 若 JSON 格式有誤，回傳預設設定
		return defaultConf
	}

	// 檢查讀取到的格式欄位，若為空字串則回退到預設格式
	if conf.Format == "" {
		conf.Format = defaultFormat
	}
	return conf
}

// SaveConfig 將傳入的設定物件以格式化的 JSON 形式儲存到指定路徑
func SaveConfig(conf Config, configPath string) error {
	// 將結構體轉換為具備縮排（2個空格）的 JSON 格式
	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}

	// 將 JSON 資料寫入檔案，權限設定為 0644 (擁有者可讀寫，其餘人唯讀)
	return os.WriteFile(configPath, data, 0644)
}

// formatSlice 將字串切片格式化為逗號分隔的顯示字串，若為空則回傳「(無)」
func formatSlice(s []string) string {
	if len(s) == 0 {
		return "(none)"
	}
	return strings.Join(s, ", ")
}

// parsePatterns 將使用者輸入的逗號分隔字串解析為去除空白後的樣式切片
func parsePatterns(input string) []string {
	parts := strings.Split(input, ",")
	var result []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}

// matchesAnyPattern 檢查 filename 是否符合 patterns 中的任一 glob 樣式
func matchesAnyPattern(filename string, patterns []string) bool {
	for _, pattern := range patterns {
		if matched, _ := filepath.Match(pattern, filename); matched {
			return true
		}
	}
	return false
}
