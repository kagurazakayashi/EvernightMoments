# [EvernightMoments](https://github.com/kagurazakayashi/EvernightMoments)

[English](README.md) | [简体中文](README.zh-Hans.md) | [繁體中文](README.zh-Hant.md) | **日本語**

**瞬きに永遠を、常夜に温もりを。**
EvernightMoments は、写真の撮影日時を抽出し、ファイル名を自動でリネームするツールです。

## ダウンロード

最新バージョン: v1.1.0 (go 1.26.0)

[Releases](https://github.com/kagurazakayashi/EvernightMoments/releases) から最新バージョンをダウンロードしてください。

|   OS    |  プロセッサ   | ビット | ソフトウェアパッケージ名          |
| :-----: | :-----------: | :----: | --------------------------------- |
| windows | Intel/AMD x86 |   32   | EvernightMoments_windows-x86.7z   |
| windows | Intel/AMD x86 |   64   | EvernightMoments_windows-x64.7z   |
| windows |      ARM      |   64   | EvernightMoments_windows-arm64.7z |
|  macOS  |   Intel x86   |   64   | EvernightMoments_macos-x64.7z     |
|  macOS  | Apple silicon |   64   | EvernightMoments_macos-arm64.7z   |
|  Linux  | Intel/AMD x86 |   32   | EvernightMoments_linux-x86.7z     |
|  Linux  | Intel/AMD x86 |   64   | EvernightMoments_linux-x64.7z     |
|  Linux  |      ARM      |   32   | EvernightMoments_linux-arm32.7z   |
|  Linux  |      ARM      |   64   | EvernightMoments_linux-arm64.7z   |

## 使い方

### クイックスタート

#### 基本コマンド

`[実行ファイル名] [写真パス1] [写真パス2] ...`

- **実行ファイル名**:
  - `Windows コマンドプロンプト` では `EvernightMoments.exe` など
  - `Windows PowerShell` では `.\EvernightMoments.exe` など
  - macOS / Linux `sh` では `./EvernightMoments` など
- **写真のファイルパス**:
  - **複数のファイル** に対応：例 `EvernightMoments.exe "C:\album1\photo1.arw" "C:\album1\photo2.arw"`
  - **複数のフォルダ** に対応：例 `EvernightMoments.exe "C:\album1" "C:\album2"`
    - フォルダを指定した場合、そのフォルダ内の **すべてのファイル** のリネームを試みます。
    - デフォルトではサブフォルダ内のファイルは変更されません。サブフォルダも含める場合は、引数 `-r` を追加してください。

#### WindowsのGUIで使用する場合

- 「エクスプローラー」で、1つ以上の **写真ファイル** または **写真のみが入ったフォルダ** を、この .exe ファイルの上に直接ドラッグ＆ドロップするとリネームが開始されます。
- また、プログラムを「送る」メニューに追加して、右クリックメニューからいつでも使用できるようにすることも可能です。
  - `Windows + R` キーで「ファイル名を指定して実行」ダイアログを開き、`shell:sendto` と入力してエンターキーを押します。開いたフォルダにプログラム（またはショートカット）をコピーしてください。
  - その後、写真やフォルダを右クリックし、「送る」メニューの中に本プログラムが表示されるようになります。

### ソフトウェア設定

言語、リネーム形式、確認プロンプトのオン/オフなどを変更したい場合は、まず設定を行ってください。

**引数なし** で `./EvernightMoments` を直接実行する（Windowsの場合は .exe をダブルクリックする）と、設定モードに入ります。

表示されるメッセージに従って回答し、エンターキーを押してください。
以下の設定が可能です：

#### 1. 言語設定

- 最初に表示言語を選択します。番号を入力してエンターキーを押してください。
- 設定保存後、次回設定に入る際はエンターキーでスキップできます。

#### 2. リネーム形式の設定

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

##### 注意事項

1. **システムでサポートされていない記号は入力しないでください！** 例: `\ / : ? ' " | < * >`
   - Windowsの場合、ファイル名を `CON, PRN, AUX, NUL, COM1~COM9, LPT1~LPT9` にすることはできず、末尾を `.` にすることもできません。
   - 不適切な記号を使用すると、リネームの失敗や、他のファイル・ファイルシステムの破損を招く恐れがあります。
2. `<hh>` (12時間表記) には対応していません。

##### 例

元のファイル名が `photo.jpg` の時、デフォルト形式 `<YYYY><MM><DD>_<HH><mm><ss>_<*>` を使用すると、 `20260220_122937_Photo.jpg` が出力されます。

その他の例：

- `<YYYY><MM><DD>_<HH><mm><ss>*` -> `20250502_090703Photo.jpg`
- `<YY>年<M>月<D>日—*` -> `25年5月2日-Photo.jpg`
- `<YYYY>-<MM>-<DD>_<HH>-<mm>-<ss>_<*>` -> `2025-05-02_09-07-03_Photo.jpg`

#### 3. リネーム前にプレビュー確認しますか？

- 有効にすると、変更内容が事前に表示され、続行するかどうかを確認できます。
- `y` または `n` を入力してください。
  - デフォルト `y`: 毎回確認する。
  - `n`: 確認せずにリネームを開始する。

#### 4. 実行終了後、「エンターキーで終了」の待機を表示しますか？

- 結果を確認するために画面を保持するかどうかを設定します。
- `y` または `n` を入力してください。
  - デフォルト `y`: 終了時にユーザーの入力を待つ。
  - `n`: 終了後すぐに閉じる。

## ビルド

まず [Go](https://go.dev/) をインストールしてください。バージョンは `1.26.0` 以上である必要があります。

### Windows システムの場合

以下のスクリプトを使用できます：

- **ビルド**
  - `build.bat`: 各プラットフォーム向けにビルドします（`bin` フォルダに出力されます）。
- **テスト**
  - まず `TestPhotos` フォルダを作成し、その中にテスト用の写真ファイルをいくつか入れます。
  - `conf.bat`: ビルドして設定モードに入ります。
  - `test_dir.bat`: ビルドして `TestPhotos` フォルダ（全ファイルを含む）の処理をテストします。
  - `test_files.bat`: 複数のファイル入力（`TestPhotos` 内の全ファイルを対象）をテストします。

### すべてのシステム

ソースコードのフォルダに `cd` で移動し、以下を実行します：

```bash
go generate
go build .
```

## LICENSE

Copyright (c) 2026 KagurazakaYashi. EvernightMoments is licensed under Mulan PSL v2. You can use this software according to the terms and conditions of the Mulan PSL v2. You may obtain a copy of Mulan PSL v2 at: http://license.coscl.org.cn/MulanPSL2. THIS SOFTWARE IS PROVIDED ON AN “AS IS” BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE. See the Mulan PSL v2 for more details.
