# EvernightMoments — AI Agent Skill

## Description

EvernightMoments is a CLI photo renaming tool written in Go. It extracts original capture timestamps from photo EXIF metadata and renames files to a configurable time‑based format. It supports three time sources (priority order): ExifTool external tool → goexif built‑in parser → filesystem modification time. A full‑screen TUI (tview) configuration mode is also available.

## Entry Point

`main.go:26` — `func main()`

- **No arguments** → launches TUI configuration mode (`runConfigMode()`)
- **Arguments** → parses CLI flags and positional file paths, then enters rename mode (`runRenameMode()`)

## Build

```bash
go build -o EvernightMoments .
```

Go version ≥ 1.26.0 required. Cross‑platform build scripts: `build.bat` / `build.sh` (output to `bin/`).

## Tests

```bash
go test ./...
```

Test files: `config_test.go` (glob matching, parsePatterns), `language_test.go` (I18N matching).

## CLI Interface

```
EvernightMoments [flags] <file_or_folder_path...>
```

File/folder paths are positional arguments placed at the end. Multiple paths and wildcards are supported. Flags temporarily override config‑file values.

### Flags

| Short | Long             | Type       | Config Key      | Description                                     |
|-------|------------------|------------|-----------------|-------------------------------------------------|
| `-l`  | `--language`     | `string`   | `language`      | UI language: `en`, `zh-Hans`, `zh-Hant`, `ja`   |
| `-f`  | `--format`       | `string`   | `format`        | Renaming format template (see placeholders)     |
| `-e`  | `--exclude`      | `string[]` | `exclude`       | Glob pattern to exclude (repeatable)            |
| `-s`  | `--sync`         | `string[]` | `sync`          | Glob pattern for companion files (repeatable)   |
| `-y`  | `--confirm`      | `bool`     | `confirm`       | Enable preview confirmation                     |
| `-ny` | `--no-confirm`   | `bool`     | `confirm`       | Disable preview confirmation                    |
| `-p`  | `--pause`        | `bool`     | `endpause`      | Enable end‑of‑run pause (Windows safety)        |
| `-np` | `--no-pause`     | `bool`     | `endpause`      | Disable end‑of‑run pause                        |
| `-x`  | `--exiftool`     | `string`   | `exiftoolpath`  | Path to ExifTool executable (empty = disable)   |
| `-r`  | `--recursive`    | `bool`     | —               | Recurse into subdirectories                     |

`-h` / `--help` prints the full flag list.

### Flag Parsing Logic

`cli.go:22` — `parseCLIFlags()` uses `flag.NewFlagSet` with `flag.ContinueOnError`. String flags use `flag.Func` to set `*string` pointers (nil = not provided). Bool flags use `flag.BoolFunc` with separate on/off flags to distinguish "not set" from "explicitly false". `-e`/`-s` accumulate into slices.

`cli.go:150` — `applyCLIOverrides()` merges overrides into the `Config` struct; only fields with non‑nil pointers are applied.

## Configuration

### File Location

| OS      | Path                                                    |
|---------|---------------------------------------------------------|
| Windows | `%APPDATA%\EvernightMoments\config.json`                |
| macOS   | `~/Library/Application Support/EvernightMoments/config.json` |
| Linux   | `~/.config/EvernightMoments/config.json`                |

Legacy config (same directory as executable, named `<exe>.json`) is auto‑migrated.

### JSON Schema (`config.go:12`)

```json
{
  "language": "zh-Hans",
  "format": "<YYYY><MM><DD>_<HH><mm><ss>_<*>",
  "exclude": ["*.dop", "CaptureOne\\Settings153\\*.cos"],
  "sync": ["*.dop"],
  "confirm": true,
  "endpause": true,
  "exiftoolpath": "C:\\Tools\\exiftool.exe"
}
```

