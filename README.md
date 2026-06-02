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

Latest Version: v1.2.0

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

Packages with `ExifTool` indicate that the ExifTool executable is already built-in. ExifTool is optional, but it provides extensive media file support and updated camera support. For details, please see [Install ExifTool to Improve Recognition Capabilities](#install-exiftool-to-improve-recognition-capabilities).

This program has been tested on the following platforms:

- Windows 11 25H2 (Intel & AMD x86_64)
- UOS 1072 (Intel & AMD x86_64)
- Arch Linux 6.16 (Intel & AMD x86_64)
- macOS 15.7 (Intel)

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

### Installation on Linux

1. Extract to a location where write access does not require root privileges.
2. Open the **Terminal** and use the following commands:

```bash
cd path_to_extracted_directory
chmod +x install.sh
./install.sh
```

- This script will automatically:
  - Install the program to `~/.local/bin/EvernightMoments`.
  - Install the icon resources to `~/.local/share/icons/EvernightMoments.png`.
  - Create an application menu entry for the configuration program at `~/.local/share/applications/EvernightMoments.desktop`, categorized under "Graphics".
  - Create a desktop icon for the configuration program at `~/Desktop/EvernightMoments.desktop`.
  - Add the command to the `~/.local/bin` folder so it can be easily accessed in the terminal at any time.
    - If `~/.local/bin` is not in your `PATH` environment variable, it will attempt to add it automatically.
- **Note: Some operating systems will block unsigned programs from running.** You need to allow unsigned programs to execute in order to use this program. Generally, in a graphical environment, you will see a popup alert if it is blocked by the system.
- The program's configuration will be stored in `~/.config/EvernightMoments/config.json`.

#### Uninstallation on Linux

```bash
cd path_to_extracted_directory
chmod +x uninstall.sh
./uninstall.sh
```

Then manually delete all related files and environment variables.

### Installation on macOS

1. Extract to a location where write access does not require root privileges.
2. Open the **Terminal** and use the following commands:

```bash
cd path_to_extracted_directory
chmod +x install.sh
./install-mac.sh
```

- This script will automatically:
  - Install the program to `~/.local/bin/EvernightMoments`.
  - Generate the configuration program at `~/Applications/EvernightMoments Config.app`.
  - Create a desktop icon for the configuration program at `~/Desktop/EvernightMoments Config.app`.
    - You can move it from the desktop to the `/Applications` :
    - `mv "$HOME/Desktop/EvernightMoments Config.app" "/Applications/EvernightMoments Config.app"`
  - Add the command to the `~/.local/bin` folder so it can be easily accessed in the terminal at any time.
    - If `~/.local/bin` is not in your `PATH` environment variable, it will attempt to add it automatically.
- **Note: The system might block unsigned programs from running.** You need to allow unsigned programs to execute in order to use this program.
  - If blocked, please go to "Privacy & Security" in "System Settings", scroll down to the bottom, find the "Allow" (or "Open Anyway") button, and click it.
- The program's configuration will be stored in `~/Library/Application Support/EvernightMoments/config.json`.

#### Uninstallation on macOS

```bash
cd path_to_extracted_directory
chmod +x uninstall.sh
./uninstall-mac.sh
```

Then manually delete all related files and environment variables.

### Install ExifTool to Improve Recognition Capabilities

The built-in EXIF parser for this program is [goexif](https://github.com/rwcarlsen/goexif), which can only handle a limited number of formats. Furthermore, due to the update limitations of both this library and this program, timely support for new camera devices cannot be guaranteed. It is strongly recommended that you install ExifTool to obtain the latest camera support.

[ExifTool](https://github.com/exiftool/exiftool) is a free and open-source software developed by Phil Harvey, specifically designed for reading, writing, and editing metadata in images, videos, and audio. By keeping this tool updated, you can equip this program with the latest RAW camera support and the ability to process more file formats.

If the program finds a valid ExifTool, the file path of the "EXIF Extractor" will be displayed in the startup output. Additionally, during the renaming process, if ExifTool successfully retrieves the timestamp, "ExifTool" will be shown as the "Time Source."

Regarding the download and installation instructions for ExifTool:

#### Install ExifTool using a Package Manager

If you have a package manager installed on your system, you can use your preferred one for a quick installation. For example:

- Windows: `choco install exiftool` or `scoop install exiftool`
- macOS: `brew install exiftool`
- Debian / Ubuntu / Mint: `sudo apt update && sudo apt install perl libimage-exiftool-perl`
- CentOS / RHEL / Fedora: `sudo dnf install perl perl-Image-ExifTool`
- Arch Linux: `sudo pacman -S perl perl-image-exiftool`

If you do not have a package manager, you can install it following these steps:

#### Download and Install ExifTool

First, go to the [ExifTool Homepage](https://github.com/exiftool/exiftool) to download the latest version of the program for your specific operating system.

- **Windows** has two methods:
  - **Global Installation**: Please refer to the official [Installation and Uninstallation Instructions](https://exiftool.org/install.html#Windows). Once completed, ensure that you can enter the `exiftool.exe` command from any location in the "Command Prompt".
  - **For this program only**: You will receive a zip file. Extract all ExifTool files from this zip archive into this program's folder (make sure `exiftool(-k).exe` or `exiftool.exe` is located in the same folder as `EvernightMoments.exe`).
- **macOS**: Please refer to the official [Installation and Uninstallation Instructions](https://exiftool.org/install.html#MacOS).
- **Linux**: Please refer to the official [Installation and Uninstallation Instructions](https://exiftool.org/install.html#Unix).

## Instructions

### Using the Windows GUI

- In **File Explorer**, you can drag and drop one or more **photo files** or **folders containing only photos** directly onto this .exe file to start the renaming process.
- You can also use it via the "Send to" menu; see [Installation on Windows](#installation-on-windows).

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

### Command-Line Parameters

You can temporarily override any configuration item via command-line flags. Each flag has a short form (`-x`) and a long form (`--xxx`). File/folder paths do not need a flag name — place them at the end. Multiple files, folders, and wildcards are supported.

| Short | Long              | Description                                          | Example                                         |
| :---- | :---------------- | :--------------------------------------------------- | :---------------------------------------------- |
| `-l`  | `--language`      | Set display language (`en`/`zh-Hans`/`zh-Hant`/`ja`) | `-l en`                                         |
| `-f`  | `--format`        | Set renaming format template                         | `-f "<YYYY>-<MM>-<DD>_<*>""`                    |
| `-e`  | `--exclude`       | Add an exclude glob pattern (repeatable)             | `-e "*.dop" -e "*.cos"`                         |
| `-s`  | `--sync`          | Add a sync glob pattern (repeatable)                 | `-s "*.dop"`                                    |
| `-y`  | `--confirm`       | Enable preview confirmation                          | `-y`                                            |
| `-ny` | `--no-confirm`    | Disable preview confirmation                         | `-ny`                                           |
| `-p`  | `--pause`         | Enable pause before exit                             | `-p`                                            |
| `-np` | `--no-pause`      | Disable pause before exit                            | `-np`                                           |
| `-x`  | `--exiftool`      | Set ExifTool executable path (empty to disable)      | `-x "C:\Tools\exiftool.exe"`                    |
| `-r`  | `--recursive`     | Process subdirectories recursively                   | `-r`                                            |
| `-me` | `--multi-ext`     | Treat multi-level extensions as part of filename     | `-me`                                           |
| `-nc` | `--no-color`      | Disable colored terminal output                      | `-nc`                                           |

**Examples:**

```bash
# Override format and language, disable confirmation, with recursive
EvernightMoments -f "<YYYY>-<MM>-<DD>_<*>" -l en -ny -r "C:\Photos"

# Override ExifTool path, add exclude patterns
EvernightMoments -x "D:\exiftool.exe" -e "*.dop" -e "*.cos" "C:\album1" "C:\album2"

# Rename .ARW files in multiple folders, with .ARW.dop sidecars synced
EvernightMoments -f "<YYYY><MM><DD>_<HH><mm><ss><*>" -s "*.ARW.dop" -ny "D:\DCIM\10860213" "D:\DCIM\11051228"
```

### AI Agent Integration

The project includes a [`SKILL.md`](SKILL.md) file that describes the tool's architecture, CLI interface, configuration schema, and typical workflows in a format optimized for AI agents. To let your AI assistant understand and operate EvernightMoments, load this file as a skill or provide it as context.

## Software Configuration

To change the language, filename format, exclude/sync patterns, the ExifTool path, or interactive prompts, enter the configuration mode:

Run `./EvernightMoments` (or double-click the `.exe` in Windows) **without any parameters** to open the full-screen TUI settings interface.

How to navigate the interface:

- Use `Tab` or the `Arrow keys` to move focus between fields.
- Type directly into input fields; toggle switches with `Enter` or `Space`.
- Choose **"Save & Exit"** to write the configuration, or press `Esc` / choose **"Quit"** to discard all changes.

You can configure the following:

### 1. Language Settings

- Pick a display language from the "Language" dropdown; the interface text switches to the selected language **instantly**.
- Your choice is saved along with the other settings and pre-selected next time.

### 2. Filename Format

- Edit the renaming format in the "Naming Format" input field; the footer shows a **live preview** of the generated name, and illegal characters are blocked as you type.
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

### 3. Exclude Patterns (skip these files)

- Enter path patterns to exclude in the "Exclude" input field, separated by commas.
- Supports both absolute and relative paths, not just extensions. Examples:
  - `CaptureOne\Settings153\*.cos`
  - `*.dop`
- Matching files will not be renamed. Leave it empty to exclude nothing.

### 4. Sync Patterns (renamed alongside the primary file)

- Enter path patterns in the "Sync" input field, separated by commas. Supports both absolute and relative paths, not just extensions. Examples:
  - `CaptureOne\Settings153\*.cos`
  - `*.dop`
- When a primary photo file is renamed, any matching companion file in the same folder is renamed using the **same** new filename.
- **Multi‑extension companions:** For files with multiple extensions (e.g. `photo.ARW.dop`), the tool automatically strips the intermediate extensions to locate the primary file (e.g. matches `photo.ARW` or `photo`).
- **Note:** Files matched by sync patterns are also treated as excluded by default (they will not be renamed based on their own content).

### 5. ExifTool Path

- Specify the path to the `exiftool` executable in the "ExifTool path" field to read more accurate capture times.
- The default value is the path auto-detected by the program; it is empty if nothing was found.
- **Leave it empty** to skip ExifTool and use only the built-in parser.
- Click the **"Auto-detect"** button below to detect the path from the system `PATH` again.

### 6. Enable Preview Confirmation?

- The "Ask preview" toggle: when checked, the tool shows a preview of the changes and asks for confirmation before proceeding.
- Uncheck it to start renaming immediately (proceed with caution).

### 7. Wait for "Press Enter to Exit" on finish?

- The "Wait on exit" toggle: when checked, the program pauses after finishing so you can review the results; uncheck it to exit immediately.

### 8. Treat Multi-Level Extensions as Filename?

- The "Long ext file" toggle: when checked, only the last extension is stripped — intermediate extensions (e.g. `.ARW` in `photo.ARW.dop`) are treated as part of the filename and preserved. When unchecked (default), all extensions are stripped, keeping only the bare base name.

### 9. Enable Colored Output?

- The "Color output" toggle: when checked, ANSI color codes are enabled in terminal output. Uncheck it to disable colors — recommended when redirecting output or using the tool via an AI agent.

## Build

First, install [Go](https://go.dev/). The version must be `1.26.0` or higher.

### Compiling on Windows

1. Use `mdhtml.bat` to create help documentation (output to the `readme` folder).
2. Use `build.bat` to compile for various platforms (output to the `bin` folder).

#### Testing

- First, create a `TestPhotos` folder and place some test photo files inside.
- `conf.bat`: Build and enter configuration mode.
- `test_dir.bat`: Build and test processing the `TestPhotos` folder (including all files).
- `test_files.bat`: Test multiple file inputs (using all files from `TestPhotos`).
- If the photos in the `TestPhotos` folder were renamed using the default format during testing, you can run `python test_undo.py` to undo the renaming.

### Compiling on macOS / Linux

Same as above, just replace `.bat` with `.sh`.

## LICENSE

Copyright (c) 2026 KagurazakaYashi. EvernightMoments is licensed under **Mulan PSL v2**. You can use this software according to the terms and conditions of the Mulan PSL v2. You may obtain a copy of Mulan PSL v2 at: http://license.coscl.org.cn/MulanPSL2. THIS SOFTWARE IS PROVIDED ON AN “AS IS” BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FITNESS FOR A PARTICULAR PURPOSE. See the Mulan PSL v2 for more details.
