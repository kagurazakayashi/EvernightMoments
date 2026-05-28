package main

import "testing"

// TestMatchIndex 驗證 MatchIndex 能將語系字串正確對應到 support 清單的索引
// support 順序：0=英文, 1=簡體中文, 2=繁體中文, 3=日文
func TestMatchIndex(t *testing.T) {
	mgr := NewI18nManager()

	cases := []struct {
		lang string
		want int
	}{
		{"en", 0},
		{"zh-Hans", 1},
		{"zh-Hant", 2},
		{"ja", 3},
	}

	for _, c := range cases {
		if got := mgr.MatchIndex(c.lang); got != c.want {
			t.Errorf("MatchIndex(%q) = %d, want %d", c.lang, got, c.want)
		}
	}
}
