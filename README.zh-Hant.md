![icon](ico/icon.ico)

# [EvernightMoments](https://github.com/kagurazakayashi/EvernightMoments)

[English](README.md) | [简体中文](README.zh-Hans.md) | **繁體中文** | [日本語](README.ja.md)

**予瞬息以永恆，於長夜留餘溫。**

EvernightMoments 是一款透過提取照片原始拍攝時間，為您自動重新命名的工具。

## 演示

演示時採用的設定：

1. 重新命名格式：`<YYYY>-<MM>-<DD>_<HH>-<mm>-<ss>_<*>`
2. 預覽確認：停用
3. 結束時暫停：停用

### 演示 1 (Windows 圖形介面)

Windows 10 Pro 22H2

![Demo1](ico/demo1.gif)

### 演示 2 (Bash 命令列終端)

UOS 1072 Pro

![Demo1](ico/demo2.gif)

## 下載

最新版本: v1.2.0

前往 [Releases](https://github.com/kagurazakayashi/EvernightMoments/releases) 下載最新版本。

| 作業系統 |    處理器     | 位數 | 軟體壓縮包名稱                         |
| :------: | :-----------: | :--: | -------------------------------------- |
| windows  | Intel/AMD x86 |  32  | `EvernightMoments_v*_windows-x86.7z`   |
| windows  | Intel/AMD x86 |  64  | `EvernightMoments_v*_windows-x64.7z`   |
| windows  |      ARM      |  64  | `EvernightMoments_v*_windows-arm64.7z` |
|  macOS   |   Intel x86   |  64  | `EvernightMoments_v*_macos-x64.7z`     |
|  macOS   | Apple silicon |  64  | `EvernightMoments_v*_macos-arm64.7z`   |
|  Linux   | Intel/AMD x86 |  32  | `EvernightMoments_v*_linux-x86.7z`     |
|  Linux   | Intel/AMD x86 |  64  | `EvernightMoments_v*_linux-x64.7z`     |
|  Linux   |      ARM      |  32  | `EvernightMoments_v*_linux-arm32.7z`   |
|  Linux   |      ARM      |  64  | `EvernightMoments_v*_linux-arm64.7z`   |

標有 `ExifTool` 的軟體包表示已內建 ExifTool 執行檔。ExifTool 是選配的，但它提供了更廣泛的媒體檔案支援以及最新的相機型號支援。詳情請參閱[安裝 ExifTool 以提升辨識能力](#安裝-exiftool-以提升識別能力)。

本程式已在以下平台完成測試：

- Windows 11 25H2 (Intel & AMD x86_64)
- UOS 1072 (Intel & AMD x86_64)
- Arch Linux 6.16 (Intel & AMD x86_64)
- macOS 15.7 (Intel)

## 安裝與解除安裝

### 在 Windows 上安裝

直接解壓縮到一個不需要管理者權限就能寫入的任意位置即可。

#### 安裝到「傳送到」選單

將程式放到「傳送到」選單中，可以隨時透過在相片上按右鍵選單中的「傳送到」來使用本程式。

1. 按下 `Windows+R` 快速鍵開啟「執行」對話框，輸入 `shell:sendto` 並按下回車（Enter）。
2. 將程式複製到新開啟的資料夾視窗中。
3. 選取一個或多個 **相片檔案** 或 **只含有相片的資料夾** 並按下滑鼠右鍵，顯示完整右鍵選單，找到「傳送到」項目，即可看到本程式。

#### 在命令列中隨時使用此命令

1. 建立一個資料夾，並將程式複製進去。
2. 複製該資料夾的路徑。
3. 按下 `Windows+R` 快速鍵開啟「執行」對話框，輸入 `rundll32.exe sysdm.cpl,EditEnvironmentVariables` 並按下回車。
4. 選取 `Path` 並按下「編輯」按鈕。
5. 按下「新增」按鈕，將該資料夾的路徑貼進去。

#### 在 Windows 上解除安裝

- 刪除所有相關檔案及環境變數即可。

### 在 Linux 上安裝

1. 解壓縮到不需 root 權限就能寫入的地方。
2. 進入**終端機**，使用以下指令：

```bash
cd 剛解壓縮的檔案目錄
chmod +x install.sh
./install.sh
```

- 該腳本會自動完成：
  - 安裝程式到 `~/.local/bin/EvernightMoments`。
  - 安裝圖示資源到 `~/.local/share/icons/EvernightMoments.png`。
  - 為設定程式建立應用程式列表項目 `~/.local/share/applications/EvernightMoments.desktop`，分類為「圖形」。
  - 為設定程式建立桌面圖示 `~/Desktop/EvernightMoments.desktop`。
  - 將指令新增至 `~/.local/bin` 資料夾中，以便在終端機裡隨時使用。
    - 如果 `~/.local/bin` 不在 `PATH` 環境變數中，則會嘗試自動加入。
- **注意：部分作業系統會拒絕未經簽署的程式執行**，您需要允許未經簽署的程式執行才能使用本程式。通常在圖形介面下，如果被系統拒絕執行，會看到彈出視窗提示。
- 程式的設定將儲存至 `~/.local/bin/EvernightMoments.json`。

#### 在 Linux 上解除安裝

```bash
cd 剛解壓縮的檔案目錄
chmod +x uninstall.sh
./uninstall.sh
```

然後刪除所有相關檔案與環境變數。

### 在 macOS 上安裝

1. 解壓縮到不需 root 權限就能寫入的地方。
2. 進入**終端機**，使用以下指令：

```bash
cd 剛解壓縮的檔案目錄
chmod +x install.sh
./install-mac.sh
```

- 該腳本會自動完成：
  - 安裝程式到 `~/.local/bin/EvernightMoments`。
  - 產生設定程式 `~/Applications/EvernightMoments Config.app`。
  - 為設定程式建立桌面圖示 `~/Desktop/EvernightMoments Config.app`。
    - 你可以將它從桌面上移動到 `/Applications` (應用程式資料夾) 裡面：
    - `mv "$HOME/Desktop/EvernightMoments Config.app" "/Applications/EvernightMoments Config.app"`
  - 將指令新增至 `~/.local/bin` 資料夾中，以便在終端機裡隨時使用。
    - 如果 `~/.local/bin` 不在 `PATH` 環境變數中，則會嘗試自動加入。
- **注意：系統可能會拒絕未經簽署的程式執行**，您需要允許未經簽署的程式執行才能使用本程式。
  - 如果遇到阻擋，請前往「系統設定」中的「隱私權與安全性」，捲動到最底部，找到「強制打開」或「允許」按鈕並點擊。
- 程式的設定將儲存至 `~/.local/bin/EvernightMoments.json`。

#### 在 macOS 上解除安裝

```bash
cd 剛解壓縮的檔案目錄
chmod +x uninstall.sh
./uninstall-mac.sh
```

然後刪除所有相關檔案與環境變數。

### 安裝 ExifTool 以提升識別能力

本程式內建的 EXIF 解析器為 [goexif](https://github.com/rwcarlsen/goexif)，它只能處理有限的格式。並且由於該函式庫和本程式的更新限制，對新的相機設備無法提供及時的支援。強烈建議您安裝 ExifTool 來獲取最新的相機支援能力。

[ExifTool](https://github.com/exiftool/exiftool) 是 Phil Harvey 開發的自由開源軟體，專門用於處理影像、影片及音訊的 metadata，只要保持更新該程式，就能為本程式提供最新的相機 RAW 支援和處理更多類型的檔案格式。

如果本程式找到了可用的 ExifTool，會在起始輸出的地方顯示「EXIF 擷取器」的檔案路徑。並且在重新命名的時候，如果 ExifTool 成功取得時間，「時間來源」處會顯示「ExifTool」。

有關 ExifTool 的下載和安裝說明：

#### 使用套件管理員安裝 ExifTool

如果你的系統中安裝有套件管理員，可以使用你熟悉的套件管理員完成快速安裝。例如：

- Windows: `choco install exiftool` 或者 `scoop install exiftool`
- macOS: `brew install exiftool`
- Debian / Ubuntu / Mint: `sudo apt update && sudo apt install perl libimage-exiftool-perl`
- CentOS / RHEL / Fedora: `sudo dnf install perl perl-Image-ExifTool`
- Arch Linux: `sudo pacman -S perl perl-image-exiftool`

如果沒有套件管理員，可以按下面步驟安裝：

#### 下載安裝 ExifTool

先前往 [ExifTool 主頁](https://github.com/exiftool/exiftool) 下載對應系統的最新版本程式。

- **Windows** 有兩種方式:
  - **全域安裝**: 請查看官方[安裝與解除安裝說明](https://exiftool.org/install.html#Windows)操作。在完成後，應確保在「命令提示字元」中任意位置可以輸入 `exiftool.exe` 命令。
  - **只限本程式使用**: 你將得到一個 zip 檔案，將該 zip 檔案中所有的 ExifTool 檔案解壓縮到本程式資料夾內（確保 `exiftool(-k).exe` 或 `exiftool.exe` 和 `EvernightMoments.exe` 在同一個資料夾內）。
- **macOS**: 請查看官方[安裝與解除安裝說明](https://exiftool.org/install.html#MacOS)操作。
- **Linux**: 請查看官方[安裝與解除安裝說明](https://exiftool.org/install.html#Unix)操作。

## 使用說明

### 在 Windows 的圖形介面中使用

- 在「檔案總管」中，可以將一個或多個 **照片檔案** 或者 **內部只有照片的資料夾** 直接拖曳到這個 .exe 檔案上，開始重新命名操作。
- 可以透過「傳送到」選單使用，詳見[在 Windows 上安裝](#在-windows-上安裝)。

### 通用指令

`[程式可執行檔名] [照片檔案路徑1] [照片檔案路徑2] ...`

- **程式執行檔名稱**:
  - 已設定環境變數：
    - 直接在任何目錄下使用 `EvernightMoments`。
  - 從程式所在目錄執行：
    - Windows 命令提示字元中類似於 `EvernightMoments.exe`
    - `Windows PowerShell` 中類似於 `.\EvernightMoments.exe`
    - macOS / Linux `sh` 中類似於 `./EvernightMoments`
- **照片檔案路徑**:
  - 支援 **多個檔案** ，例如 `EvernightMoments.exe "C:\album1\photo1.arw" "C:\album1\photo2.arw"`
  - 支援 **多個資料夾** ，例如 `EvernightMoments.exe "C:\album1" "C:\album2"`
    - 如果指定的是一個資料夾，將嘗試重新命名該資料夾中的 **所有檔案** 。
    - 預設不會修改子資料夾中的檔案，如果需要同時修改子資料夾中的檔案，請新增引數 `-r` 。

## 軟體配置

如果需要更改 語言、檔案重新命名的格式、互動詢問開關 等，請先進行軟體配置：

**不帶引數** 直接執行 `./EvernightMoments` (或者在 Windows 中直接雙擊 .exe )，將會進入配置模式。

根據提示資訊，回答每個問題並回車。
你可以進行以下設定：

### 1. 語言設定

- 首先將會讓你選擇顯示的語言，輸入序號並回車即可。
- 設定被最終儲存後，下次再進入設定時此處直接按回車跳過即可。

### 2. 配置檔案重新命名的格式

- 你可以看到當前的檔案重新命名的格式，如果沒有問題，可以直接留空並回車，將跳過修改。
- 如果需要修改，請輸入新的格式。在格式中，可以使用以下**替代符**代指重新命名時的資訊。
- 注意：區分大小寫。

| 替代符   |    輸出示例 | 含義         |
| :------- | ----------: | ------------ |
| `<YY>`   |        `25` | 兩位年份     |
| `<YYYY>` |      `2025` | 完整年份     |
| `<M>`    |         `5` | 月份         |
| `<MM>`   |        `05` | 兩位月份     |
| `<D>`    |         `2` | 日           |
| `<DD>`   |        `02` | 兩位日       |
| `<H>`    |         `9` | 小時         |
| `<HH>`   |        `09` | 兩位小時     |
| `<m>`    |         `7` | 分鐘         |
| `<mm>`   |        `07` | 兩位分鐘     |
| `<s>`    |         `3` | 秒           |
| `<ss>`   |        `03` | 兩位秒       |
| `<#>`    |         `1` | 編號         |
| `<##>`   |        `01` | 統一位數編號 |
| `<*>`    | `photo.jpg` | 原始檔名     |

#### 注意事項

1. **不要輸入系統不支援的符號！** 包括以下符號: `\ / : ? ' " | < * >`
   - 其中 Windows 中最終命名結果不能為 `CON, PRN, AUX, NUL, COM1~COM9, LPT1~LPT9`, 檔名結尾不能是 `.`
   - 使用系統不支援的符號可能導致重新命名失敗甚至損壞其他檔案或檔案系統。
2. 不支援 `<hh>` (12 小時格式)

#### 示例

原始檔名為 `photo.jpg` 時，採用預設值格式 `<YYYY><MM><DD>_<HH><mm><ss>_<*>` ，將會輸出 `20260220_122937_Photo.jpg` 。

更多示例：

- `<YYYY><MM><DD>_<HH><mm><ss>*` -> `20250502_090703Photo.jpg`
- `<*>_<HH><mm><ss>` -> `Photo_193030.jpg`
- `<YY>年<M>月<D>日—*` -> `25年5月2日-Photo.jpg`
- `<YYYY>-<MM>-<DD>_<HH>-<mm>-<ss>_<*>` -> `2025-05-02_09-07-03_Photo.jpg`
- `<MM>-<DD>-<YYYY>_<HH>-<mm>-<ss>_<*>` -> `05-02-2025_09-07-03_Photo.jpg`
- `<DD>.<MM>.<YYYY>_<HH>.<mm>.<ss>_<*>` -> `02.05.2025_09.07.03_Photo.jpg`

### 3. 是否在重新命名前預覽確認？

- 如果開啟，會先向你展示會修改成什麼樣子，讓你確認要不要繼續。
- 請輸入 `y` 或者 `n`。
  - 預設值`y`: 每次先詢問我。
  - `n`: 直接開始重新命名。

### 4. 在執行結束後，需要提示“按回車鍵退出”以便停留檢視結果嗎？

- 請輸入 `y` 或者 `n`。
  - 預設值`y`: 結束後等待使用者按回車鍵。
  - `n`: 結束後直接退出。

## 編譯

先安裝 [Go](https://go.dev/)，版本需大於等於 `1.26.0`。

### 在 Windows 系統中編譯

1. 使用 `mdhtml.bat` 建立說明文件（建立至 `readme` 資料夾）。
2. 使用 `build.bat` 編譯至各個平台（編譯至 `bin` 資料夾）。

#### 測試

- 先建立 `TestPhotos` 資料夾，在裡面放一些測試用的照片檔案。
- `conf.bat`: 編譯並進入設定模式。
- `test_dir.bat`: 編譯並測試處理 `TestPhotos` 資料夾（包含所有檔案）。
- `test_files.bat`: 測試多個檔案輸入（取自 `TestPhotos` 中的所有檔案）。
- 如果測試時使用了預設格式對 `TestPhotos` 資料夾中的照片進行了重新命名，可以執行 `python test_undo.py` 復原重新命名。

### 在 macOS / Linux 系統中編譯

同上，將 `.bat` 換成 `.sh` 即可。

## LICENSE

Copyright (c) 2026 KagurazakaYashi EvernightMoments is licensed under Mulan PSL v2. You can use this software according to the terms and conditions of the Mulan PSL v2. You may obtain a copy of Mulan PSL v2 at: http://license.coscl.org.cn/MulanPSL2 THIS SOFTWARE IS PROVIDED ON AN “AS IS” BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE. See the Mulan PSL v2 for more details.
