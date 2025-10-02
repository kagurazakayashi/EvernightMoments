![icon](ico/icon.ico)

# [EvernightMoments](https://github.com/kagurazakayashi/EvernightMoments)

**English** | [简体中文](README.zh-Hans.md) | [繁體中文](README.zh-Hant.md) | [日本語](README.ja.md)

**Bestowing eternity upon the fleeting, and warmth to the everlasting night.**
EvernightMoments is a utility that automatically renames your visual archives by extracting original capture timestamps.

## Download

Latest Version: v1.1.0 (go 1.26.0)

Go to [Releases](https://github.com/kagurazakayashi/EvernightMoments/releases) to download the latest version.

|   OS    |   Processor   | Arch | Package Name                      |
| :-----: | :-----------: | :--: | --------------------------------- |
| Windows | Intel/AMD x86 |  32  | EvernightMoments_windows-x86.7z   |
| Windows | Intel/AMD x86 |  64  | EvernightMoments_windows-x64.7z   |
| Windows |      ARM      |  64  | EvernightMoments_windows-arm64.7z |
|  macOS  |   Intel x86   |  64  | EvernightMoments_macos-x64.7z     |
|  macOS  | Apple Silicon |  64  | EvernightMoments_macos-arm64.7z   |
|  Linux  | Intel/AMD x86 |  32  | EvernightMoments_linux-x86.7z     |
|  Linux  | Intel/AMD x86 |  64  | EvernightMoments_linux-x64.7z     |
|  Linux  |      ARM      |  32  | EvernightMoments_linux-arm32.7z   |
|  Linux  |      ARM      |  64  | EvernightMoments_linux-arm64.7z   |

## Usage

### Quick Start

#### General Commands

`[executable_name] [photo_path1] [photo_path2] ...`

- **Executable name**:
  - In `Windows Command Prompt`: e.g., `EvernightMoments.exe`
  - In `Windows PowerShell`: e.g., `.\EvernightMoments.exe`
  - In macOS / Linux `sh`: e.g., `./EvernightMoments`
- **Photo file paths**:
  - Supports **multiple files**: e.g., `EvernightMoments.exe "C:\album1\photo1.arw" "C:\album1\photo2.arw"`
  - Supports **multiple folders**: e.g., `EvernightMoments.exe "C:\album1" "C:\album2"`
    - If a folder is specified, the tool will attempt to rename **all files** within that folder.
    - By default, subfolders are not modified. To include subfolders, add the `-r` parameter.

#### Using Graphical Interface (Windows)

- In **File Explorer**, you can directly drag and drop one or more **photo files** or **folders containing photos** onto the `.exe` file to start the renaming process.
- You can also add the program to the **"Send to"** menu for quick access via the right-click menu:
  - Press `Windows+R` to open the "Run" dialog, type `shell:sendto`, and press Enter. Copy the program into the folder that opens.
  - Select one or more photos/folders, right-click, find the "Send to" menu, and select this program.

### Software Configuration

To change the language, filename format, or interactive prompts, enter the configuration mode:

Run `./EvernightMoments` (or double-click the `.exe` in Windows) **without any parameters**.

Follow the prompts and press Enter after each answer. You can configure the following:

#### 1. Language Settings

- Select your preferred display language by entering its corresponding number.
- Once saved, you can press Enter to skip this step in the future.

#### 2. Filename Format

- You will see the current renaming format. Press Enter to keep it, or type a new format.
- Use the following **placeholders** (Case-sensitive):

| Placeholder | Example Output | Meaning             |
| :---------- | -------------: | ------------------- |
| `<YY>`      |           `25` | 2-digit Year        |
| `<YYYY>`    |         `2025` | 4-digit Year        |
| `<M>`       |            `5` | Month               |
| `<MM>`      |           `05` | 2-digit Month       |
| `<D>`       |            `2` | Day                 |
| `<DD>`      |           `02` | 2-digit Day         |
| `<H>`       |            `9` | Hour                |
| `<HH>`      |           `09` | 2-digit Hour        |
| `<m>`       |            `7` | Minute              |
| `<mm>`      |           `07` | 2-digit Minute      |
| `<s>`       |            `3` | Second              |
| `<ss>`      |           `03` | 2-digit Second      |
| `<#>`       |            `1` | Index Number        |
| `<##>`      |           `01` | Padded Index Number |
| `<*>`       |    `photo.jpg` | Original Filename   |

##### Important Notes

1. **Do not use illegal characters!** Prohibited symbols include: `\ / : ? ' " | < * >`
   - In Windows, filenames cannot be reserved names like `CON`, `PRN`, `AUX`, etc., and cannot end with a period `.`.
2. `<hh>` (12-hour format) is **not supported**.

##### Examples

With the original filename `photo.jpg`, the default format `<YYYY><MM><DD>_<HH><mm><ss>_<*>` will output `20260220_122937_photo.jpg`.

Additional examples:

- `<YYYY><MM><DD>_<HH><mm><ss>*` -> `20250502_090703photo.jpg`
- `<YY>-<M>-<D>-*` -> `25-5-2-photo.jpg`
- `<YYYY>-<MM>-<DD>_<HH>-<mm>-<ss>_<*>` -> `2025-05-02_09-07-03_photo.jpg`

#### 3. Enable Preview Confirmation?

- If enabled, the tool will show a preview of the changes and ask for confirmation before proceeding.
- Enter `y` (default) to ask every time, or `n` to rename immediately.

#### 4. "Press Enter to Exit" Prompt?

- Decide if the program should wait for a keypress after finishing so you can review the results.
- Enter `y` (default) to wait, or `n` to exit immediately.

## Build

First, install [Go](https://go.dev/). Version `1.26.0` or higher is required.

### Windows

You can use the following scripts:

- **Build**
  - `build.bat`: Compiles for various platforms (output in the `bin` folder).
- **Test**
  - First, create a `TestPhotos` folder and place some test photo files inside.
  - `conf.bat`: Builds and enters configuration mode.
  - `test_dir.bat`: Builds and tests processing the `TestPhotos` folder (includes all files).
  - `test_files.bat`: Tests multiple file inputs (takes all files from `TestPhotos`).

### All Systems

`cd` into the source code folder and execute:

```bash
go generate
go build .
```

## LICENSE

Copyright (c) 2026 KagurazakaYashi. EvernightMoments is licensed under **Mulan PSL v2**. You can use this software according to the terms and conditions of the Mulan PSL v2. You may obtain a copy of Mulan PSL v2 at: http://license.coscl.org.cn/MulanPSL2. THIS SOFTWARE IS PROVIDED ON AN “AS IS” BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FITNESS FOR A PARTICULAR PURPOSE. See the Mulan PSL v2 for more details.
