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

// renameEntry 記錄主檔案更名前後的完整檔名與基底名稱，供同步檔案匹配使用
type renameEntry struct {
	oldFullName string // 更名前含副檔名的完整檔名（如 "KYS0001.ARW"）
	newFullName string // 更名後含副檔名的完整檔名（如 "20260214_171635_KYS0001.ARW"）
	newBase     string // 更名後不含副檔名的基底名稱（如 "20260214_171635_KYS0001"）
}

// runRenameMode 執行更名模式的完整流程
// 參數 files 是從命令列傳入的檔案或目錄路徑清單
// 參數 cliOpts 包含從命令列旗標解析出的設定覆蓋值，可為 nil
func runRenameMode(files []string, cliOpts *CLIFlags) {
	// 先在載入設定前依 CLI 旗標預設色彩輸出狀態
	// 避免第一行輸出（標題分隔線）在 -nc 時仍產生色彩碼
	if cliOpts != nil && cliOpts.NoColor != nil {
		noColorOutput = *cliOpts.NoColor
	}

	fmt.Println(Dim(outLine))
	fmt.Println(Bold(evernightMoments + " v" + evernightMomentsVersion))

	// 載入設定檔並初始化多國語言
	conf := LoadConfig()

	// 套用命令列覆蓋值（若使用者有指定）
	if cliOpts != nil {
		applyCLIOverrides(&conf, cliOpts)
	}

	// 依設定最終決定是否停用彩色輸出
	noColorOutput = conf.NoColor

	var plans []RenamePlan
	counter := 1 // 檔名序號計數器，用於 <#> 標籤
	i18n.SetLanguage(conf.Language)
	// 依設定決定實際使用的 ExifTool 路徑（未設定時自動偵測）
	ApplyExiftoolConfig(conf)

	// 從 CLI 覆蓋選項中取得遞迴旗標（向後相容舊的 -r 手動解析）
	recursive := false
	if cliOpts != nil {
		recursive = cliOpts.Recursive
	}
	cleanFiles := files

	fmt.Println(Dim(outLine))
	fmt.Println(Dim(i18n.T("当前格式") + ": ") + conf.Format)
	fmt.Println(Dim(i18n.T("要配置格式") + " " + evernightMoments))
	fmt.Println(Dim(GetExiftoolPathI18n()))
	fmt.Println(Dim(outLine))
	fmt.Println(Cyan(i18n.T("正在分析") + "..."))

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
		if matchesAnyPattern(path, conf.Exclude) {
			continue
		}
		if matchesAnyPattern(path, conf.Sync) {
			syncPaths = append(syncPaths, path)
		} else {
			primaryPaths = append(primaryPaths, path)
		}
	}

	totalFiles := len(primaryPaths) + len(syncPaths)
	if totalFiles == 0 {
		fmt.Println(Yellow(i18n.T("没有文件")))
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
			fmt.Printf("[%s/%d] %s %s: %s\n", PadNumberByReference(i+1, totalFiles), totalFiles, Yellow(i18n.T("跳过")), i18n.T("文件错误"), Red(path))
			continue
		}

		rawTimeStr := t.Format("2006-01-02 15:04:05")
		newName := GenerateNewName(conf.Format, t, path, counter, conf.MultiExt)

		if isInvalid, char := ContainsInvalidChars(newName); isInvalid {
			fmt.Printf("[%s/%d] %s\n", PadNumberByReference(i+1, totalFiles), totalFiles, Red(i18n.T("非法字符格式", char)))
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

		// 記錄主檔案的新舊名稱對照
		// 同時以舊基底名稱與舊完整檔名作為鍵，處理雙副檔名伴隨檔案的精確匹配
		oldFullName := filepath.Base(path)
		oldBase := strings.TrimSuffix(oldFullName, filepath.Ext(oldFullName))
		newBase := strings.TrimSuffix(newName, filepath.Ext(newName))
		entry := renameEntry{
			oldFullName: oldFullName,
			newFullName: newName,
			newBase:     newBase,
		}
		renameMap[filepath.Join(dir, oldFullName)] = entry
		renameMap[filepath.Join(dir, oldBase)] = entry

		fmt.Printf("[%s/%d] %s : %s\n", Cyan(PadNumberByReference(i+1, totalFiles)), totalFiles, i18n.T("原文件"), Cyan(absPath))
		fmt.Printf("-> %s : %s : %s\n", i18n.T("依据"), Cyan(source), Yellow(rawTimeStr))
		fmt.Println("-> " + i18n.T("新文件名") + " : " + Green(newPath))
		counter++
	}

	// 5. 處理同步檔案：依照主檔案的更名結果同步更名
	syncIdx := len(primaryPaths)
	for _, syncPath := range syncPaths {
		dir := filepath.Dir(syncPath)
		syncExt := filepath.Ext(syncPath)
		syncStem := strings.TrimSuffix(filepath.Base(syncPath), syncExt)

		// 先嘗試以 syncStem 直接匹配，失敗則逐層剝離副檔名後再試
		// 這用於處理雙副檔名伴隨檔案（例如 KYS0001.ARW.dop 對應主檔案 KYS0001.ARW）
		entry, ok := renameMap[filepath.Join(dir, syncStem)]
		matchedKey := syncStem
		if !ok {
			altStem := syncStem
			for {
				ext := filepath.Ext(altStem)
				if ext == "" {
					break
				}
				altStem = strings.TrimSuffix(altStem, ext)
				if entry, ok = renameMap[filepath.Join(dir, altStem)]; ok {
					matchedKey = altStem
					break
				}
			}
		}
		if !ok {
			continue
		}

		syncIdx++
		// 根據匹配到的鍵決定新檔名：
		//   - 若匹配到主檔案的完整舊檔名（含副檔名），保留主副檔名 + 伴隨副檔名
		//     例如 KYS0001.ARW.dop → 20260214_171635_KYS0001.ARW.dop
		//   - 若匹配到主檔案的基底名稱（不含副檔名），僅使用新基底 + 伴隨副檔名
		//     例如 KYS0001.dop → 20260214_171635_KYS0001.dop
		var newName string
		if matchedKey == entry.oldFullName {
			newName = entry.newFullName + syncExt
		} else {
			newName = entry.newBase + syncExt
		}
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

		fmt.Printf("[%s/%d] %s : %s\n", Cyan(PadNumberByReference(syncIdx, totalFiles)), totalFiles, Yellow(i18n.T("原文件")), Cyan(absPath))
		fmt.Printf("-> %s: %s\n", i18n.T("同步依据"), Cyan(syncPath))
		fmt.Println("-> " + i18n.T("新文件名") + " : " + Green(newPath))
	}

	// 若未發現符合條件的檔案，則提示並結束
	if len(plans) == 0 {
		fmt.Println(Yellow(i18n.T("没有文件")))
		return
	}

	// 6. 使用者確認邏輯：根據設定決定是否顯示預覽確認
	proceed := true
	fmt.Println(Dim(outLine))
	if conf.Confirm {
		fmt.Printf("%s%s? (y/N): ", Yellow(i18n.T("共计", len(plans))), Yellow(i18n.T("确认")))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input != "y" && input != "yes" {
			proceed = false
		}
	}

	// 7. 正式執行檔案更名作業
	if proceed {
		fmt.Println(Cyan(i18n.T("开始") + "..."))
		successCount := 0
		for i, p := range plans {
			var totalPlans int = len(plans)
			fmt.Printf("[%s/%d] %s : %s\n", Cyan(PadNumberByReference(i+1, totalPlans)), totalPlans, i18n.T("原文件"), Cyan(p.AbsPath))
			fmt.Printf("-> %s : %s : %s\n", i18n.T("依据"), Cyan(p.Source), Yellow(p.RawTime))

			if p.OldPath == p.NewPath {
				fmt.Println("-> " + Yellow(i18n.T("跳过")) + " : " + Yellow(i18n.T("无变化")))
				continue
			}

			fmt.Println("-> " + i18n.T("新文件名") + " : " + Green(p.NewPath))
			if _, err := os.Stat(p.NewPath); err == nil {
				fmt.Println("-> " + Red(i18n.T("重命名失败")) + " : " + Red(i18n.T("已存在")))
				continue
			}

			err := os.Rename(p.OldPath, p.NewPath)
			if err != nil {
				fmt.Printf("-> %s: %v\n", Red(i18n.T("重命名失败")), err)
			} else {
				fmt.Println("-> " + Green(i18n.T("重命名成功")))
				successCount++
			}
		}
		fmt.Println(Green(i18n.T("处理结果", successCount, len(plans)-successCount, len(plans))))
	} else {
		fmt.Println(Yellow(i18n.T("取消")))
	}

	if conf.EndPause {
		EndPause()
	}
}
