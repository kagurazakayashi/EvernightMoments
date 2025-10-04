import os


def batch_rename():
    """
    執行批次更名功能。
    此函式會掃描指定的資料夾，將檔名中最後一個底線 (_) 之前的內容移除，
    僅保留底線後的名稱作為新檔名。
    """
    # 取得目前工作目錄下的 "TestPhotos" 資料夾完整路徑
    folder_path = os.path.join(os.getcwd(), "TestPhotos")
    count = 0

    # 輸出目前正在處理的目錄路徑
    print(f"DIR: {folder_path}")

    # 遍歷該資料夾內的所有檔案與子目錄
    for filename in os.listdir(folder_path):
        # 如果檔名剛好跟目前的腳本名稱相同，則跳過不處理
        if filename == os.path.basename(__file__):
            continue

        # 檢查檔名中是否包含底線
        if "_" in filename:
            # 以最後一個底線為基準進行分割，並取得最後一部分作為新檔名
            new_name = filename.rsplit("_", 1)[-1]

            # 組合出原始檔案與目標檔案的完整路徑
            old_file = os.path.join(folder_path, filename)
            new_file = os.path.join(folder_path, new_name)

            try:
                # 執行檔案更名動作
                os.rename(old_file, new_file)
                print(f"RENAME: {filename} -> {new_name}")
                count += 1
            except Exception as e:
                # 若更名過程發生錯誤（例如權限不足或檔案被佔用），捕捉例外並顯示
                print(f"ERR: {filename} : {e}")

    # 最終輸出成功處理的檔案總數
    print(f"OK: {count}")


if __name__ == "__main__":
    # 程式進入點：呼叫批次更名函式
    batch_rename()
