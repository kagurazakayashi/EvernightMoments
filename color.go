package main

// noColorOutput 全域旗標，決定是否停用 ANSI 彩色輸出
// 由 runRenameMode 在載入設定後依據 Config.NoColor 設定
var noColorOutput bool

// Dim 回傳暗色（dim）文字；若停用彩色則原樣回傳
func Dim(s string) string {
	if noColorOutput {
		return s
	}
	return "\033[2m" + s + "\033[0m"
}

// Bold 回傳粗體文字
func Bold(s string) string {
	if noColorOutput {
		return s
	}
	return "\033[1m" + s + "\033[0m"
}

// Cyan 回傳青色文字
func Cyan(s string) string {
	if noColorOutput {
		return s
	}
	return "\033[36m" + s + "\033[0m"
}

// Yellow 回傳黃色文字
func Yellow(s string) string {
	if noColorOutput {
		return s
	}
	return "\033[33m" + s + "\033[0m"
}

// Red 回傳紅色文字
func Red(s string) string {
	if noColorOutput {
		return s
	}
	return "\033[31m" + s + "\033[0m"
}

// Green 回傳綠色文字
func Green(s string) string {
	if noColorOutput {
		return s
	}
	return "\033[32m" + s + "\033[0m"
}
