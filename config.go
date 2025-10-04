package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// Config 定義了應用程式的設定結構
type Config struct {
	Language string `json:"language"` // 介面語言設定
	Format   string `json:"format"`   // 輸出格式
	Confirm  bool   `json:"confirm"`  // 是否在執行前顯示確認預覽
	EndPause bool   `json:"endpause"` // 程式結束後是否暫停（避免視窗直接關閉）
}

// getConfigPath 根據目前執行檔的名稱與位置，自動推導設定檔 (.json) 的路徑
func getConfigPath() string {
	// 取得目前執行檔的完整路徑
	exePath, err := os.Executable()
	if err != nil {
		// 若無法取得路徑，則回傳目前目錄下的預設檔名
		return "config.json"
	}

	// 取得執行檔所在的目錄
	dir := filepath.Dir(exePath)
	// 取得執行檔檔名（包含副檔名）
	base := filepath.Base(exePath)
	// 去除執行檔的副檔名（例如 .exe），取得純檔名
	name := strings.TrimSuffix(base, filepath.Ext(base))

	// 將目錄、純檔名與 .json 副檔名組合起來
	return filepath.Join(dir, name+".json")
}

// LoadConfig 從硬碟讀取設定檔，若檔案不存在或解析出錯則回傳預設值
func LoadConfig() Config {
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