| Key            | Type       | Default                                         | Semantics                                                        |
|----------------|------------|-------------------------------------------------|------------------------------------------------------------------|
| `language`     | `string`   | `""` (auto‑detect)                              | Empty = system locale detection                                  |
| `format`       | `string`   | `"<YYYY><MM><DD>_<HH><mm><ss>_<*>"`            | Renaming template; empty falls back to default                   |
| `exclude`      | `[]string` | `null` (none)                                   | Glob patterns for files to skip                                  |
| `sync`         | `[]string` | `null` (none)                                   | Glob patterns for companion files renamed alongside primary      |
| `confirm`      | `bool`     | `true`                                          | Show preview and ask y/N before renaming                         |
| `endpause`     | `bool`     | `true`                                          | Wait for Enter before exit (prevents Windows console auto‑close) |
| `exiftoolpath` | `*string`  | `null` (auto‑detect on next run if unset)       | `null` = auto‑detect, `""` = disable ExifTool, `"path"` = use    |

## Format Placeholders

Defined in `format.go`. Case‑sensitive.

| Placeholder | Meaning              | Example Output |
|-------------|----------------------|----------------|
| `<YYYY>`    | 4‑digit year         | `2025`         |
| `<YY>`      | 2‑digit year         | `25`           |
| `<MM>`      | 2‑digit month        | `05`           |
| `<M>`       | month (no padding)   | `5`            |
| `<DD>`      | 2‑digit day          | `02`           |
| `<D>`       | day (no padding)     | `2`            |
| `<HH>`      | 2‑digit hour         | `09`           |
| `<H>`       | hour (no padding)    | `9`            |
| `<mm>`      | 2‑digit minute       | `07`           |
| `<m>`       | minute (no padding)  | `7`            |
| `<ss>`      | 2‑digit second       | `03`           |
| `<s>`       | second (no padding)  | `3`            |
| `<#>`       | sequence number      | `1`            |
| `<##>`      | zero‑padded sequence | `01`           |
| `<*>`       | original filename    | `photo.jpg`    |

Default format: `"<YYYY><MM><DD>_<HH><mm><ss>_<*>"`

## Architecture

| File                  | Role                                                    |
|-----------------------|---------------------------------------------------------|
| `main.go`             | Entry point, constants, `EndPause()` helper             |
| `cli.go`              | CLI flag definitions, parsing, override application     |
| `config.go`           | `Config` struct, JSON load/save, path matching, legacy migration |
| `format.go`           | Filename generation, illegal‑char & reserved‑name checks |
| `meta.go`             | EXIF time extraction (ExifTool → goexif → modtime), ExifTool auto‑detect |
| `runConfigMode.go`    | TUI configuration form (tview)                          |
| `runRenameMode.go`    | Rename workflow: scan → classify → plan → confirm → execute |
| `language.go`         | `I18nManager` — language matching & translation dispatch |
| `language_en.go`      | English translation map                                 |
| `language_zh-Hans.go` | Simplified Chinese translation map                      |
| `language_zh-Hant.go` | Traditional Chinese translation map                     |
| `language_ja.go`      | Japanese translation map                                |
| `formbutton.go`       | Custom tview `FormItem` for buttons in forms            |

### Rename Pipeline (`runRenameMode.go`)

1. Load config → apply CLI overrides → init I18N
2. Collect all file paths (handle directories, recursion, globs)
3. Classify: exclude‑matched → skip; sync‑matched → companion; rest → primary
4. For each primary: extract time → generate new name → store rename plan + name mapping
5. For each companion: match against primary mapping → generate corresponding rename plan
6. If `confirm`: print preview, ask y/N
7. Execute `os.Rename()` for each plan
8. If `endpause`: wait for Enter

## Dependencies (`go.mod`)

| Package                     | Purpose              |
|-----------------------------|----------------------|
| `github.com/gdamore/tcell/v2` | Terminal cell library |
| `github.com/rivo/tview`       | TUI widget framework |
| `github.com/rwcarlsen/goexif` | Built‑in EXIF parser |
| `golang.org/x/text`           | I18N matching & printing |

## Typical Usage Patterns

```bash
# Batch rename with custom format, no confirmation, recursive
EvernightMoments -f "<YYYY>-<MM>-<DD>_<HH>-<mm>-<ss>_<*>" -ny -r "C:\Photos"

# Override language and ExifTool path, exclude raw sidecars
EvernightMoments -l ja -x "D:\exiftool.exe" -e "*.dop" "C:\album"

# Quick rename using config defaults, just pass files
EvernightMoments "C:\Photos\IMG_001.jpg" "C:\Photos\IMG_002.jpg"

# Enter configuration TUI (no arguments)
EvernightMoments
```
