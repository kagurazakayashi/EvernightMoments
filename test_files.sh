#!/bin/bash
go generate
go build -gcflags="all=-N -l" -v -o EvernightMoments
cd TestPhotos || exit
file_list=()
shopt -s nullglob
for f in *.jpg *.arw *.cr3 *.nef *.JPG *.ARW *.CR3 *.NEF; do
    file_list+=("TestPhotos/$f")
done
shopt -u nullglob
cd ..
./EvernightMoments "${file_list[@]}"
ls -l TestPhotos
