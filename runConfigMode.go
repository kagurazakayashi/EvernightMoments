package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// i18n 是全域的多國語言管理器執行個體，負責處理介面文字的翻譯
var i18n = NewI18nManager()

// runConfigMode 執行互動式設定模式，引導使用者透過命令列配置程式參數
func runConfigMode() {
	// 載入目前的設定檔，若無設定檔則取得預設值
	conf := LoadConfig()
	// 建立標準輸入（stdin）的讀取器
	reader := bufio.NewReader(os.Stdin)

	// 若設定檔中未指定語言，則嘗試讀取作業系統的語系
	if conf.Language == "" {
		conf.Language = i18n.GetSystemLanguage()
	}

	// --- 語系設定部分 ---
	fmt.Println("Current language: " + conf.Language)
	fmt.Println("Change display language (Press Enter to keep current):")
	fmt.Println("1. English  2. 简体中文  3. 繁體中文  4. 日本語")
	fmt.Print("> ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		// 將輸入的字串轉換為整數序號
		choice, err := strconv.Atoi(input)
		// 驗證輸入是否在有效範圍內（1 到 支援語系總數）
		if err == nil && choice >= 1 && choice <= len(i18n.support) {
			selectedTag := i18n.support[choice-1]
			conf.Language = selectedTag.String()
		} else {
			fmt.Println("Invalid input. No changes made.")
		}
	}
	// 更新當前執行個體的語系設定並套用
	i18n.SetLanguage(conf.Language)
	fmt.Println("Current language: " + conf.Language)

	// --- 顯示軟體資訊與說明 ---
	fmt.Println(outLine)
	fmt.Println(evernightMoments + " v" + evernightMomentsVersion)
	fmt.Println(i18n.T("介绍1"))
	fmt.Println(evernightMoments + " " + i18n.T("介绍2"))
	fmt.Println("https://github.com/kagurazakayashi/EvernightMoments")
	fmt.Println(GetExiftoolPathI18n())
	fmt.Println(outLine)
	fmt.Println(i18n.T("使用方式", evernightMoments))
	fmt.Println(outLine)
	fmt.Println(i18n.T("软件配置"))

	// --- 設定 1/5：命名格式 ---
	fmt.Println(outLine)
	fmt.Println("(1/5) " + i18n.T("配置文件名命名格式"))
	fmt.Println(i18n.T("当前格式") + ": " + conf.Format)
	fmt.Println(i18n.T("可用变量"))

	// 以目前的預設格式產生一個範例，讓使用者參考
	exampleName := GenerateNewName(defaultFormat, time.Now(), "Photo.jpg", 1)
	fmt.Print(i18n.T("默认格式例", defaultFormat, exampleName))

	var formatInput string
	// 進入無窮迴圈直到使用者輸入合法的格式或選擇維持現狀
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// 若直接按 Enter，則沿用原有的格式設定
		if input == "" {
			formatInput = conf.Format
			break
		}

		// 檢查輸入是否包含作業系統不允許的非法字元
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

	// 顯示新格式的預覽結果
	exampleName = GenerateNewName(conf.Format, time.Now(), "Photo.jpg", 1)
	fmt.Printf("-> "+i18n.T("示例输出")+": %s\n", exampleName)

	// --- 設定 2/5：排除副檔名 ---
	fmt.Println(outLine)
	fmt.Println("(2/5) " + i18n.T("排除项配置"))
	fmt.Println(i18n.T("当前排除项") + ": " + formatSlice(conf.Exclude))
	fmt.Println(i18n.T("排除项说明"))
	fmt.Print("> ")

	excludeInput, _ := reader.ReadString('\n')
	excludeInput = strings.TrimSpace(excludeInput)
	if excludeInput != "" {
		conf.Exclude = parsePatterns(excludeInput)
	}
	fmt.Printf("-> "+i18n.T("排除项已设定")+": %s\n", formatSlice(conf.Exclude))

	// --- 設定 3/5：同步副檔名 ---
	fmt.Println(outLine)
	fmt.Println("(3/5) " + i18n.T("同步项配置"))
	fmt.Println(i18n.T("当前同步项") + ": " + formatSlice(conf.Sync))
	fmt.Println(i18n.T("同步项说明"))
	fmt.Print("> ")

	syncInput, _ := reader.ReadString('\n')
	syncInput = strings.TrimSpace(syncInput)
	if syncInput != "" {
		conf.Sync = parsePatterns(syncInput)
	}
	fmt.Printf("-> "+i18n.T("同步项已设定")+": %s\n", formatSlice(conf.Sync))

	// --- 設定 4/5：確認預覽開關 ---
	fmt.Println(outLine)
	fmt.Println("(4/5) " + i18n.T("询问预览"))
	fmt.Printf("%s: %v\n", i18n.T("当前设置"), conf.Confirm)
	fmt.Print(i18n.T("询问预览说明"))

	confirmInput, _ := reader.ReadString('\n')
	confirmInput = strings.TrimSpace(strings.ToLower(confirmInput))
	// 只有輸入 'n' 才會關閉，其餘輸入（或直接 Enter）皆視為開啟
	if confirmInput == "n" {
		conf.Confirm = false
		fmt.Println("-> " + i18n.T("预览确认关"))
	} else {
		conf.Confirm = true
		fmt.Println("-> " + i18n.T("预览确认开"))
	}

	// --- 設定 5/5：結束後暫停視窗開關 ---
	fmt.Println(outLine)
	fmt.Println("(5/5) " + i18n.T("结束等待"))
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

	// --- 儲存與結束 ---
	fmt.Println(outLine)
	configPath := getConfigPath()
	// 將所有變更寫回 JSON 設定檔
	if err := SaveConfig(conf, configPath); err != nil {
		fmt.Println(i18n.T("保存配置失败", configPath, err))
	} else {
		fmt.Println(i18n.T("保存配置成功") + configPath)
	}

	// 若使用者開啟了結束等待功能，則呼叫 EndPause 避免視窗閃退
	if conf.EndPause {
		EndPause()
	}
}
