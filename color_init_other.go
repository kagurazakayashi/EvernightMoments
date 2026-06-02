//go:build !windows

package main

// initConsoleColors 在非 Windows 平台上為空操作（ANSI 逸出碼原生支援）
func initConsoleColors() {}
