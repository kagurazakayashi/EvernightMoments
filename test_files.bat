setlocal enabledelayedexpansion
go generate
go build -gcflags="all=-N -l" -v -o EvernightMoments.exe
CD TestPhotos
SET "file_list="
@ECHO OFF
for %%f in (*.jpg *.arw *.cr3 *.nef) do (
    if defined file_list (
        SET "file_list=!file_list! TestPhotos/%%f"
    ) else (
        SET "file_list=TestPhotos/%%f"
    )
)
ECHO ON
CD ..
EvernightMoments.exe !file_list!
DIR TestPhotos
