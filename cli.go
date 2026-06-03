package main

import (
	"flag"
)

// CLIFlags 儲存從命令列解析出的設定覆蓋值
// 每個欄位為 nil 表示使用者未在命令列指定該選項，應使用設定檔的值
type CLIFlags struct {
	Language     *string  // 介面語言覆蓋值，nil 表示未指定
	Format       *string  // 命名格式覆蓋值，nil 表示未指定
	Exclude      []string // 排除樣式清單（可多次指定）
	Sync         []string // 同步樣式清單（可多次指定）
	Confirm      *bool    // 預覽確認覆蓋值，nil 表示未指定
	EndPause     *bool    // 結束暫停覆蓋值，nil 表示未指定
	ExiftoolPath *string  // ExifTool 路徑覆蓋值，nil 表示未指定
	MultiExt     *bool    // 多層副檔名覆蓋值，nil 表示未指定
	NoColor      *bool    // 停用彩色輸出覆蓋值，nil 表示未指定
	Recursive    bool     // 是否遞迴處理子目錄
}

// parseCLIFlags 定義命令列旗標並解析傳入的參數
// 回傳解析後的 CLI 選項結構與剩餘的位置參數（檔案/目錄路徑）
func parseCLIFlags(args []string) (*CLIFlags, []string) {
	f := flag.NewFlagSet(evernightMoments, flag.ContinueOnError)
	var opts CLIFlags

	// -- 語言設定 --
	// -l / --language：指定介面語言，接受 "en"、"zh-Hans"、"zh-Hant"、"ja"
	f.Func("l", "Set display language (en/zh-Hans/zh-Hant/ja)", func(s string) error {
		v := s
		opts.Language = &v
		return nil
	})
	f.Func("language", "Set display language", func(s string) error {
		v := s
		opts.Language = &v
		return nil
	})

	// -- 命名格式 --
	// -f / --format：指定重新命名格式範本
	f.Func("f", "Set renaming format template", func(s string) error {
		v := s
		opts.Format = &v
		return nil
	})
	f.Func("format", "Set renaming format template", func(s string) error {
		v := s
		opts.Format = &v
		return nil
	})

	// -- 排除樣式（可重複指定） --
	// -e / --exclude：每次指定將累加到排除清單中
	f.Func("e", "Add an exclude glob pattern (can be repeated)", func(s string) error {
		opts.Exclude = append(opts.Exclude, s)
		return nil
	})
	f.Func("exclude", "Add an exclude glob pattern (can be repeated)", func(s string) error {
		opts.Exclude = append(opts.Exclude, s)
		return nil
	})

	// -- 同步樣式（可重複指定） --
	// -s / --sync：每次指定將累加到同步清單中
	f.Func("s", "Add a sync glob pattern (can be repeated)", func(s string) error {
		opts.Sync = append(opts.Sync, s)
		return nil
	})
	f.Func("sync", "Add a sync glob pattern (can be repeated)", func(s string) error {
		opts.Sync = append(opts.Sync, s)
		return nil
	})

	// -- 預覽確認（啟用） --
	// -y / --confirm：強制啟用重新命名前的預覽確認
	f.BoolFunc("y", "Enable preview confirmation before renaming", func(s string) error {
		v := true
		opts.Confirm = &v
		return nil
	})
	f.BoolFunc("confirm", "Enable preview confirmation before renaming", func(s string) error {
		v := true
		opts.Confirm = &v
		return nil
	})

	// -- 預覽確認（停用） --
	// -ny / --no-confirm：強制停用重新命名前的預覽確認
	f.BoolFunc("ny", "Disable preview confirmation before renaming", func(s string) error {
		v := false
		opts.Confirm = &v
		return nil
	})
	f.BoolFunc("no-confirm", "Disable preview confirmation before renaming", func(s string) error {
		v := false
		opts.Confirm = &v
		return nil
	})

	// -- 結束暫停（啟用） --
	// -p / --pause：強制啟用程式結束後的等待暫停
	f.BoolFunc("p", "Enable pause before exit", func(s string) error {
		v := true
		opts.EndPause = &v
		return nil
	})
	f.BoolFunc("pause", "Enable pause before exit", func(s string) error {
		v := true
		opts.EndPause = &v
		return nil
	})

	// -- 結束暫停（停用） --
	// -np / --no-pause：強制停用程式結束後的等待暫停
	f.BoolFunc("np", "Disable pause before exit", func(s string) error {
		v := false
		opts.EndPause = &v
		return nil
	})
	f.BoolFunc("no-pause", "Disable pause before exit", func(s string) error {
		v := false
		opts.EndPause = &v
		return nil
	})

	// -- ExifTool 路徑 --
	// -x / --exiftool-path：指定 ExifTool 執行檔路徑；傳入空字串可停用 ExifTool
	f.Func("x", "Set ExifTool executable path (empty to disable)", func(s string) error {
		v := s
		opts.ExiftoolPath = &v
		return nil
	})
	f.Func("exiftool-path", "Set ExifTool executable path (empty to disable)", func(s string) error {
		v := s
		opts.ExiftoolPath = &v
		return nil
	})

	// -- 遞迴處理 --
	// -r / --recursive：處理目錄時包含子目錄
	f.BoolVar(&opts.Recursive, "r", false, "Process subdirectories recursively")
	f.BoolVar(&opts.Recursive, "recursive", false, "Process subdirectories recursively")

	// -- 多重副檔名（啟用） --
	// -m / --multi-ext：支援多重副檔名（剝離所有副檔名）
	f.BoolFunc("m", "Support multi-level extensions (strip all extensions)", func(s string) error {
		v := true
		opts.MultiExt = &v
		return nil
	})
	f.BoolFunc("multi-ext", "Support multi-level extensions (strip all extensions)", func(s string) error {
		v := true
		opts.MultiExt = &v
		return nil
	})

	// -- 多重副檔名（停用） --
	// -nm / --no-multi-ext：不支援多重副檔名（僅剝離最後一層副檔名）
	f.BoolFunc("nm", "Disable multi-level extension support (strip only last)", func(s string) error {
		v := false
		opts.MultiExt = &v
		return nil
	})
	f.BoolFunc("no-multi-ext", "Disable multi-level extension support (strip only last)", func(s string) error {
		v := false
		opts.MultiExt = &v
		return nil
	})

	// -- 停用彩色輸出 --
	// -nc / --no-color：停用 ANSI 彩色輸出
	f.BoolFunc("nc", "Disable colored output", func(s string) error {
		v := true
		opts.NoColor = &v
		return nil
	})
	f.BoolFunc("no-color", "Disable colored output", func(s string) error {
		v := true
		opts.NoColor = &v
		return nil
	})

	f.Parse(args)
	return &opts, f.Args()
}

// applyCLIOverrides 將 CLI 覆蓋值套用到設定結構
// 僅覆蓋使用者明確指定的欄位，未指定的欄位保留設定檔原始值
func applyCLIOverrides(conf *Config, opts *CLIFlags) {
	if opts.Language != nil {
		conf.Language = *opts.Language
	}
	if opts.Format != nil {
		conf.Format = *opts.Format
	}
	if len(opts.Exclude) > 0 {
		conf.Exclude = opts.Exclude
	}
	if len(opts.Sync) > 0 {
		conf.Sync = opts.Sync
	}
	if opts.Confirm != nil {
		conf.Confirm = *opts.Confirm
	}
	if opts.EndPause != nil {
		conf.EndPause = *opts.EndPause
	}
	if opts.ExiftoolPath != nil {
		conf.ExiftoolPath = opts.ExiftoolPath
	}
	if opts.MultiExt != nil {
		conf.MultiExt = *opts.MultiExt
	}
	if opts.NoColor != nil {
		conf.NoColor = *opts.NoColor
	}
}
