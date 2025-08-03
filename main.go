package main

import (
	"bufio"
	"fmt"
	"os"
)

const defaultFormat = "YYYYMMDD_HHmmss_*"
const outLine = "--------------------------------"
const evernightMoments = "EvernightMoments"
const evernightMomentsVersion = "1.0.0"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		runConfigMode()
		return
	}
	runRenameMode(args)
}

func EndPause() {
	fmt.Print(i18n.T("回车退出") + "...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}
