package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RenamePlan 存储重命名任务的细节
type RenamePlan struct {
	OldPath string
	NewPath string
	NewName string
	Source  string
	RawTime string
}

func runRenameMode(files []string) {
	conf := LoadConfig()
	var plans []RenamePlan
	counter := 1

	fmt.Println("正在准备重命名计划...")

	for _, pattern := range files {
		// 处理可能的通配符
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
				NewPath: newPath,
				NewName: newName,
				Source:  source,
				RawTime: rawTimeStr,
			})

			fmt.Printf("原文件: %s\n", absPath)
			fmt.Printf("-> 依据 %s : %s , 应用格式: %s\n", source, rawTimeStr, conf.Format)
			fmt.Printf("-> 新文件名: %s\n\n", newName)

			counter++
		}
	}

	if len(plans) == 0 {
		fmt.Println("未发现可处理的照片文件。")
		return
	}

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

	if proceed {
		fmt.Println("\n开始处理...")
		successCount := 0
		for _, p := range plans {
			if p.OldPath == p.NewPath {
				fmt.Printf("[忽略] 文件名未变化: %s\n", filepath.Base(p.OldPath))
				continue
			}

			if _, err := os.Stat(p.NewPath); err == nil {
				fmt.Printf("[跳过] 目标已存在: %s\n", p.NewName)
				continue
			}

			if err := os.Rename(p.OldPath, p.NewPath); err != nil {
				fmt.Printf("[失败] %s -> %v\n", filepath.Base(p.OldPath), err)
			} else {
				successCount++
			}
		}
		fmt.Printf("\n处理完成！成功: %d, 总计: %d\n", successCount, len(plans))
	} else {
		fmt.Println("\n已取消操作。")
	}
}
