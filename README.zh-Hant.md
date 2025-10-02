![icon](ico/icon.ico)

# [EvernightMoments](https://github.com/kagurazakayashi/EvernightMoments)

[English](README.md) | [简体中文](README.zh-Hans.md) | **繁體中文** | [日本語](README.ja.md)

**予瞬息以永恆，於長夜留餘溫。**
EvernightMoments 是一款透過提取照片原始拍攝時間，為您自動重新命名的工具。

## 下載

最新版本: v1.1.0 (go 1.26.0)

前往 [Releases](https://github.com/kagurazakayashi/EvernightMoments/releases) 下載最新版本。

| 作業系統 |    處理器     | 位數 | 軟體壓縮包名稱                    |
| :------: | :-----------: | :--: | --------------------------------- |
| windows  | Intel/AMD x86 |  32  | EvernightMoments_windows-x86.7z   |
| windows  | Intel/AMD x86 |  64  | EvernightMoments_windows-x64.7z   |
| windows  |      ARM      |  64  | EvernightMoments_windows-arm64.7z |
|  macOS   |   Intel x86   |  64  | EvernightMoments_macos-x64.7z     |
|  macOS   | Apple silicon |  64  | EvernightMoments_macos-arm64.7z   |
|  Linux   | Intel/AMD x86 |  32  | EvernightMoments_linux-x86.7z     |
|  Linux   | Intel/AMD x86 |  64  | EvernightMoments_linux-x64.7z     |
|  Linux   |      ARM      |  32  | EvernightMoments_linux-arm32.7z   |
|  Linux   |      ARM      |  64  | EvernightMoments_linux-arm64.7z   |

## 使用方法

### 快速開始

#### 通用命令

`[程式可執行檔名] [照片檔案路徑1] [照片檔案路徑2] ...`

- **程式可執行檔名**:
  - `Windows Command Prompt` 中類似於 `EvernightMoments.exe`
  - `Windows PowerShell` 中類似於 `.\EvernightMoments.exe`
  - macOS / Linux `sh` 中類似於 `./EvernightMoments`
- **照片檔案路徑**:
  - 支援 **多個檔案** ，例如 `EvernightMoments.exe "C:\album1\photo1.arw" "C:\album1\photo2.arw"`
  - 支援 **多個資料夾** ，例如 `EvernightMoments.exe "C:\album1" "C:\album2"`
    - 如果指定的是一個資料夾，將嘗試重新命名該資料夾中的 **所有檔案** 。
    - 預設不會修改子資料夾中的檔案，如果需要同時修改子資料夾中的檔案，請新增引數 `-r` 。

#### 在 Windows 的圖形畫面中使用

- 在“檔案資源管理器”中，可以將一個或多個 **照片檔案** 或者 **內部只有照片的資料夾** 直接拖拽到這個 .exe 檔案上，開始重新命名操作。
- 你還可以將程式放到“傳送到”選單裡面，可以隨時透過在照片上點右鍵中的“傳送到”選單來使用本程式。
  - 按 `Windows+R` 快捷鍵開啟“執行”對話方塊，然後輸入 `shell:sendto` 並回車，將程式複製到新開啟的資料夾視窗中。
  - 然後選中一個或多個 **照片檔案** 或者 **內部只有照片的資料夾** 並點按滑鼠右鍵，顯示完整右鍵選單，找到“傳送到”選單項，可以看到本程式。

### 軟體配置

如果需要更改 語言、檔案重新命名的格式、互動詢問開關 等，請先進行軟體配置：

**不帶引數** 直接執行 `./EvernightMoments` (或者在 Windows 中直接雙擊 .exe )，將會進入配置模式。

根據提示資訊，回答每個問題並回車。
你可以進行以下設定：

#### 1. 語言設定

- 首先將會讓你選擇顯示的語言，輸入序號並回車即可。
- 設定被最終儲存後，下次再進入設定時此處直接按回車跳過即可。

#### 2. 配置檔案重新命名的格式

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

##### 注意事項

1. **不要輸入系統不支援的符號！** 包括以下符號: `\ / : ? ' " | < * >`
   - 其中 Windows 中最終命名結果不能為 `CON, PRN, AUX, NUL, COM1~COM9, LPT1~LPT9`, 檔名結尾不能是 `.`
   - 使用系統不支援的符號可能導致重新命名失敗甚至損壞其他檔案或檔案系統。
2. 不支援 `<hh>` (12 小時格式)

##### 示例

原始檔名為 `photo.jpg` 時，採用預設值格式 `<YYYY><MM><DD>_<HH><mm><ss>_<*>` ，將會輸出 `20260220_122937_Photo.jpg` 。

更多示例：

- `<YYYY><MM><DD>_<HH><mm><ss>*` -> `20250502_090703Photo.jpg`
- `<YY>年<M>月<D>日—*` -> `25年5月2日-Photo.jpg`
- `<YYYY>-<MM>-<DD>_<HH>-<mm>-<ss>_<*>` -> `2025-05-02_09-07-03_Photo.jpg`
- `<MM>-<DD>-<YYYY>_<HH>-<mm>-<ss>_<*>` -> `05-02-2025_09-07-03_Photo.jpg`
- `<DD>.<MM>.<YYYY>_<HH>.<mm>.<ss>_<*>` -> `02.05.2025_09.07.03_Photo.jpg`

#### 3. 是否在重新命名前預覽確認？

- 如果開啟，會先向你展示會修改成什麼樣子，讓你確認要不要繼續。
- 請輸入 `y` 或者 `n`。
  - 預設值`y`: 每次先詢問我。
  - `n`: 直接開始重新命名。

#### 4. 在執行結束後，需要提示“按回車鍵退出”以便停留檢視結果嗎？

- 請輸入 `y` 或者 `n`。
  - 預設值`y`: 結束後等待使用者按回車鍵。
  - `n`: 結束後直接退出。

## 編譯

先安裝 [Go](https://go.dev/)，版本需高於或等於 `1.26.0`。

### Windows 系統下

可以使用以下腳本：

- **編譯**
  - `build.bat`：編譯到各個平台（位於 `bin` 資料夾中）。
- **測試**
  - 先建立 `TestPhotos` 資料夾，並在其中放入一些測試用的照片檔案。
  - `conf.bat`：編譯並進入設定模式。
  - `test_dir.bat`：編譯並測試處理 `TestPhotos` 資料夾（包括所有檔案）。
  - `test_files.bat`：測試多個檔案輸入（取自 `TestPhotos` 中的所有檔案）。

### 所有系統下

`cd` 進入原始碼資料夾並執行：

```bash
go generate
go build .
```

## LICENSE

Copyright (c) 2026 KagurazakaYashi EvernightMoments is licensed under Mulan PSL v2. You can use this software according to the terms and conditions of the Mulan PSL v2. You may obtain a copy of Mulan PSL v2 at: http://license.coscl.org.cn/MulanPSL2 THIS SOFTWARE IS PROVIDED ON AN “AS IS” BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE. See the Mulan PSL v2 for more details.
