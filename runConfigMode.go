package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// i18n 是全域的多國語言管理器執行個體，負責處理介面文字的翻譯
var i18n = NewI18nManager()

// isInteractiveTerminal 判斷目前的標準輸入與輸出是否皆連接到互動式終端機（TTY）
// 當程式以管線、重新導向或無 TTY 的環境執行時，TUI 無法正常運作，需提前回退處理
func isInteractiveTerminal() bool {
	// 檢查標準輸入是否為字元裝置（終端機）
	if fi, err := os.Stdin.Stat(); err != nil || (fi.Mode()&os.ModeCharDevice) == 0 {
		return false
	}
	// 檢查標準輸出是否為字元裝置（終端機）
	if fo, err := os.Stdout.Stat(); err != nil || (fo.Mode()&os.ModeCharDevice) == 0 {
		return false
	}
	return true
}

// runConfigMode 啟動以 tview 實作的全螢幕 TUI 設定介面，引導使用者配置程式參數
func runConfigMode() {
	// 非互動式終端機環境下無法繪製 TUI，直接輸出錯誤訊息並結束
	if !isInteractiveTerminal() {
		fmt.Fprintln(os.Stderr, evernightMoments+": configuration mode requires an interactive terminal (TTY).")
		os.Exit(1)
	}

	// 初始化檔名保留字清單，供格式合法性檢查使用
	InitInvalidChars()

	// 載入目前的設定檔，若無設定檔則取得預設值
	conf := LoadConfig()
	// 若設定檔中未指定語言，則嘗試讀取作業系統的語系
	if conf.Language == "" {
		conf.Language = i18n.GetSystemLanguage()
	}
	// 套用目前語系，讓後續介面文字以正確語言顯示
	i18n.SetLanguage(conf.Language)

	// 用於記錄 TUI 結束後的儲存結果，供離開後於命令列輸出
	var saved bool
	var saveErr error
	var savedPath string

	// 建立 tview 應用程式執行個體
	app := tview.NewApplication()

	// rebuild 會以目前語系重新建立整個介面並設為根節點
	// 切換語言時呼叫此函式即可達成介面文字的即時重繪
	var rebuild func()
	rebuild = func() {
		app.SetRoot(buildConfigUI(app, &conf, &saved, &saveErr, &savedPath, rebuild), true)
	}
	rebuild()

	// 進入 TUI 主事件迴圈，直到使用者儲存或離開
	if err := app.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// TUI 結束後，依儲存結果於命令列輸出對應訊息
	if saved {
		if saveErr != nil {
			fmt.Println(i18n.T("保存配置失败", savedPath, saveErr))
		} else {
			fmt.Println(i18n.T("保存配置成功") + savedPath)
		}
		// 若使用者開啟了結束等待功能，則呼叫 EndPause 避免視窗閃退
		if conf.EndPause {
			EndPause()
		}
	}
}

