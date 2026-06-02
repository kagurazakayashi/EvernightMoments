![icon](ico/icon.ico)

# [EvernightMoments](https://github.com/kagurazakayashi/EvernightMoments)

[English](README.md) | [简体中文](README.zh-Hans.md) | [繁體中文](README.zh-Hant.md) | **日本語**

**瞬きに永遠を、常夜に温もりを。**

EvernightMoments は、写真の撮影日時を抽出し、ファイル名を自動でリネームするツールです。

## デモ

デモに使用した設定：

1. リネーム形式：`<YYYY>-<MM>-<DD>_<HH>-<mm>-<ss>_<*>`
2. プレビュー確認：無効
3. 終了時に一時停止：無効

### デモ 1 (Windows グラフィカル画面)

Windows 10 Pro 22H2

![Demo1](ico/demo1.gif)

### デモ 2 (Bash コマンドターミナル)

UOS 1072 Pro

![Demo1](ico/demo2.gif)

## ダウンロード

最新バージョン: v1.2.0

[Releases](https://github.com/kagurazakayashi/EvernightMoments/releases) から最新バージョンをダウンロードしてください。

|   OS    |  プロセッサ   | ビット | ソフトウェアパッケージ名               |
| :-----: | :-----------: | :----: | -------------------------------------- |
| windows | Intel/AMD x86 |   32   | `EvernightMoments_v*_windows-x86.7z`   |
| windows | Intel/AMD x86 |   64   | `EvernightMoments_v*_windows-x64.7z`   |
| windows |      ARM      |   64   | `EvernightMoments_v*_windows-arm64.7z` |
|  macOS  |   Intel x86   |   64   | `EvernightMoments_v*_macos-x64.7z`     |
|  macOS  | Apple silicon |   64   | `EvernightMoments_v*_macos-arm64.7z`   |
|  Linux  | Intel/AMD x86 |   32   | `EvernightMoments_v*_linux-x86.7z`     |
|  Linux  | Intel/AMD x86 |   64   | `EvernightMoments_v*_linux-x64.7z`     |
|  Linux  |      ARM      |   32   | `EvernightMoments_v*_linux-arm32.7z`   |
|  Linux  |      ARM      |   64   | `EvernightMoments_v*_linux-arm64.7z`   |

`ExifTool` 付きのパッケージには、ExifTool 実行ファイルが同梱されています。ExifTool はオプションですが、より幅広いメディアファイル形式や最新機種のサポートが可能になります。詳細は [認識機能を向上させるための ExifTool のインストール](#認識機能を向上させるための-exiftool-のインストール) をご確認ください。

このプログラムは以下のプラットフォームで動作確認済みです：

- Windows 11 25H2 (Intel & AMD x86_64)
- UOS 1072 (Intel & AMD x86_64)
- Arch Linux 6.16 (Intel & AMD x86_64)
- macOS 15.7 (Intel)

## インストールとアンインストール

### Windows でのインストール

管理者権限を必要とせずに書き込み可能な任意の場所に、直接解凍するだけで完了です。

#### 「送る」メニューへの登録

プログラムを「送る」メニューに追加すると、写真を右クリックしていつでも本プログラムを呼び出せるようになります。

1. `Windows+R` キーを押して「ファイル名を指定して実行」ダイアログを開き、`shell:sendto` と入力して Enter キーを押します。
2. 開いたフォルダの中に、プログラム（またはそのショートカット）をコピーします。
3. 1つまたは複数の **画像ファイル**、もしくは **画像のみが含まれるフォルダ** を選択して右クリックし、コンテキストメニューの「送る」項目から本プログラムを選択してください。

#### コマンドプロンプトから実行できるようにする

1. フォルダを作成し、その中にプログラムをコピーします。
2. そのフォルダのパスをコピーします。
3. `Windows+R` キーを押し、`rundll32.exe sysdm.cpl,EditEnvironmentVariables` と入力して Enter キーを押します。
4. 「Path」を選択し、「編集」ボタンをクリックします。
5. 「新規」ボタンをクリックし、コピーしたフォルダのパスを貼り付けます。

#### Windows でのアンインストール

- 関連するファイルおよび環境変数を削除するだけで完了です。

### Linux でのインストール

1. root 権限なしで書き込み可能な場所に解凍します。
2. **ターミナル**を開き、以下のコマンドを実行します：

```bash
cd 解凍したディレクトリのパス
chmod +x install.sh
./install.sh
```

- このスクリプトは以下の作業を自動的に行います：
  - プログラムを `~/.local/bin/EvernightMoments` にインストールします。
  - アイコンリソースを `~/.local/share/icons/EvernightMoments.png` にインストールします。
  - 設定プログラム用のアプリケーションリスト項目 `~/.local/share/applications/EvernightMoments.desktop` を作成し、「グラフィックス」カテゴリに分類します。
  - 設定プログラム用のデスクトップアイコン `~/Desktop/EvernightMoments.desktop` を作成します。
  - ターミナルでいつでも使用できるように、コマンドを `~/.local/bin` フォルダに追加します。
    - `~/.local/bin` が `PATH` 環境変数にない場合は、自動追加を試みます。
- **注意：一部のオペレーティングシステムでは、署名されていないプログラムの実行が拒否されます。** 本プログラムを使用するには、署名されていないプログラムの実行を許可する必要があります。通常、グラフィカルインターフェース上でシステムに実行が拒否された場合、ポップアップで警告が表示されます。
- プログラムの設定は `~/.config/EvernightMoments/config.json` に保存されます。

#### Linux でのアンインストール

```bash
cd 解凍したディレクトリのパス
chmod +x uninstall.sh
./uninstall.sh
```

その後、関連するすべてのファイルと環境変数を削除してください。

### macOS でのインストール

1. root 権限なしで書き込み可能な場所に解凍します。
2. **ターミナル**を開き、以下のコマンドを実行します：

```bash
cd 解凍したディレクトリのパス
chmod +x install.sh
./install-mac.sh
```

- このスクリプトは以下の作業を自動的に行います：
  - プログラムを `~/.local/bin/EvernightMoments` にインストールします。
  - 設定プログラム `~/Applications/EvernightMoments Config.app` を生成します。
  - 設定プログラム用のデスクトップアイコン `~/Desktop/EvernightMoments Config.app` を作成します。
    - それをデスクトップから `/Applications` (アプリケーションフォルダ) に移動できます：
    - `mv "$HOME/Desktop/EvernightMoments Config.app" "/Applications/EvernightMoments Config.app"`
  - ターミナルでいつでも使用できるように、コマンドを `~/.local/bin` フォルダに追加します。
    - `~/.local/bin` が `PATH` 環境変数にない場合は、自動追加を試みます。
- **注意：システムが署名されていないプログラムの実行を拒否する場合があります。** 本プログラムを使用するには、署名されていないプログラムの実行を許可する必要があります。
  - ブロックされた場合は、「システム設定」の「プライバシーとセキュリティ」を開き、一番下までスクロールして「このまま開く」または「許可」ボタンをクリックしてください。
- プログラムの設定は `~/Library/Application Support/EvernightMoments/config.json` に保存されます。

#### macOS でのアンインストール

```bash
cd 解凍したディレクトリのパス
chmod +x uninstall.sh
./uninstall-mac.sh
```

その後、関連するすべてのファイルと環境変数を削除してください。

### 認識機能を向上させるための ExifTool のインストール

本プログラムに内蔵されている EXIF パーサーは [goexif](https://github.com/rwcarlsen/goexif) ですが、これは限られたフォーマットしか処理できません。また、このライブラリおよび本プログラムの更新制限により、新しいカメラデバイスに対するタイムリーなサポートを提供することができません。最新のカメラサポートを得るために、ExifTool をインストールすることを強くお勧めします。

[ExifTool](https://github.com/exiftool/exiftool) は Phil Harvey 氏によって開発されたフリーでオープンソースのソフトウェアであり、画像、動画、音声のメタデータ処理に特化しています。このプログラムを最新の状態に保つことで、本プログラムに最新のカメラ RAW サポートを提供し、より多くの種類のファイルフォーマットを処理できるようになります。

本プログラムが利用可能な ExifTool を検出した場合、開始時の出力に「EXIF エクストラクター」のファイルパスが表示されます。また、リネーム時に ExifTool が正常に日時を取得できた場合、「時間ソース」の項目に「ExifTool」と表示されます。

ExifTool のダウンロードおよびインストール手順について：

#### パッケージマネージャーを使用した ExifTool のインストール

システムにパッケージマネージャーがインストールされている場合は、使い慣れたパッケージマネージャーを使用して素早くインストールを完了できます。例：

- Windows: `choco install exiftool` または `scoop install exiftool`
- macOS: `brew install exiftool`
- Debian / Ubuntu / Mint: `sudo apt update && sudo apt install perl libimage-exiftool-perl`
- CentOS / RHEL / Fedora: `sudo dnf install perl perl-Image-ExifTool`
- Arch Linux: `sudo pacman -S perl perl-image-exiftool`

パッケージマネージャーがない場合は、以下の手順でインストールできます：

#### ExifTool のダウンロードとインストール

まず、[ExifTool のホームページ](https://github.com/exiftool/exiftool) にアクセスし、お使いのシステムに対応する最新バージョンのプログラムをダウンロードしてください。

- **Windows** の場合は2つの方法があります:
  - **グローバルインストール**: 公式の[インストールおよびアンインストール手順](https://exiftool.org/install.html#Windows)を参照して操作してください。完了後、「コマンドプロンプト」の任意の場所で `exiftool.exe` コマンドが実行できることを確認してください。
  - **本プログラムでのみ使用する場合**: zip ファイルがダウンロードされます。その zip ファイル内のすべての ExifTool ファイルを本プログラムのフォルダー内に展開します（`exiftool(-k).exe` または `exiftool.exe` が `EvernightMoments.exe` と同じフォルダー内にあることを確認してください）。
- **macOS**: 公式の[インストールおよびアンインストール手順](https://exiftool.org/install.html#MacOS)を参照して操作してください。
- **Linux**: 公式の[インストールおよびアンインストール手順](https://exiftool.org/install.html#Unix)を参照して操作してください。

## 使用方法

### Windows の GUI で使用する

- 「エクスプローラー」で、1 つまたは複数の **写真ファイル**、または **写真のみが含まれるフォルダー** をこの .exe ファイルに直接ドラッグ＆ドロップすると、リネーム操作が開始されます。
- 「送る」メニューから使用することもできます。詳細は [Windows でのインストール](#windows-でのインストール) を参照してください。

### 共通コマンド

`[実行ファイル名] [写真パス1] [写真パス2] ...`

- **プログラムの実行ファイル名**:
  - 環境変数が設定済みの場合：
    - 任意のディレクトリから直接 `EvernightMoments` を使用。
  - プログラムのあるディレクトリから実行する場合：
    - Windows コマンドプロンプトでは `EvernightMoments.exe` のようになります
    - `Windows PowerShell` では `.\EvernightMoments.exe` のようになります
    - macOS / Linux `sh` では `./EvernightMoments` のようになります
- **写真のファイルパス**:
  - **複数のファイル** に対応：例 `EvernightMoments.exe "C:\album1\photo1.arw" "C:\album1\photo2.arw"`
  - **複数のフォルダ** に対応：例 `EvernightMoments.exe "C:\album1" "C:\album2"`
    - フォルダを指定した場合、そのフォルダ内の **すべてのファイル** のリネームを試みます。
     - デフォルトではサブフォルダ内のファイルは変更されません。サブフォルダも含める場合は、引数 `-r` を追加してください。

### コマンドラインパラメータ

コマンドラインフラグを使用して、設定ファイルの任意の項目を一時的に上書きできます。各パラメータには短い形式（`-x`）と長い形式（`--xxx`）があります。ファイル/フォルダのパスにはフラグ名は不要で、末尾に配置します。複数のファイル、フォルダ、ワイルドカードをサポートします。

| 短形式 | 長形式             | 説明                                                 | 例                                              |
| :----- | :----------------- | :--------------------------------------------------- | :---------------------------------------------- |
| `-l`   | `--language`       | 表示言語を設定（`en`/`zh-Hans`/`zh-Hant`/`ja`）        | `-l en`                                         |
| `-f`   | `--format`         | リネーム形式テンプレートを設定                           | `-f "<YYYY>-<MM>-<DD>_<*>""`                    |
| `-e`   | `--exclude`        | 除外 glob パターンを追加（繰り返し可能）                 | `-e "*.dop" -e "*.cos"`                         |
| `-s`   | `--sync`           | 同期 glob パターンを追加（繰り返し可能）                 | `-s "*.dop"`                                    |
| `-y`   | `--confirm`        | プレビュー確認を有効化                                  | `-y`                                            |
| `-ny`  | `--no-confirm`     | プレビュー確認を無効化                                  | `-ny`                                           |
| `-p`   | `--pause`          | 終了前の一時停止を有効化                                | `-p`                                            |
| `-np`  | `--no-pause`       | 終了前の一時停止を無効化                                | `-np`                                           |
| `-x`   | `--exiftool`       | ExifTool 実行ファイルのパスを設定（空で無効化）          | `-x "C:\Tools\exiftool.exe"`                    |
| `-r`   | `--recursive`      | サブディレクトリを再帰的に処理                           | `-r`                                            |
| `-me`  | `--multi-ext`      | 複数レベルの拡張子をファイル名の一部として扱う              | `-me`                                           |
| `-nc`  | `--no-color`       | カラー端末出力を無効化                                  | `-nc`                                           |

**例：**

```bash
# フォーマットと言語を上書き、確認を無効化、再帰的に処理
EvernightMoments -f "<YYYY>-<MM>-<DD>_<*>" -l en -ny -r "C:\Photos"

# ExifTool パスを上書き、除外パターンを追加
EvernightMoments -x "D:\exiftool.exe" -e "*.dop" -e "*.cos" "C:\album1" "C:\album2"

# 複数フォルダの .ARW ファイルをリネームし、.ARW.dop サイドカーを同期
EvernightMoments -f "<YYYY><MM><DD>_<HH><mm><ss><*>" -s "*.ARW.dop" -ny "D:\DCIM\10860213" "D:\DCIM\11051228"
```

### AI エージェント統合

本プロジェクトには [`SKILL.md`](SKILL.md) ファイルが含まれており、AI エージェント向けに最適化された形式でツールのアーキテクチャ、CLI インターフェース、設定スキーマ、典型的なワークフローを説明しています。AI アシスタントに EvernightMoments を理解して操作させるには、このファイルをスキルとして読み込むか、コンテキストとして提供してください。

### ソフトウェア設定

言語、リネーム形式、除外パターン、同期パターン、ExifTool パス、確認プロンプトのオン/オフなどを変更したい場合は、まず設定を行ってください。

**引数なし** で `./EvernightMoments` を直接実行する（Windowsの場合は .exe をダブルクリックする）と、全画面の TUI 設定画面が開きます。

画面の操作方法：

- `Tab` または `方向キー` で各項目間のフォーカスを移動します。
- 入力欄には直接入力できます。トグル項目は `Enter` または `スペース` で切り替えます。
- 完了したら **「保存して終了」** ボタンで設定を書き込みます。`Esc` を押すか **「保存せずに終了」** を選ぶと、すべての変更が破棄されます。

以下の設定が可能です：

### 1. 言語設定

- 「言語」ドロップダウンから表示言語を選択すると、画面の文字が**即座に**選択した言語へ切り替わります。
- 選択内容は他の設定と一緒に保存され、次回起動時に自動で選択されます。

### 2. リネーム形式の設定

- 「命名形式」入力欄でリネーム形式を編集します。下部の情報エリアに生成結果の**ライブプレビュー**が表示され、不正な文字は入力時にブロックされます。
- 以下の **置換記号** を使用できます。
- 注意：大文字と小文字を区別します。

| 置換記号 |      出力例 | 意味           |
| :------- | ----------: | :------------- |
| `<YY>`   |        `25` | 2桁の年        |
| `<YYYY>` |      `2025` | 4桁の年        |
| `<M>`    |         `5` | 月             |
| `<MM>`   |        `05` | 2桁の月        |
| `<D>`    |         `2` | 日             |
| `<DD>`   |        `02` | 2桁の日        |
| `<H>`    |         `9` | 時             |
| `<HH>`   |        `09` | 2桁の時        |
| `<m>`    |         `7` | 分             |
| `<mm>`   |        `07` | 2桁の分        |
| `<s>`    |         `3` | 秒             |
| `<ss>`   |        `03` | 2桁の秒        |
| `<#>`    |         `1` | 連番           |
| `<##>`   |        `01` | 桁を揃えた連番 |
| `<*>`    | `photo.jpg` | 元のファイル名 |

#### 注意事項

1. **システムでサポートされていない記号は入力しないでください！** 例: `\ / : ? ' " | < * >`
   - Windowsの場合、ファイル名を `CON, PRN, AUX, NUL, COM1~COM9, LPT1~LPT9` にすることはできず、末尾を `.` にすることもできません。
   - 不適切な記号を使用すると、リネームの失敗や、他のファイル・ファイルシステムの破損を招く恐れがあります。
2. `<hh>` (12時間表記) には対応していません。

#### 例

元のファイル名が `photo.jpg` の時、デフォルト形式 `<YYYY><MM><DD>_<HH><mm><ss>_<*>` を使用すると、 `20260220_122937_Photo.jpg` が出力されます。

その他の例：

- `<YYYY><MM><DD>_<HH><mm><ss>*` -> `20250502_090703Photo.jpg`
- `<*>_<HH><mm><ss>` -> `Photo_193030.jpg`
- `<YY>年<M>月<D>日—*` -> `25年5月2日-Photo.jpg`
- `<YYYY>-<MM>-<DD>_<HH>-<mm>-<ss>_<*>` -> `2025-05-02_09-07-03_Photo.jpg`
- `<MM>-<DD>-<YYYY>_<HH>-<mm>-<ss>_<*>` -> `05-02-2025_09-07-03_Photo.jpg`
- `<DD>.<MM>.<YYYY>_<HH>.<mm>.<ss>_<*>` -> `02.05.2025_09.07.03_Photo.jpg`

### 3. 除外パターン（これらのファイルをスキップ）

- 「除外」入力欄に除外するパスパターンをカンマ区切りで入力します。絶対パスと相対パスをサポートし、拡張子のみに限りません。例：
  - `CaptureOne\Settings153\*.cos`
  - `*.dop`
- 一致したファイルはリネームされません。空欄の場合は何も除外しません。

### 4. 同期パターン（主ファイルと一緒にリネーム）

- 「同期」入力欄に同期リネームするパスパターンをカンマ区切りで入力します。絶対パスと相対パスをサポートし、拡張子のみに限りません。例：
  - `CaptureOne\Settings153\*.cos`
  - `*.dop`
- 主写真ファイルがリネームされるとき、同じフォルダ内の一致する付随ファイルは**同じ**新しいファイル名でリネームされます。
- **複数拡張子の付随ファイル：** 複数の拡張子を持つファイル（例：`photo.ARW.dop`）の場合、ツールは自動的に中間の拡張子を除去して主ファイルを特定します（例：`photo.ARW` または `photo` に一致）。
- **注意：** 同期パターンに一致するファイルはデフォルトで除外対象にもなります（それらは自身のコンテンツに基づいて名前を計算しません）。

### 5. ExifTool パス

- 「ExifTool パス」入力欄に `exiftool` 実行ファイルのパスを指定すると、より正確な撮影日時を読み取れます。
- デフォルト値はプログラムが自動検出したパスです。検出されなかった場合は空欄になります。
- **空欄**にすると ExifTool を使用せず、内蔵解析のみを使用します。
- 下の **「自動検出」** ボタンでシステムの `PATH` から再検出できます。

### 6. リネーム前にプレビュー確認しますか？

- 「プレビュー確認」トグル：オンにすると、変更内容を事前に表示し、続行するかどうかを確認します。
- オフにすると、確認せずにすぐにリネームを開始します（操作にご注意ください）。

### 7. 終了後に「エンターキーで終了」の待機を表示しますか？

- 「終了待機」トグル：オンにすると、結果を確認できるよう終了時に画面を保持します。オフにすると終了後すぐに閉じます。

### 8. 複数レベルの拡張子をファイル名として扱いますか？

- 「複数拡張子」トグル：オンにすると最後の拡張子のみ除去し、中間層（例: `photo.ARW.dop` の `.ARW`）をファイル名の一部として保持します。オフの場合（デフォルト）はすべての拡張子を除去し、ベース名のみを保持します。

### 9. カラー出力を有効にしますか？

- 「カラー出力」トグル：オンにすると端末出力に ANSI カラーコードが含まれます。出力をリダイレクトする場合や AI エージェント経由でツールを使用する場合は、オフにすることを推奨します。

## コンパイル

まず [Go](https://go.dev/) をインストールしてください。バージョンは `1.26.0` 以上である必要があります。

### Windows システムでのコンパイル

1. `mdhtml.bat` を使用してヘルプドキュメントを作成します（`readme` フォルダ内に作成されます）。
2. `build.bat` を使用して各プラットフォーム向けにビルドします（`bin` フォルダ内に出力されます）。

#### テスト

- まず `TestPhotos` フォルダを作成し、その中にテスト用の写真ファイルをいくつか入れます。
- `conf.bat`: コンパイルして設定モードに入ります。
- `test_dir.bat`: コンパイルして `TestPhotos` フォルダ（全ファイルを含む）の処理をテストします。
- `test_files.bat`: 複数のファイル入力（`TestPhotos` 内のすべてのファイルを使用）をテストします。
- テスト時にデフォルトの形式で `TestPhotos` フォルダ内の写真の名前が変更された場合は、`python test_undo.py` を実行して名前変更を取り消すことができます。

### macOS / Linux システムでのコンパイル

上記と同様に、`.bat` を `.sh` に変更するだけです。

## LICENSE

Copyright (c) 2026 KagurazakaYashi. EvernightMoments is licensed under Mulan PSL v2. You can use this software according to the terms and conditions of the Mulan PSL v2. You may obtain a copy of Mulan PSL v2 at: http://license.coscl.org.cn/MulanPSL2. THIS SOFTWARE IS PROVIDED ON AN “AS IS” BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE. See the Mulan PSL v2 for more details.
