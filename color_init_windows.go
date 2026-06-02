//go:build windows

package main

import "golang.org/x/sys/windows"

// initConsoleColors 在 Windows 上啟用虛擬終端機處理以支援 ANSI 逸出碼
// 若控制台不支援則靜默忽略
func initConsoleColors() {
	var mode uint32
	handle := windows.Stdout
	if err := windows.GetConsoleMode(handle, &mode); err == nil {
		mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
		windows.SetConsoleMode(handle, mode)
	}
}
