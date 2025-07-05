package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Format  string `json:"format"`
	Confirm bool   `json:"confirm"` // 是否需要預覽確認
}

const defaultFormat = "YYYYMMDD_HHmmss_*"

// 獲取配置檔案路徑
func getConfigPath() string {
	exePath, err := os.Executable()
	if err != nil {
		return "config.json"
	}
	dir := filepath.Dir(exePath)
	base := filepath.Base(exePath)
	name := strings.TrimSuffix(base, filepath.Ext(base))
	return filepath.Join(dir, name+".json")
}

// LoadConfig 讀取完整配置
func LoadConfig() Config {
	configPath := getConfigPath()
	data, err := os.ReadFile(configPath)
	// 預設配置：開啟確認預覽
	defaultConf := Config{Format: defaultFormat, Confirm: true}

	if err != nil {
		return defaultConf
	}

	var conf Config
	if err := json.Unmarshal(data, &conf); err != nil {
		return defaultConf
	}

	// 如果讀取到的格式為空，回退到預設格式
	if conf.Format == "" {
		conf.Format = defaultFormat
	}
	return conf
}

// SaveConfig 儲存完整配置
func SaveConfig(conf Config, configPath string) error {
	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}
