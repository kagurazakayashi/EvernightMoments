package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var i18n = NewI18nManager()

func runConfigMode() {
	conf := LoadConfig()
	reader := bufio.NewReader(os.Stdin)
	if conf.Language == "" {
		conf.Language = i18n.GetSystemLanguage()
	}
	fmt.Println("Current language: " + conf.Language)
	fmt.Println("Change display language (Press Enter to keep current):")
	fmt.Println("1. English  2. 简体中文  3. 繁體中文  4. 日本語")
	fmt.Print("> ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		choice, err := strconv.Atoi(input)
		if err == nil && choice >= 1 && choice <= len(i18n.support) {
			selectedTag := i18n.support[choice-1]
			conf.Language = selectedTag.String()
		} else {
			fmt.Println("Invalid input. No changes made.")
		}
	}
	i18n.SetLanguage(conf.Language)
	fmt.Println("Current language: " + conf.Language)

	fmt.Println(outLine)
	fmt.Println(evernightMoments + " v" + evernightMomentsVersion)
	fmt.Println(i18n.T("介绍1"))
	fmt.Println(evernightMoments + " " + i18n.T("介绍2"))
	fmt.Println("https://github.com/kagurazakayashi/EvernightMoments")
	fmt.Println(outLine)
	fmt.Println(i18n.T("使用方式", evernightMoments))
	fmt.Println(outLine)
	fmt.Println(i18n.T("软件配置"))

	fmt.Println(outLine)
	fmt.Println("(1/3) " + i18n.T("配置文件名命名格式"))
	fmt.Println(i18n.T("当前格式") + ": " + conf.Format)
	fmt.Println(i18n.T("可用变量"))
	exampleName := GenerateNewName(defaultFormat, time.Now(), "Photo.jpg", 1)
	fmt.Print(i18n.T("默认格式例", defaultFormat, exampleName))

	var formatInput string
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			formatInput = conf.Format
			break
		}

		hasInvalid, char := ContainsInvalidChars(input)
		if hasInvalid {
			fmt.Println(i18n.T("非法字符格式", char))
			fmt.Println(i18n.T("非法字符"))
			fmt.Print(i18n.T("重新输入") + " > ")
			continue
		}

		formatInput = input
		break
	}
	conf.Format = formatInput
	fmt.Printf("-> "+i18n.T("格式已设定")+": %s\n", conf.Format)
	fmt.Printf("-> "+i18n.T("示例输出")+": %s\n", exampleName)

	fmt.Println(outLine)
	fmt.Println("(2/3) " + i18n.T("询问预览"))
	fmt.Printf("%s: %v\n", i18n.T("当前设置"), conf.Confirm)
	fmt.Print(i18n.T("询问预览说明"))
	confirmInput, _ := reader.ReadString('\n')
	confirmInput = strings.TrimSpace(strings.ToLower(confirmInput))
	if confirmInput == "n" {
		conf.Confirm = false
		fmt.Println("-> " + i18n.T("预览确认关"))
	} else {
		conf.Confirm = true
		fmt.Println("-> " + i18n.T("预览确认开"))
	}

	fmt.Println(outLine)
	fmt.Println("(3/3) " + i18n.T("结束等待"))
	fmt.Printf("%s: %v\n", i18n.T("当前设置"), conf.EndPause)
	fmt.Print(i18n.T("结束等待说明"))
	confirmInput, _ = reader.ReadString('\n')
	confirmInput = strings.TrimSpace(strings.ToLower(confirmInput))
	if confirmInput == "n" {
		conf.EndPause = false
		fmt.Println("-> " + i18n.T("结束等待关"))
	} else {
		conf.EndPause = true
		fmt.Println("-> " + i18n.T("结束等待开"))
	}

	fmt.Println(outLine)
	configPath := getConfigPath()
	if err := SaveConfig(conf, configPath); err != nil {
		fmt.Println(i18n.T("保存配置失败", configPath, err))
	} else {
		fmt.Println(i18n.T("保存配置成功") + configPath)
	}
	if conf.EndPause {
		EndPause()
	}
}
