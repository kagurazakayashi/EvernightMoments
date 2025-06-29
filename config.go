package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Config 配置檔案結構
type Config struct {
	Format string `json:"format"`
}

const defaultFormat = "YYYYMMDD_HHmmss_*"

// getConfigPath 獲取與程式同名的 JSON 配置檔案路徑
func getConfigPath() string {
	exePath, err := os.Executable()
	if err != nil {
		return "config.json"
	}

	// 獲取程式所在的目錄
	dir := filepath.Dir(exePath)
	// 獲取程式檔名（如 EvernightMoments.exe）
	base := filepath.Base(exePath)
	// 去掉副檔名（如 EvernightMoments）
	name := strings.TrimSuffix(base, filepath.Ext(base))

	// 返回 路徑/程式名.json
	return filepath.Join(dir, name+".json")
}

// SaveFormat 儲存格式到 JSON 檔案
func SaveFormat(format string) error {
	configPath := getConfigPath()
	conf := Config{Format: format}

	// 將結構體序列化為帶縮排的 JSON
	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(configPath)
	return os.WriteFile(configPath, data, 0644)
}

// LoadFormat 從 JSON 檔案讀取格式
func LoadFormat() string {
	configPath := getConfigPath()
	data, err := os.ReadFile(configPath)
	if err != nil {
		return defaultFormat
	}

	var conf Config
	if err := json.Unmarshal(data, &conf); err != nil {
		return defaultFormat
	}

	if conf.Format == "" {
		return defaultFormat
	}
	return conf.Format
}

// DeleteFormat 刪除配置檔案
func DeleteFormat() error {
	return os.Remove(getConfigPath())
}
