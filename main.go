//go:generate goversioninfo -o=resource_windows_386.syso -64=false -icon=ico/icon.ico -manifest=main.exe.manifest
//go:generate goversioninfo -o=resource_windows_amd64.syso -64=true -icon=ico/icon.ico -manifest=main.exe.manifest
//go:generate goversioninfo -o=resource_windows_arm64.syso -arm=true -icon=ico/icon.ico -manifest=main.exe.manifest
package main

import (
	"bufio"
	"fmt"
	"os"
)

const defaultFormat = "<YYYY><MM><DD>_<HH><mm><ss>_<*>"
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
