package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// buttonFormItem 將 tview.Button 包裝成可放入表單欄位序列的 FormItem。
//
// 一般的 tview.Button 只能透過 Form.AddButton 放在表單底部按鈕列；此包裝讓按鈕
// 能以獨立一行插入欄位之間（例如緊接在「ExifTool 路徑」輸入框下方），並讓焦點
// 沿用表單原生的 Tab 巡覽機制（Enter 觸發、Tab/Backtab/Esc 切換焦點）。
type buttonFormItem struct {
	*tview.Button
	labelWidth int // 由 Form 透過 SetFormAttributes 傳入，用於將按鈕縮排對齊欄位輸入框
}

// 編譯期斷言：確保 buttonFormItem 完整實作 tview.FormItem 介面
var _ tview.FormItem = (*buttonFormItem)(nil)

// newButtonFormItem 建立一個包裝按鈕，label 為按鈕文字，selected 為點選時的回呼。
func newButtonFormItem(label string, selected func()) *buttonFormItem {
	return &buttonFormItem{Button: tview.NewButton(label).SetSelectedFunc(selected)}
}

// GetLabel 回傳空字串：按鈕本身即顯示文字，左側不需要額外的標籤欄。
func (b *buttonFormItem) GetLabel() string { return "" }

// SetFormAttributes 記錄表單統一的標籤欄寬度（供繪製時縮排對齊），色彩屬性對按鈕無作用。
func (b *buttonFormItem) SetFormAttributes(labelWidth int, _, _, _, _ tcell.Color) tview.FormItem {
	b.labelWidth = labelWidth
	return b
}

// GetFieldWidth 回傳按鈕所需寬度（文字寬度加上左右邊距）。
func (b *buttonFormItem) GetFieldWidth() int {
	return tview.TaggedStringWidth(b.Button.GetLabel()) + 4
}

// GetFieldHeight 回傳 1，使按鈕僅佔一行（避免套用多行欄位的預設高度）。
func (b *buttonFormItem) GetFieldHeight() int { return 1 }

// SetFinishedFunc 委派給按鈕的離開回呼，使 Tab/Backtab/Esc 能讓表單切換焦點。
func (b *buttonFormItem) SetFinishedFunc(handler func(key tcell.Key)) tview.FormItem {
	b.Button.SetExitFunc(handler)
	return b
}

// SetDisabled 委派給按鈕，並回傳 FormItem 介面型別以符合介面簽章。
func (b *buttonFormItem) SetDisabled(disabled bool) tview.FormItem {
	b.Button.SetDisabled(disabled)
	return b
}

// Draw 將按鈕繪製在標籤欄寬度之後的位置，使其左緣與上方輸入框對齊。
func (b *buttonFormItem) Draw(screen tcell.Screen) {
	x, y, width, height := b.GetRect()
	// 以表單統一的標籤欄寬度作為縮排
	indent := b.labelWidth
	if indent > width {
		indent = width
	}
	// 計算按鈕寬度，並避免超出可用範圍
	buttonWidth := tview.TaggedStringWidth(b.Button.GetLabel()) + 4
	if buttonWidth > width-indent {
		buttonWidth = width - indent
	}
	if buttonWidth < 0 {
		buttonWidth = 0
	}
	b.Button.SetRect(x+indent, y, buttonWidth, height)
	b.Button.Draw(screen)
}
