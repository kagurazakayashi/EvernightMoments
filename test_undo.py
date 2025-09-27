import os
def batch_rename():
    folder_path = os.path.join(os.getcwd(), "TestPhotos")
    count = 0
    print(f"DIR: {folder_path}")
    for filename in os.listdir(folder_path):
        if filename == os.path.basename(__file__):
            continue
        if "_" in filename:
            new_name = filename.rsplit('_', 1)[-1]
            old_file = os.path.join(folder_path, filename)
            new_file = os.path.join(folder_path, new_name)
            try:
                os.rename(old_file, new_file)
                print(f"RENAME: {filename} -> {new_name}")
                count += 1
            except Exception as e:
                print(f"ERR: {filename} : {e}")
    print(f"OK: {count}")
if __name__ == "__main__":
    batch_rename()
