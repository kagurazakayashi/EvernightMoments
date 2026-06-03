package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// reservedNames 儲存作業系統（主要是 Windows）限制使用的保留檔案名稱
var reservedNames = make(map[string]bool)

// GenerateNewName 根據使用者定義的格式字串與檔案資訊，生成新的檔案名稱
// 參數說明：
//
//	format: 格式字串，例如 "<YYYY>-<MM>-<DD>_image_<#>"
//	t: 用於命名的時間基準（通常為照片拍攝時間）
//	originalPath: 原始檔案的完整路徑
//	index: 當前處理的檔案序號，用於 <#> 標籤
//	multiExt: 是否將多層副檔名視為檔名的一部分（true = 僅剝離最後一層，false = 剝離全部）
//
// 回傳值：
//
//	結合格式化後的名稱與原始副檔名的新檔名
func GenerateNewName(format string, t time.Time, originalPath string, index int, multiExt bool) string {
	// 取得原始檔名（含副檔名）
	originalName := filepath.Base(originalPath)
	// 取得最後一層副檔名（例如 .dop）
	ext := filepath.Ext(originalName)
	// 依 multiExt 設定決定 nameWithoutExt 的計算方式
	var nameWithoutExt string
	if multiExt {
		// 支援多重副檔名：將所有副檔名都視為副檔名，全部剝離
		// 例如 "KYS0001.ARW.dop" → "KYS0001"
		nameWithoutExt = stripAllExtensions(originalName)
	} else {
		// 僅最後一層視為副檔名，中間層保留為檔名的一部分
		// 例如 "KYS0001.ARW.dop" → "KYS0001.ARW"
		nameWithoutExt = strings.TrimSuffix(originalName, ext)
	}

	// 定義替換對映清單，將格式標籤映射至實際數值
	// 注意：長標籤（如 <YYYY>）必須排在短標籤（如 <YY>）前面，以避免部分匹配導致錯誤
	replacements := []struct {
		old string
		new string
	}{
		{"<YYYY>", fmt.Sprintf("%04d", t.Year())},   // 四位數年份
		{"<YY>", fmt.Sprintf("%02d", t.Year()%100)}, // 兩位數年份
		{"<MM>", fmt.Sprintf("%02d", t.Month())},    // 兩位數月份
		{"<M>", fmt.Sprintf("%d", t.Month())},       // 單/雙位數月份
		{"<DD>", fmt.Sprintf("%02d", t.Day())},      // 兩位數日期
		{"<D>", fmt.Sprintf("%d", t.Day())},         // 單/雙位數日期
		{"<HH>", fmt.Sprintf("%02d", t.Hour())},     // 兩位數小時（24小時制）
		{"<H>", fmt.Sprintf("%d", t.Hour())},        // 單/雙位數小時
		{"<mm>", fmt.Sprintf("%02d", t.Minute())},   // 兩位數分鐘
		{"<m>", fmt.Sprintf("%d", t.Minute())},      // 單/雙位數分鐘
		{"<ss>", fmt.Sprintf("%02d", t.Second())},   // 兩位數秒數
		{"<s>", fmt.Sprintf("%d", t.Second())},      // 單/雙位數秒數
		{"<##>", fmt.Sprintf("%02d", index)},        // 兩位數序號（不足補零）
		{"<#>", fmt.Sprintf("%d", index)},           // 原始序號
		{"<*>", nameWithoutExt},                     // 原始檔名（不含副檔名）
	}

	result := format

	// 依序執行字串替換
	for _, r := range replacements {
		result = strings.ReplaceAll(result, r.old, r.new)
	}

	// 重新拼接副檔名並回傳
	return result + ext
}

// stripAllExtensions 從檔名中逐層剝離所有副檔名，僅回傳不含任何副檔名的基底名稱
// 例如 "KYS0001.ARW.dop" → "KYS0001"；"photo.jpg" → "photo"；"README" → "README"
func stripAllExtensions(filename string) string {
	for {
		ext := filepath.Ext(filename)
		if ext == "" {
			return filename
		}
		filename = strings.TrimSuffix(filename, ext)
	}
}

// InitInvalidChars 初始化保留名稱清單
// 這些名稱在 Windows 系統中具有特殊用途（如設備名稱），不能作為一般檔案名稱使用
func InitInvalidChars() {
	// 基本保留名稱
	base := []string{"CON", "PRN", "AUX", "NUL"}
	for _, n := range base {
		reservedNames[n] = true
	}
	// 通訊埠與印表機連接埠相關保留名稱（COM1-9, LPT1-9）
	for i := 1; i <= 9; i++ {
		reservedNames[fmt.Sprintf("COM%d", i)] = true
		reservedNames[fmt.Sprintf("LPT%d", i)] = true
	}
}

// ContainsInvalidChars 檢查字串是否包含檔案系統不允許的非法字元或保留名稱
// 回傳值：
//
//	bool: 是否包含非法內容
//	string: 觸發錯誤的字元或名稱
func ContainsInvalidChars(s string) (bool, string) {
	// 空字串不視為非法（由其他邏輯判斷是否必填）
	if s == "" {
		return false, ""
	}

	// Windows 系統禁用的字元
	invalidChars := "\\/:?\"|"
	for _, char := range s {
		if strings.ContainsRune(invalidChars, char) {
			return true, string(char)
		}
	}

	// 檔案名稱末尾不可為句點
	if strings.HasSuffix(s, ".") {
		return true, "."
	}

	// 轉換為大寫以進行保留名稱比對（不分大小寫）
	upperS := strings.ToUpper(s)
	if reservedNames[upperS] {
		return true, s
	}

	// 通過所有檢查，回傳合法
	return false, ""
}

// PadNumberByReference 根據 ref 的位數給 val 前面補 0
func PadNumberByReference(val int, ref int) string {
	// 1. 處理負數情況：通常位數判斷基於絕對值
	absRef := ref
	if ref < 0 {
		absRef = -ref
	}

	// 2. 將第二個 int 轉為字串，獲取其長度（即位數）
	width := len(strconv.Itoa(absRef))

	// 3. 使用 fmt.Sprintf 進行格式化
	// %0*d 中，* 是一個佔位符，表示寬度由後面的引數 width 決定
	// 0 表示長度不足時在前面補 0
	return fmt.Sprintf("%0*d", width, val)
}
