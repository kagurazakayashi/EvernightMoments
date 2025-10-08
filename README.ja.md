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

最新バージョン: v1.1.0

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
- プログラムの設定は `~/.local/bin/EvernightMoments.json` に保存されます。

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
- プログラムの設定は `~/.local/bin/EvernightMoments.json` に保存されます。

#### macOS でのアンインストール

```bash
cd 解凍したディレクトリのパス
chmod +x uninstall.sh
./uninstall-mac.sh
```

その後、関連するすべてのファイルと環境変数を削除してください。

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

### ソフトウェア設定

言語、リネーム形式、確認プロンプトのオン/オフなどを変更したい場合は、まず設定を行ってください。

**引数なし** で `./EvernightMoments` を直接実行する（Windowsの場合は .exe をダブルクリックする）と、設定モードに入ります。

表示されるメッセージに従って回答し、エンターキーを押してください。
以下の設定が可能です：

### 1. 言語設定

- 最初に表示言語を選択します。番号を入力してエンターキーを押してください。
- 設定保存後、次回設定に入る際はエンターキーでスキップできます。

### 2. リネーム形式の設定

- 現在のリネーム形式が表示されます。問題なければ空欄のままエンターキーを押してスキップしてください。
- 変更する場合は、新しい形式を入力してください。以下の **置換記号** を使用できます。
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

### 3. リネーム前にプレビュー確認しますか？

- 有効にすると、変更内容が事前に表示され、続行するかどうかを確認できます。
- `y` または `n` を入力してください。
  - デフォルト `y`: 毎回確認する。
  - `n`: 確認せずにリネームを開始する。

### 4. 実行終了後、「エンターキーで終了」の待機を表示しますか？

- 結果を確認するために画面を保持するかどうかを設定します。
- `y` または `n` を入力してください。
  - デフォルト `y`: 終了時にユーザーの入力を待つ。
  - `n`: 終了後すぐに閉じる。

## コンパイル

まず [Go](https://go.dev/) をインストールしてください。バージョンは `1.26.0` 以上である必要があります。

### Windows システムでのコンパイル

`build.bat` スクリプトを使用して、各プラットフォーム向けにコンパイルできます（`bin` フォルダに出力されます）。

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
