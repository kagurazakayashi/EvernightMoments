package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type RenamePlan struct {
	OldPath string
	AbsPath string
	NewPath string
	NewName string
	Source  string
	RawTime string
}

func runRenameMode(files []string) {
	fmt.Println(outLine)
	fmt.Println(evernightMoments + " v" + evernightMomentsVersion)
	conf := LoadConfig()
	var plans []RenamePlan
	counter := 1
	i18n.SetLanguage(conf.Language)

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
	fmt.Println(i18n.T("要配置格式") + ": " + evernightMoments)
	fmt.Println(outLine)
	fmt.Println(i18n.T("正在分析") + "...")
	fmt.Println()

	addFileToPlan := func(path string) {
		info, err := os.Stat(path)
		if err != nil || info.IsDir() {
			return
		}

		t, source, err := GetPhotoTime(path)
		if err != nil {
			fmt.Printf("%s %s: %s\n", i18n.T("跳过"), i18n.T("文件错误"), path)
			return
		}

		rawTimeStr := t.Format("2006-01-02 15:04:05")
		newName := GenerateNewName(conf.Format, t, path, counter)
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

		fmt.Printf("[%d] %s: %s\n", len(plans), i18n.T("原文件"), absPath)
		fmt.Printf("-> %s %s : %s\n", i18n.T("依据"), source, rawTimeStr)
		fmt.Printf("-> %s: %s\n", i18n.T("新文件名"), newName)
		fmt.Println()
		counter++
	}

	for _, pattern := range cleanFiles {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			fmt.Printf("%s %s: %s\n", i18n.T("错误"), i18n.T("路径错误"), pattern)
			continue
		}

		for _, path := range matches {
			info, err := os.Stat(path)
			if err != nil {
				continue
			}

			if info.IsDir() {
				if recursive {
					filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
						if err == nil && !d.IsDir() {
							addFileToPlan(p)
						}
						return nil
					})
				} else {
					entries, err := os.ReadDir(path)
					if err == nil {
						for _, entry := range entries {
							if !entry.IsDir() {
								addFileToPlan(filepath.Join(path, entry.Name()))
							}
						}
					}
				}
			} else {
				// 如果是普通文件
				addFileToPlan(path)
			}
		}
	}

	// --- 后续逻辑（确认和执行）保持不变 ---
	if len(plans) == 0 {
		fmt.Println(i18n.T("没有文件"))
		return
	}

	proceed := true
	fmt.Println(outLine)
	if conf.Confirm {
		fmt.Printf("%s %s? (y/n): ", i18n.T("共计", len(plans)), i18n.T("确认"))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input != "y" && input != "yes" {
			proceed = false
		}
	}

	if proceed {
		fmt.Println(i18n.T("开始") + "...")
		fmt.Println()
		successCount := 0
		for i, p := range plans {
			fmt.Printf("[%d] %s: %s\n", i+1, i18n.T("原文件"), p.AbsPath)
			// ... 这里的 os.Rename 逻辑与你之前的一致 ...
			if p.OldPath == p.NewPath {
				fmt.Printf("-> %s: %s\n", i18n.T("跳过"), i18n.T("无变化"))
				fmt.Println()
				continue
			}
			if _, err := os.Stat(p.NewPath); err == nil {
				fmt.Println("-> " + i18n.T("重命名失败") + i18n.T("已存在"))
				fmt.Println()
				continue
			}
			err := os.Rename(p.OldPath, p.NewPath)
			if err != nil {
				fmt.Printf("-> %s: %v\n", i18n.T("重命名失败"), err)
			} else {
				fmt.Println("-> " + i18n.T("重命名成功"))
				successCount++
			}
			fmt.Println()
		}
		fmt.Println(i18n.T("处理结果", successCount, len(plans)-successCount, len(plans)))
	} else {
		fmt.Println(i18n.T("取消"))
	}
	if conf.EndPause {
		EndPause()
	}
}
