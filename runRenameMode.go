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

// renameEntry 記錄主檔案更名後的基底名稱（不含副檔名），供同步檔案匹配使用
type renameEntry struct {
	newBase string // 更名後不含副檔名的基底名稱
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
	// 依設定決定實際使用的 ExifTool 路徑（未設定時自動偵測）
	ApplyExiftoolConfig(conf)

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

	// 3. 根據排除與同步設定，將檔案分為主要檔案與同步檔案
	var primaryPaths []string
	var syncPaths []string
	for _, path := range allPaths {
		filename := filepath.Base(path)
		if matchesAnyPattern(filename, conf.Exclude) {
			continue
		}
		if matchesAnyPattern(filename, conf.Sync) {
			syncPaths = append(syncPaths, path)
		} else {
			primaryPaths = append(primaryPaths, path)
		}
	}

	totalFiles := len(primaryPaths) + len(syncPaths)
	if totalFiles == 0 {
		fmt.Println(i18n.T("没有文件"))
		return
	}

	// 4. 先處理主要檔案：生成更名計劃並建立名稱對照表
	renameMap := make(map[string]renameEntry)

	for i, path := range primaryPaths {
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

		// 記錄主檔案的舊基底名稱 → 新基底名稱對照
		oldBase := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
		newBase := strings.TrimSuffix(newName, filepath.Ext(newName))
		renameMap[filepath.Join(dir, oldBase)] = renameEntry{newBase: newBase}

		fmt.Printf("[%s/%d] %s : %s\n", PadNumberByReference(i+1, totalFiles), totalFiles, i18n.T("原文件"), absPath)
		fmt.Printf("-> %s : %s : %s\n", i18n.T("依据"), source, rawTimeStr)
		fmt.Printf("-> %s : %s\n", i18n.T("新文件名"), newPath)
		counter++
	}

	// 5. 處理同步檔案：依照主檔案的更名結果同步更名
	syncIdx := len(primaryPaths)
	for _, syncPath := range syncPaths {
		dir := filepath.Dir(syncPath)
		syncExt := filepath.Ext(syncPath)
		syncBase := strings.TrimSuffix(filepath.Base(syncPath), syncExt)

		entry, ok := renameMap[filepath.Join(dir, syncBase)]
		if !ok {
			// 找不到對應的主檔案，跳過
			continue
		}

		syncIdx++
		newName := entry.newBase + syncExt
		newPath := filepath.Join(dir, newName)
		absPath, _ := filepath.Abs(syncPath)

		plans = append(plans, RenamePlan{
			OldPath: syncPath,
			AbsPath: absPath,
			NewPath: newPath,
			NewName: newName,
			Source:  i18n.T("同步依据"),
			RawTime: "-",
		})

		fmt.Printf("[%s/%d] %s : %s\n", PadNumberByReference(syncIdx, totalFiles), totalFiles, i18n.T("原文件"), absPath)
		fmt.Printf("-> %s: %s\n", i18n.T("同步依据"), syncPath)
		fmt.Printf("-> %s : %s\n", i18n.T("新文件名"), newPath)
	}

	// 若未發現符合條件的檔案，則提示並結束
	if len(plans) == 0 {
		fmt.Println(i18n.T("没有文件"))
		return
	}

	// 6. 使用者確認邏輯：根據設定決定是否顯示預覽確認
	proceed := true
	fmt.Println(outLine)
	if conf.Confirm {
		fmt.Printf("%s%s? (y/N): ", i18n.T("共计", len(plans)), i18n.T("确认"))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input != "y" && input != "yes" {
			proceed = false
		}
	}

	// 7. 正式執行檔案更名作業
	if proceed {
		fmt.Println(i18n.T("开始") + "...")
		successCount := 0
		for i, p := range plans {
			var totalPlans int = len(plans)
			fmt.Printf("[%s/%d] %s : %s\n", PadNumberByReference(i+1, totalPlans), totalPlans, i18n.T("原文件"), p.AbsPath)
			fmt.Printf("-> %s : %s : %s\n", i18n.T("依据"), p.Source, p.RawTime)

			if p.OldPath == p.NewPath {
				fmt.Printf("-> %s : %s\n", i18n.T("跳过"), i18n.T("无变化"))
				continue
			}

			fmt.Printf("-> %s : %s\n", i18n.T("新文件名"), p.NewPath)
			if _, err := os.Stat(p.NewPath); err == nil {
				fmt.Println("-> " + i18n.T("重命名失败") + i18n.T("已存在"))
				continue
			}

			err := os.Rename(p.OldPath, p.NewPath)
			if err != nil {
				fmt.Printf("-> %s: %v\n", i18n.T("重命名失败"), err)
			} else {
				fmt.Println("-> " + i18n.T("重命名成功"))
				successCount++
			}
		}
		fmt.Println(i18n.T("处理结果", successCount, len(plans)-successCount, len(plans)))
	} else {
		fmt.Println(i18n.T("取消"))
	}

	if conf.EndPause {
		EndPause()
	}
}
