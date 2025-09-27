setlocal enabledelayedexpansion
go generate
go build -gcflags="all=-N -l" -v -o EvernightMoments.exe
EvernightMoments.exe TestPhotos
DIR TestPhotos
