package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RenamePlan 儲存單一檔案的更名任務資訊
type RenamePlan struct {
	OldPath string // 原始相對路徑
	AbsPath string // 檔案的絕對路徑
	NewPath string // 預計更名後的新路徑
	NewName string // 預計生成的新檔名
	Source  string // 時間資訊的來源（例如：EXIF 或 修改日期）
	RawTime string // 格式化後的原始時間字串
}

// runRenameMode 執行更名模式的完整流程
// 參數 files 是從命令列傳入的檔案或目錄路徑清單
func runRenameMode(files []string) {
	fmt.Println(outLine)
	fmt.Println(evernightMoments + " v" + evernightMomentsVersion)

	// 載入設定檔並初始化多國語言
	conf := LoadConfig()
	var plans []RenamePlan
	counter := 1 // 檔名序號計數器，用於 <#> 標籤
	i18n.SetLanguage(conf.Language)

	// 1. 處理參數：從傳入的檔案清單中提取並檢查是否包含遞迴標記 (-r)
	recursive := false
	var cleanFiles []string
	for _, f := range files {
		if f == "-r" {
			recursive = true
		} else {
			cleanFiles = append(cleanFiles, f)
		}
	}

	fmt.Println(outLine)
	fmt.Println(i18n.T("当前格式") + ": " + conf.Format)
	fmt.Println(i18n.T("要配置格式") + " " + evernightMoments)
	fmt.Println(GetExiftoolPathI18n())
	fmt.Println(outLine)
	fmt.Println(i18n.T("正在分析") + "...")

	// 2. 預掃描邏輯，先獲取所有檔案路徑
	var allPaths []string
	for _, pattern := range cleanFiles {
		matches, _ := filepath.Glob(pattern)
		for _, path := range matches {
			info, err := os.Stat(path)
			if err != nil {
				continue
			}
			if info.IsDir() {
				if recursive {
					filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
						if err == nil && !d.IsDir() {
							allPaths = append(allPaths, p)
						}
						return nil
					})
				} else {
					entries, _ := os.ReadDir(path)
					for _, entry := range entries {
						if !entry.IsDir() {
							allPaths = append(allPaths, filepath.Join(path, entry.Name()))
						}
					}
				}
			} else {
				allPaths = append(allPaths, path)
			}
		}
	}

	totalFiles := len(allPaths) // 總檔案數
	if totalFiles == 0 {
		fmt.Println(i18n.T("没有文件"))
		return
	}

	// 3. 走訪邏輯：支援多種輸入方式（檔案、目錄、萬用字元）
	for i, path := range allPaths {
		info, err := os.Stat(path)
		if err != nil || info.IsDir() {
			continue
		}

		t, source, err := GetPhotoTime(path)
		if err != nil {
			fmt.Printf("[%d/%d] %s %s: %s\n", i+1, totalFiles, i18n.T("跳过"), i18n.T("文件错误"), path)
			continue
		}

		rawTimeStr := t.Format("2006-01-02 15:04:05")
		newName := GenerateNewName(conf.Format, t, path, counter)

		// 检查非法字符
		if isInvalid, char := ContainsInvalidChars(newName); isInvalid {
			fmt.Printf("[%d/%d] %s\n", i+1, totalFiles, i18n.T("非法字符格式", char))
			continue
		}

		dir := filepath.Dir(path)
		newPath := filepath.Join(dir, newName)
		absPath, _ := filepath.Abs(path)

		plans = append(plans, RenamePlan{
			OldPath: path,
			AbsPath: absPath,
			NewPath: newPath,
			NewName: newName,
			Source:  source,
			RawTime: rawTimeStr,
		})

		// 修改后的打印输出：[当前序号/总数]
		fmt.Printf("[%s/%d] %s : %s\n", PadNumberByReference(i+1, totalFiles), totalFiles, i18n.T("原文件"), absPath)
		fmt.Printf("-> %s : %s : %s\n", i18n.T("依据"), source, rawTimeStr)
		fmt.Printf("-> %s : %s\n", i18n.T("新文件名"), newPath)
		counter++
	}

	// 若未發現符合條件的檔案，則提示並結束
	if len(plans) == 0 {
		fmt.Println(i18n.T("没有文件"))
		return
	}

	// 4. 使用者確認邏輯：根據設定決定是否顯示預覽確認
	proceed := true
	fmt.Println(outLine)
	if conf.Confirm {
		// 顯示總計數量並詢問使用者
		fmt.Printf("%s%s? (y/N): ", i18n.T("共计", len(plans)), i18n.T("确认"))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		// 使用者必須輸入 'y' 或 'yes' 才會繼續執行
		if input != "y" && input != "yes" {
			proceed = false
		}
	}

	// 5. 正式執行檔案更名作業
	if proceed {
		fmt.Println(i18n.T("开始") + "...")
		successCount := 0
		for i, p := range plans {
			var totalPlans int = len(plans)
			fmt.Printf("[%s/%d] %s : %s\n", PadNumberByReference(i+1, totalPlans), totalPlans, i18n.T("原文件"), p.AbsPath)
			fmt.Printf("-> %s : %s : %s\n", i18n.T("依据"), p.Source, p.RawTime)

			// 檢查檔名是否真的有變動，若無則跳過
			if p.OldPath == p.NewPath {
				fmt.Printf("-> %s : %s\n", i18n.T("跳过"), i18n.T("无变化"))
				continue
			}

			// fmt.Printf("-> %s : %s\n", i18n.T("新文件名"), p.NewName)
			fmt.Printf("-> %s : %s\n", i18n.T("新文件名"), p.NewPath)
			// 檢查目標路徑是否已經存在檔案，避免覆蓋
			if _, err := os.Stat(p.NewPath); err == nil {
				fmt.Println("-> " + i18n.T("重命名失败") + i18n.T("已存在"))
				continue
			}

			// 執行實體檔案更名
			err := os.Rename(p.OldPath, p.NewPath)
			if err != nil {
				fmt.Printf("-> %s: %v\n", i18n.T("重命名失败"), err)
			} else {
				fmt.Println("-> " + i18n.T("重命名成功"))
				successCount++
			}
			// fmt.Println("-> DEBUG MODE") //DEBUG
			// successCount++                       //DEBUG
		}
		// 輸出最終處理結果統計
		fmt.Println(i18n.T("处理结果", successCount, len(plans)-successCount, len(plans)))
	} else {
		fmt.Println(i18n.T("取消"))
	}

	// 根據設定決定是否在視窗關閉前暫停
	if conf.EndPause {
		EndPause()
	}
}
