package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runConfigMode() {
	conf := LoadConfig()
	fmt.Println("=== EvernightMoments 配置 ===")
	fmt.Printf("当前格式: %s\n", conf.Format)
	fmt.Printf("预览确认: %v\n", conf.Confirm)
	fmt.Println("---------------------------------------------------------")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("1. 请输入文件名格式 (直接回车将恢复默认格式):")
	fmt.Print("> ")
	formatInput, _ := reader.ReadString('\n')
	formatInput = strings.TrimSpace(formatInput)
	if formatInput == "" {
		conf.Format = defaultFormat
		fmt.Println("-> 格式已重置为默认。")
	} else {
		conf.Format = formatInput
		fmt.Printf("-> 格式已设定为: %s\n", conf.Format)
	}

	fmt.Println("\n2. 是否在重命名前开启预览确认? (y: 开启 / n: 直接重命名):")
	fmt.Print("> ")
	confirmInput, _ := reader.ReadString('\n')
	confirmInput = strings.TrimSpace(strings.ToLower(confirmInput))
	if confirmInput == "n" {
		conf.Confirm = false
		fmt.Println("-> 已关闭预览确认。请小心操作。")
	} else {
		conf.Confirm = true
		fmt.Println("-> 已开启预览确认。")
	}

	// 儲存配置
	configPath := getConfigPath()
	if err := SaveConfig(conf, configPath); err != nil {
		fmt.Printf("保存配置到 %s 失败: %v\n", configPath, err)
	} else {
		fmt.Println("\n配置已成功保存到 " + configPath)
	}

	fmt.Println("\n按回车键退出...")
	reader.ReadString('\n')
}

func runRenameMode(files []string) {
}

func main() {
	args := os.Args[1:]

	// --- 模式 1: 配置模式 (无参数) ---
	if len(args) == 0 {
		runConfigMode()
		return
	}

	// --- 模式 2: 重命名模式 (有参数) ---
	runRenameMode(args)
}
