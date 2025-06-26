package main

import (
	"os"
	"path/filepath"
	"strings"
)

const configFile = ".format_config"
const defaultFormat = "YYYYMMDD_HHmmss_*"

// GetExecutableDir 獲取當前程式所在目錄
func GetExecutableDir() string {
	exe, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(exe)
}

// SaveFormat 儲存格式字串到配置檔案
func SaveFormat(format string) error {
	configPath := filepath.Join(GetExecutableDir(), configFile)
	return os.WriteFile(configPath, []byte(format), 0644)
}

// LoadFormat 讀取儲存的格式，如果不存在則返回預設值
func LoadFormat() string {
	configPath := filepath.Join(GetExecutableDir(), configFile)
	data, err := os.ReadFile(configPath)
	if err != nil {
		return defaultFormat
	}
	format := strings.TrimSpace(string(data))
	if format == "" {
		return defaultFormat
	}
	return format
}

// DeleteFormat 刪除配置檔案，恢復系統預設
func DeleteFormat() error {
	configPath := filepath.Join(GetExecutableDir(), configFile)
	return os.Remove(configPath)
}
