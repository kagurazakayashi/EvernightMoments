![icon](ico/icon.ico)

# [EvernightMoments](https://github.com/kagurazakayashi/EvernightMoments)

**English** | [简体中文](README.zh-Hans.md) | [繁體中文](README.zh-Hant.md) | [日本語](README.ja.md)

**Bestowing eternity upon the fleeting, and warmth to the everlasting night.**

EvernightMoments is a utility that automatically renames your visual archives by extracting original capture timestamps.

## Demo

Settings used for the demonstration:

1. Renaming Format: `<YYYY>-<MM>-<DD>_<HH>-<mm>-<ss>_<*>`
2. Preview Confirmation: Disabled
3. Pause on Exit: Disabled

### Demo 1 (Windows GUI)

Windows 10 Pro 22H2

![Demo1](ico/demo1.gif)

### Demo 2 (Bash Terminal)

UOS 1072 Pro

![Demo1](ico/demo2.gif)

## Download

Latest Version: v1.1.0

Go to [Releases](https://github.com/kagurazakayashi/EvernightMoments/releases) to download the latest version.

|   OS    |   Processor   | Arch | Package Name                           |
| :-----: | :-----------: | :--: | -------------------------------------- |
| Windows | Intel/AMD x86 |  32  | `EvernightMoments_v*_windows-x86.7z`   |
| Windows | Intel/AMD x86 |  64  | `EvernightMoments_v*_windows-x64.7z`   |
| Windows |      ARM      |  64  | `EvernightMoments_v*_windows-arm64.7z` |
|  macOS  |   Intel x86   |  64  | `EvernightMoments_v*_macos-x64.7z`     |
|  macOS  | Apple Silicon |  64  | `EvernightMoments_v*_macos-arm64.7z`   |
|  Linux  | Intel/AMD x86 |  32  | `EvernightMoments_v*_linux-x86.7z`     |
|  Linux  | Intel/AMD x86 |  64  | `EvernightMoments_v*_linux-x64.7z`     |
|  Linux  |      ARM      |  32  | `EvernightMoments_v*_linux-arm32.7z`   |
|  Linux  |      ARM      |  64  | `EvernightMoments_v*_linux-arm64.7z`   |

This program has been tested on the following platforms:

- Windows 11 25H2 (Intel & AMD x86_64)
- UOS 1072 (Intel & AMD x86_64)

## Installation and Uninstallation

### Installation on Windows

Simply extract the files to any location where you have write permissions without needing administrator privileges.

#### Add to the "Send to" Menu

Adding the program to the "Send to" menu allows you to use it anytime by right-clicking on photos.

1. Press `Windows+R` to open the "Run" dialog, type `shell:sendto`, and press Enter.
2. Copy the program (or a shortcut to it) into the folder that opens.
3. Select one or more **image files** or **folders containing only images**, right-click to show the context menu, and you will find the program under the "Send to" item.

#### Use as a Command Line Tool

1. Create a folder and copy the program into it.
2. Copy the full path of that folder.
3. Press `Windows+R`, type `rundll32.exe sysdm.cpl,EditEnvironmentVariables`, and press Enter.
4. Select `Path` (under User variables) and click the "Edit" button.
5. Click "New" and paste the folder path into the list.

#### Uninstalling on Windows

- Simply delete all related files and remove the path from your environment variables.

### Installing on Linux

1. Extract the files to a location where you have write permissions without needing root privileges.
2. Open the **Terminal** and use the following commands:

```bash
cd /path/to/extracted/directory
chmod +x install.sh
./install.sh
```

- This script will automatically:
  - Create configuration icons in the application menu and on the desktop.
  - Add the command to the `/.local/bin` folder for easy access in the terminal at any time.
    - If the command still cannot be found, you will need to manually add the absolute path of this folder to your `PATH` environment variable.
- **Note: Some operating systems will refuse to run unsigned programs.** You need to allow unsigned programs to run in order to use this software. Typically, in a graphical user interface (GUI), a pop-up dialog will appear if the system blocks the execution.

#### Uninstalling on Linux

```bash
cd /path/to/extracted/directory
chmod +x uninstall.sh
./uninstall.sh
```

Then, delete all related files and environment variables.

## Instructions

### Using the Windows GUI

- In **File Explorer**, you can drag and drop one or more **photo files** or **folders containing only photos** directly onto this .exe file to start the renaming process.
- You can also use it via the "Send to" menu; see [Installing on Windows](#installing-on-windows).

### General Commands

`[executable_name] [photo_path1] [photo_path2] ...`

- **Program Executable Filename**:
  - Environment variables configured:
    - Use `EvernightMoments` directly from any directory.
  - Running from the program directory:
    - Similar to `EvernightMoments.exe` in Windows Command Prompt
    - Similar to `.\EvernightMoments.exe` in `Windows PowerShell`
    - Similar to `./EvernightMoments` in macOS / Linux `sh`
- **Photo file paths**:
  - Supports **multiple files**: e.g., `EvernightMoments.exe "C:\album1\photo1.arw" "C:\album1\photo2.arw"`
  - Supports **multiple folders**: e.g., `EvernightMoments.exe "C:\album1" "C:\album2"`
    - If a folder is specified, the tool will attempt to rename **all files** within that folder.
    - By default, subfolders are not modified. To include subfolders, add the `-r` parameter.

## Software Configuration

To change the language, filename format, or interactive prompts, enter the configuration mode:

Run `./EvernightMoments` (or double-click the `.exe` in Windows) **without any parameters**.

Follow the prompts and press Enter after each answer. You can configure the following:

### 1. Language Settings

- Select your preferred display language by entering its corresponding number.
- Once saved, you can press Enter to skip this step in the future.

### 2. Filename Format

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

#### Important Notes

1. **Do not use illegal characters!** Prohibited symbols include: `\ / : ? ' " | < * >`
   - In Windows, filenames cannot be reserved names like `CON`, `PRN`, `AUX`, etc., and cannot end with a period `.`.
2. `<hh>` (12-hour format) is **not supported**.

#### Examples

With the original filename `photo.jpg`, the default format `<YYYY><MM><DD>_<HH><mm><ss>_<*>` will output `20260220_122937_photo.jpg`.

Additional examples:

- `<YYYY><MM><DD>_<HH><mm><ss>*` -> `20250502_090703Photo.jpg`
- `<*>_<HH><mm><ss>` -> `Photo_193030.jpg`
- `<YY>年<M>月<D>日—*` -> `25年5月2日-Photo.jpg`
- `<YYYY>-<MM>-<DD>_<HH>-<mm>-<ss>_<*>` -> `2025-05-02_09-07-03_Photo.jpg`
- `<MM>-<DD>-<YYYY>_<HH>-<mm>-<ss>_<*>` -> `05-02-2025_09-07-03_Photo.jpg`
- `<DD>.<MM>.<YYYY>_<HH>.<mm>.<ss>_<*>` -> `02.05.2025_09.07.03_Photo.jpg`

### 3. Enable Preview Confirmation?

- If enabled, the tool will show a preview of the changes and ask for confirmation before proceeding.
- Enter `y` (default) to ask every time, or `n` to rename immediately.

### 4. "Press Enter to Exit" Prompt?

- Decide if the program should wait for a keypress after finishing so you can review the results.
- Enter `y` (default) to wait, or `n` to exit immediately.

## Build

First, install [Go](https://go.dev/). The version must be `1.26.0` or higher.

### On Windows

You can use the `build.bat` script to compile for various platforms (output to the `bin` folder).

#### Testing

- First, create a `TestPhotos` folder and place some test photo files inside.
- `conf.bat`: Build and enter configuration mode.
- `test_dir.bat`: Build and test processing the `TestPhotos` folder (including all files).
- `test_files.bat`: Test multiple file inputs (using all files from `TestPhotos`).
- If the photos in the `TestPhotos` folder were renamed using the default format during testing, you can run `python test_undo.py` to undo the renaming.

### On All Systems

`cd` into the source code folder and run:

```bash
go generate
go build .
```

## LICENSE

Copyright (c) 2026 KagurazakaYashi. EvernightMoments is licensed under **Mulan PSL v2**. You can use this software according to the terms and conditions of the Mulan PSL v2. You may obtain a copy of Mulan PSL v2 at: http://license.coscl.org.cn/MulanPSL2. THIS SOFTWARE IS PROVIDED ON AN “AS IS” BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FITNESS FOR A PARTICULAR PURPOSE. See the Mulan PSL v2 for more details.
