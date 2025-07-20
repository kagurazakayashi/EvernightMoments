setlocal enabledelayedexpansion
go build -gcflags="all=-N -l" -v -o ./TestPhotos/EvernightMoments.exe
CD TestPhotos
SET "file_list="
@ECHO OFF
for %%f in (*.jpg *.arw *.cr3 *.nef) do (
    if defined file_list (
        SET "file_list=!file_list! %%f"
    ) else (
        SET "file_list=%%f"
    )
)
ECHO ON
EvernightMoments.exe !file_list!
DIR
CD ..
