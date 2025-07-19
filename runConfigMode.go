package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func runConfigMode() {
	fmt.Println(outLine)
	fmt.Println(evernightMoments + " v" + evernightMomentsVersion)
	fmt.Println("予瞬息以永恒，于长夜留余温。")
	fmt.Println(evernightMoments + " 是一款通过提取照片原始拍摄时间，为您自动重命名影像文件的工具。")
	fmt.Println("https://github.com/kagurazakayashi/EvernightMoments")
	fmt.Println(outLine)
	conf := LoadConfig()
	fmt.Println("使用方式: " + evernightMoments + " [照片文件1] [照片文件2] ...")
	fmt.Println("如果直接运行而不提供照片文件，则现在进入配置模式。")
	fmt.Println(outLine)
	fmt.Println("软件配置")
	fmt.Println(outLine)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("(1/2) 配置文件名命名格式")
	fmt.Println("当前格式: " + conf.Format)
	fmt.Println("可用变量:")
	fmt.Println("  YYYY/YY: 年 | MM/M: 月 | DD/D: 日")
	fmt.Println("  HH/H: 时    | mm/m: 分 | ss/s: 秒")
	fmt.Println("  ##: 编号(01) | #: 编号(1)")
	fmt.Println("  *: 原始文件名")
	fmt.Println("  字母数代表位数，例如 2025 的 YY 为 25, 1月的 M 为 1, MM 为 01")
	exampleName := GenerateNewName(defaultFormat, time.Now(), "Photo.jpg", 1)
	fmt.Println("  例如, 默认值格式 `" + defaultFormat + "` 将会输出 `" + exampleName + "`。")
	fmt.Println("请输入新的文件名命名格式并回车, 如果留空则保留当前设定。")
	fmt.Print("> ")
	formatInput, _ := reader.ReadString('\n')
	formatInput = strings.TrimSpace(formatInput)
	if formatInput == "" {
		formatInput = conf.Format
	}
	conf.Format = formatInput
	fmt.Printf("-> 格式已设定为: %s\n", conf.Format)
	fmt.Println(outLine)
	fmt.Println("(2/2) 是否在重命名前开启预览确认? ")
	fmt.Printf("当前设置: %v\n", conf.Confirm)
	fmt.Println("如果开启，会先向你展示会修改成什么样子，让你确认要不要继续。")
	fmt.Println("请输入 `y` 或者 `n`。默认值`y`: 每次先询问我, `n`: 直接开始重命名。")
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
	fmt.Println(outLine)
	configPath := getConfigPath()
	if err := SaveConfig(conf, configPath); err != nil {
		fmt.Printf("保存配置到 %s 失败: %v\n", configPath, err)
	} else {
		fmt.Println("配置已成功保存到 " + configPath)
	}
	fmt.Println("按回车键退出...")
	reader.ReadString('\n')
}
