package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runConfigMode() {
	currentFormat := LoadFormat()
	fmt.Println("=== 照片重命名工具配置 ===")
	fmt.Printf("当前保存的格式为: %s\n", currentFormat)
	fmt.Println("请输入新的文件名格式 (直接回车将恢复默认 `YYYYMMDD_HHmmss_*` ):")
	fmt.Println("变量提示: YYYY/MM/DD, HH/mm/ss, ##(编号), *(原名)")
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var finalFormat string
	if input == "" {
		// 使用者直接回車，嘗試刪除配置
		err := DeleteFormat()
		if err != nil && !os.IsNotExist(err) {
			fmt.Printf("重置配置失败: %v\n", err)
		} else {
			fmt.Println("已恢复默认设置。")
		}
		finalFormat = defaultFormat
	} else {
		// 儲存新配置
		if err := SaveFormat(input); err != nil {
			fmt.Printf("保存配置失败: %v\n", err)
			return
		}
		fmt.Printf("格式已更新为: %s\n", input)
		finalFormat = input
	}

	// 輸出示例預覽
	fmt.Println(finalFormat)

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
