//go:generate goversioninfo -o=resource_windows_386.syso -64=false -icon=ico/icon.ico -manifest=main.exe.manifest
//go:generate goversioninfo -o=resource_windows_amd64.syso -64=true -icon=ico/icon.ico -manifest=main.exe.manifest
//go:generate goversioninfo -o=resource_windows_arm64.syso -arm=true -icon=ico/icon.ico -manifest=main.exe.manifest

package main

import (
	"bufio"
	"fmt"
	"os"
)

// 定義程式中使用的全域常數
const (
	// defaultFormat 定義預設的檔案命名格式範本
	defaultFormat = "<YYYY><MM><DD>_<HH><mm><ss>_<*>"
	// outLine 用於在命令列介面顯示分割線，提升可讀性
	outLine = "--------------------------------"
	// evernightMoments 程式名稱
	evernightMoments = "EvernightMoments"
	// evernightMomentsVersion 程式版本號
	evernightMomentsVersion = "1.2.0"
)

// main 是程式的進入點
func main() {
	// 取得命令列參數（排除執行檔路徑本身）
	args := os.Args[1:]

	// 若無參數傳入，則進入「設定模式」（Config Mode）
	if len(args) == 0 {
		runConfigMode()
		return
	}

	// 若有參數（通常是拖放檔案或命令列傳入路徑），則進入「更名模式」（Rename Mode）
	runRenameMode(args)
}

// EndPause 負責在程式執行完畢前暫停介面
// 這在 Windows 系統中特別重要，可以防止命令列視窗在執行完後立刻自動關閉
func EndPause() {
	// 顯示翻譯後的「回車退出」（按下 Enter 離開）提示訊息
	fmt.Print(i18n.T("回车退出") + "...")

	// 建立讀取器以捕捉標準輸入
	reader := bufio.NewReader(os.Stdin)
	// 持續等待直到使用者按下 Enter 鍵（換行符號）
	reader.ReadString('\n')
}