// buildConfigUI 依據目前語系與設定內容建立完整的設定介面版面，並回傳可作為根節點的元件
// 參數說明：
//
//	app:        tview 應用程式執行個體
//	conf:       目前設定（以指標傳入，欄位變更會即時寫回）
//	saved:      指標，標記是否經由「儲存」按鈕結束
//	saveErr:    指標，記錄儲存過程發生的錯誤
//	savedPath:  指標，記錄設定檔實際寫入的路徑
//	rebuild:    重建整個介面的回呼（用於語言即時切換）
func buildConfigUI(
	app *tview.Application,
	conf *Config,
	saved *bool,
	saveErr *error,
	savedPath *string,
	rebuild func(),
) tview.Primitive {
	// 語言選項的顯示名稱（固定不翻譯，順序需與 i18n.support 對應）
	langNames := []string{"English", "简体中文", "繁體中文", "日本語"}
	// 依目前設定語系定位下拉選單的初始索引
	curLangIndex := i18n.MatchIndex(conf.Language)
	if curLangIndex < 0 || curLangIndex >= len(langNames) {
		curLangIndex = 0
	}

	// 各欄位的簡短標籤文字（詳細說明改放於底部資訊區，避免標籤過長擠壓輸入框）
	lblLang := i18n.T("语言")
	lblFormat := i18n.T("格式")
	lblExclude := i18n.T("排除")
	lblSync := i18n.T("同步")
	lblExif := i18n.T("Exif路径")
	lblConfirm := i18n.T("预览确认")
	lblEndPause := i18n.T("等待退出")
	lblMultiExt := i18n.T("多层副档名")

	// 計算所有標籤的最大顯示寬度，供對齊填充使用（考量中日文字佔兩格寬）
	maxLabelWidth := 0
	for _, l := range []string{lblLang, lblFormat, lblExclude, lblSync, lblExif, lblConfirm, lblEndPause, lblMultiExt} {
		if w := tview.TaggedStringWidth(l); w > maxLabelWidth {
			maxLabelWidth = w
		}
	}
	// pad 將標籤右側補空白至統一寬度，使各欄位的輸入框起點對齊
	pad := func(label string) string {
		gap := maxLabelWidth - tview.TaggedStringWidth(label)
		if gap < 0 {
			gap = 0
		}
		return label + strings.Repeat(" ", gap) + "  "
	}

	// 預先宣告各表單欄位變數，供下方各閉包延遲取值
	var formatField, excludeField, syncField, exiftoolField *tview.InputField
	var confirmCheck, endPauseCheck, multiExtCheck *tview.Checkbox

	// 底部資訊區，顯示欄位說明、即時預覽與操作提示
	footer := tview.NewTextView().SetDynamicColors(true).SetWordWrap(true)
	// 取得目前語系的操作提示文字
	hints := i18n.T("操作提示")

	// refreshFooter 以「說明區塊 + 額外資訊 + 操作提示」組合並更新底部資訊區
	refreshFooter := func(help string, extra string) {
		var b strings.Builder
		if help != "" {
			b.WriteString(help)
		}
		if extra != "" {
			if b.Len() > 0 {
				b.WriteString("\n\n")
			}
			b.WriteString(extra)
		}
		if b.Len() > 0 {
			b.WriteString("\n\n")
		}
		b.WriteString("[gray]")
		b.WriteString(hints)
		b.WriteString("[-]")
		footer.SetText(b.String())
	}

	// formatExtra 依命名格式欄位目前的輸入內容，產生即時範例或非法字元警告
	formatExtra := func() string {
		txt := formatField.GetText()
		if invalid, ch := ContainsInvalidChars(txt); invalid {
			return "[red]" + i18n.T("非法字符格式", ch) + "[-]"
		}
		if txt != "" {
			example := GenerateNewName(txt, time.Now(), "Photo.jpg", 1, conf.MultiExt)
			return i18n.T("示例输出") + ": [green]" + example + "[-]"
		}
		return ""
	}

	// syncFormToConf 將目前各欄位的輸入值同步寫回設定物件
	syncFormToConf := func() {
		conf.Format = strings.TrimSpace(formatField.GetText())
		conf.Exclude = parsePatterns(excludeField.GetText())
		conf.Sync = parsePatterns(syncField.GetText())
		// ExifTool 路徑以指標保存：留空代表使用者明確停用（僅內建解析）
		ep := strings.TrimSpace(exiftoolField.GetText())
		conf.ExiftoolPath = &ep
		conf.Confirm = confirmCheck.IsChecked()
		conf.EndPause = endPauseCheck.IsChecked()
		conf.MultiExt = multiExtCheck.IsChecked()
	}

	// --- 語言下拉選單 ---
	langDD := tview.NewDropDown().
		SetLabel(pad(lblLang)).
		SetOptions(langNames, nil)
	// 先設定目前選項（此時尚未綁定 selected 回呼，不會觸發切換邏輯）
	langDD.SetCurrentOption(curLangIndex)
	// 再綁定選擇回呼：保留現有輸入、切換語系並即時重繪整個介面
	langDD.SetSelectedFunc(func(_ string, index int) {
		syncFormToConf()
		if index >= 0 && index < len(i18n.support) {
			conf.Language = i18n.support[index].String()
		}
		i18n.SetLanguage(conf.Language)
		rebuild()
	})
	langDD.SetFocusFunc(func() { refreshFooter("", "") })

	// --- 命名格式輸入欄（欄寬設為 0 表示自動填滿剩餘寬度，避免被長標籤擠出）---
	formatField = tview.NewInputField().
		SetLabel(pad(lblFormat)).
		SetText(conf.Format).
		SetFieldWidth(0)
	// 即時阻擋輸入作業系統不允許的非法字元（變數括號 < > 為合法字元，不在此列）
	formatField.SetAcceptanceFunc(func(_ string, lastChar rune) bool {
		return !strings.ContainsRune(`\/:?"|`, lastChar)
	})
	formatField.SetChangedFunc(func(_ string) {
		if formatField.HasFocus() {
			refreshFooter(i18n.T("可用变量"), formatExtra())
		}
	})
	formatField.SetFocusFunc(func() {
		refreshFooter(i18n.T("可用变量"), formatExtra())
	})

	// --- 排除項輸入欄 ---
	excludeField = tview.NewInputField().
		SetLabel(pad(lblExclude)).
		SetText(strings.Join(conf.Exclude, ", ")).
		SetFieldWidth(0)
	excludeField.SetFocusFunc(func() { refreshFooter(i18n.T("排除项说明"), "") })

	// --- 同步項輸入欄 ---
	syncField = tview.NewInputField().
		SetLabel(pad(lblSync)).
		SetText(strings.Join(conf.Sync, ", ")).
		SetFieldWidth(0)
	syncField.SetFocusFunc(func() { refreshFooter(i18n.T("同步项说明"), "") })

	// --- ExifTool 路徑輸入欄 ---
	// 決定初始值：已有設定則沿用（含明確留空），否則自動偵測填入
	exiftoolInitial := ""
	if conf.ExiftoolPath != nil {
		exiftoolInitial = *conf.ExiftoolPath
	} else {
		exiftoolInitial = DetectExiftoolPath()
	}
	exiftoolField = tview.NewInputField().
		SetLabel(pad(lblExif)).
		SetText(exiftoolInitial).
		SetFieldWidth(0)
	exiftoolField.SetFocusFunc(func() { refreshFooter(i18n.T("Exif说明"), "") })

	// --- 確認預覽開關 ---
	confirmCheck = tview.NewCheckbox().
		SetLabel(pad(lblConfirm)).
		SetChecked(conf.Confirm)
	confirmCheck.SetFocusFunc(func() { refreshFooter(i18n.T("询问预览"), "") })

	// --- 結束等待開關 ---
	endPauseCheck = tview.NewCheckbox().
		SetLabel(pad(lblEndPause)).
		SetChecked(conf.EndPause)
	endPauseCheck.SetFocusFunc(func() { refreshFooter(i18n.T("结束等待"), "") })

	// --- 多層副檔名開關 ---
	multiExtCheck = tview.NewCheckbox().
		SetLabel(pad(lblMultiExt)).
		SetChecked(conf.MultiExt)
	multiExtCheck.SetFocusFunc(func() { refreshFooter(i18n.T("多层副档名说明"), "") })

	// 將所有欄位依序加入表單
	form := tview.NewForm()
	form.AddFormItem(langDD)
	form.AddFormItem(formatField)
	form.AddFormItem(excludeField)
	form.AddFormItem(syncField)
	form.AddFormItem(exiftoolField)
	// 「自動偵測」按鈕以獨立一行緊接在 ExifTool 路徑欄下方
	// （透過 buttonFormItem 包裝放入欄位序列，使按鈕緊鄰其對應的輸入框）
	detectBtn := newButtonFormItem(i18n.T("自动侦测"), func() {
		path := DetectExiftoolPath()
		exiftoolField.SetText(path)
		if path == "" {
			refreshFooter("[yellow]"+i18n.T("Exif未找到")+"[-]\n\n"+i18n.T("Exif说明"), "")
		} else {
			refreshFooter("[green]"+i18n.T("Exif已找到")+"[-]: "+path+"\n\n"+i18n.T("Exif说明"), "")
		}
	})
	detectBtn.SetFocusFunc(func() { refreshFooter(i18n.T("Exif说明"), "") })
	form.AddFormItem(detectBtn)
	form.AddFormItem(confirmCheck)
	form.AddFormItem(endPauseCheck)
	form.AddFormItem(multiExtCheck)

	// 「儲存並退出」按鈕：同步輸入、檢查格式合法性後寫入設定檔
	form.AddButton(i18n.T("保存"), func() {
		syncFormToConf()
		// 儲存前再次完整檢查命名格式（涵蓋結尾句點與保留字等情況）
		if invalid, ch := ContainsInvalidChars(conf.Format); invalid {
			refreshFooter("[red]"+i18n.T("非法字符格式", ch)+"[-]\n"+i18n.T("非法字符"), "")
			app.SetFocus(formatField)
			return
		}
		// 格式為空時回退至預設格式
		if conf.Format == "" {
			conf.Format = defaultFormat
		}
		path := getConfigPath()
		err := SaveConfig(*conf, path)
		*saved = true
		*saveErr = err
		*savedPath = path
		app.Stop()
	})
	// 「退出不保存」按鈕：直接結束 TUI，不寫入任何變更
	form.AddButton(i18n.T("退出"), func() {
		*saved = false
		app.Stop()
	})

	form.SetBorder(true).SetTitle(" " + i18n.T("软件配置") + " ")

	// 頂部標題區：軟體名稱、版本與介紹語合併為單行顯示
	header := tview.NewTextView().SetDynamicColors(true)
	header.SetText("[yellow::b]" + evernightMoments + "[-:-:-] v" + evernightMomentsVersion + "  " + i18n.T("介绍1"))

	footer.SetBorder(true)
	// 初始化底部資訊區內容
	refreshFooter("", "")

	// 以垂直方向排列：單行標題、表單、底部資訊區
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 1, 0, false).
		AddItem(form, 0, 1, true).
		AddItem(footer, 10, 0, false)

	// 全域按鍵攔截：按下 Esc 視為「退出不保存」
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			*saved = false
			app.Stop()
			return nil
		}
		return event
	})

	return flex
}
