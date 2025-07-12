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
	conf := LoadConfig()
	var plans []RenamePlan
	counter := 1

	fmt.Printf("命名格式: %s\n正在准备...\n\n", conf.Format)

	for _, pattern := range files {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			fmt.Printf("![错误] 无法解析路径: %s\n", pattern)
			continue
		}

		for _, path := range matches {
			info, err := os.Stat(path)
			if err != nil || info.IsDir() {
				continue
			}

			t, source, err := GetPhotoTime(path)
			if err != nil {
				fmt.Printf("![跳过] 无法读取文件信息: %s\n", path)
				continue
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

			fmt.Printf("原文件: %s\n", absPath)
			fmt.Printf("-> 依据 %s : %s\n", source, rawTimeStr)
			fmt.Printf("-> 新文件名: %s\n", newName)
			fmt.Println()

			counter++
		}
	}

	if len(plans) == 0 {
		fmt.Println("未发现可处理的照片文件。")
		return
	}

	// 确认逻辑
	proceed := true
	if conf.Confirm {
		fmt.Printf("共计 %d 个文件，确认执行重命名? (y/n): ", len(plans))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input != "y" && input != "yes" {
			proceed = false
		}
	}

	// 正式执行重命名
	if proceed {
		fmt.Print("开始处理...\n\n")
		successCount := 0
		for _, p := range plans {
			fmt.Printf("原文件: %s\n", p.AbsPath)
			fmt.Printf("-> 依据 %s : %s\n", p.Source, p.RawTime)
			fmt.Printf("-> 新文件名: %s\n", p.NewName)

			if p.OldPath == p.NewPath {
				fmt.Println("-> 跳过: 文件名未发生变化")
				fmt.Println()
				continue
			}

			// 检查目标是否存在
			if _, err := os.Stat(p.NewPath); err == nil {
				fmt.Println("-> 重命名失败！错误原因: 目标文件已存在")
				fmt.Println()
				continue
			}

			// 执行重命名
			err := os.Rename(p.OldPath, p.NewPath)
			if err != nil {
				fmt.Printf("-> 重命名失败！错误原因: %v\n", err)
			} else {
				fmt.Println("-> 重命名成功。")
				successCount++
			}
			fmt.Println()
		}
		fmt.Printf("处理完成！成功: %d, 失败: %d, 总计: %d\n", successCount, len(plans)-successCount, len(plans))
	} else {
		fmt.Println("已取消操作。")
	}
}
