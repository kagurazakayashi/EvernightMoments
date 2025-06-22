package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runConfigMode() {
	fmt.Println("=== 照片重命名工具配置 ===")
	fmt.Println("请输入新的文件名格式 (直接回车保持不变):")
	fmt.Println("可用变量:")
	fmt.Println("  YYYY/YY: 年 | MM/M: 月 | DD/D: 日")
	fmt.Println("  HH/H: 时    | mm/m: 分 | ss/s: 秒")
	fmt.Println("  ##: 编号(01) | #: 编号(1)")
	fmt.Println("  *: 原始文件名")
	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println(input)
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
