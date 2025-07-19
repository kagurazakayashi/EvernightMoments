package main

import (
	"os"
)

const defaultFormat = "YYYYMMDD_HHmmss_*"
const outLine = "--------------------------------"
const evernightMoments = "EvernightMoments"
const evernightMomentsVersion = "1.0.0"

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
