![icon](ico/icon.ico)

# [EvernightMoments](https://github.com/kagurazakayashi/EvernightMoments)

**予瞬息以永恒，于长夜留余温。**
EvernightMoments 是一款通过提取照片原始拍摄时间，为您自动重命名影像文件的工具。

## 下载

最新版本: v1.0.0 (go 1.26.0)

前往 [Releases](https://github.com/kagurazakayashi/EvernightMoments/releases) 下载最新版本。

| 操作系统 |    处理器     | 位数 | 软件压缩包名称                    |
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

### 快速开始

1. 在 Windows 中：
   - 在“文件资源管理器”中，可以将照片文件直接拖拽到这个 .exe 文件上，开始重命名操作。
   - 你还可以将程序放到“发送到”菜单里面，可以随时通过在照片上点右键中的“发送到”菜单来使用本程序。
     - 按 `Windows+R` 快捷键打开“运行”对话框，然后输入 `shell:sendto` 并回车，将程序复制到新打开的文件夹窗口中。
   - 在“命令提示符”中，参考下面的其他平台操作方式进行操作即可。
2. 在 macOS / Linux 中：
   - 使用命令 `./EvernightMoments [照片文件路径1] [照片文件路径2] ...`

如果需要更改 语言、文件重命名的格式、交互询问开关 等，请先进行软件配置：

### 软件配置

**不带参数**直接运行 `./EvernightMoments` (或者在 Windows 中直接双击 .exe )，将会进入配置模式。

根据提示信息，回答每个问题并回车。
你可以进行以下设置：

#### 1. 语言设置

- 首先将会让你选择显示的语言，输入序号并回车即可。
- 设置被最终保存后，下次再进入设置时此处直接按回车跳过即可。

#### 2. 配置文件重命名的格式

- 你可以看到当前的文件重命名的格式，如果没有问题，可以直接留空并回车，将跳过修改。
- 如果需要修改，请输入新的格式。在格式中，可以使用以下代码代指重命名时的信息。
- 注意：区分大小写。
- **注意：不要输入系统不支持的符号！** 包括以下符号: `\  /  :  ?  "  <  >  |  \0`
  - 其中 Windows 中最终命名结果不能为 `CON, PRN, AUX, NUL, COM1~COM9, LPT1~LPT9`, 文件名结尾不能是 `.`
  - 使用系统不支持的符号可能导致重命名失败甚至损坏其他文件或文件系统。

| 代码   |    输出示例 | 含义         |
| :----- | ----------: | ------------ |
| `YY`   |        `25` | 两位年份     |
| `YYYY` |      `2025` | 完整年份     |
| `M`    |         `5` | 月份         |
| `MM`   |        `05` | 两位月份     |
| `D`    |         `2` | 日           |
| `DD`   |        `02` | 两位日       |
| `H`    |         `9` | 小时         |
| `HH`   |        `09` | 两位小时     |
| `m`    |         `7` | 分钟         |
| `mm`   |        `07` | 两位分钟     |
| `s`    |         `3` | 秒           |
| `ss`   |        `03` | 两位秒       |
| `#`    |         `1` | 编号         |
| `##`   |        `01` | 统一位数编号 |
| `*`    | `photo.jpg` | 原始文件名   |

注意，不支持 `hh` (12 小时格式)

示例：原始文件名为 `photo.jpg` 时，采用默认值格式 `YYYYMMDD_HHmmss*` ，将会输出 `20260220_122937_Photo.jpg` 。

更多示例：

- `YYYYMMDD_HHmmss*` -> `20250502_090703_Photo.jpg`
- `YY年M月D日—*` -> `25年5月2日-Photo.jpg`
- `YYYY-MM-DD_HH-mm-ss_*` -> `2025-05-02_09-07-03_Photo.jpg`
- `MM-DD-YYYY_HH-mm-ss_*` -> `05-02-2025_09-07-03_Photo.jpg`
- `DD.MM.YYYY_HH.mm.ss_*` -> `02.05.2025_09.07.03_Photo.jpg`

#### 3. 是否在重命名前预览确认？

- 如果开启，会先向你展示会修改成什么样子，让你确认要不要继续。
- 请输入 `y` 或者 `n`。
  - 默认值`y`: 每次先询问我。
  - `n`: 直接开始重命名。

#### 4. 在运行结束后，需要提示“按回车键退出”以便停留查看结果吗？

- 请输入 `y` 或者 `n`。
  - 默认值`y`: 结束后等待用户按回车键。
  - `n`: 结束后直接退出。

## 询问和预览

### 批量操作

1. 直接运行进入配置模式，将上面的第 3~4 项全部回答 `n` 来关闭交互模式。
2. 执行命令 `cd 你要批量操作的照片文件夹名称` 进入到要处理的文件夹。
3. 执行操作命令:

Windows:

```bat
FOR %i IN (*.ARW *.JPG *.CR3 *.HEIF) DO "绝对路径\EvernightMoments.exe" "%i"
```

macOS / Linux :

```bash
for f in *.ARW *.JPG *.CR3 *.HEIF; do "绝对路径/EvernightMoments" "$f"; done
```

## LICENSE

Copyright (c) 2026 KagurazakaYashi EvernightMoments is licensed under Mulan PSL v2. You can use this software according to the terms and conditions of the Mulan PSL v2. You may obtain a copy of Mulan PSL v2 at: http://license.coscl.org.cn/MulanPSL2 THIS SOFTWARE IS PROVIDED ON AN “AS IS” BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE. See the Mulan PSL v2 for more details.
