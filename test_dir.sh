#!/bin/bash
go generate
go build -gcflags="all=-N -l" -v -o EvernightMoments
./EvernightMoments TestPhotos
ls -l TestPhotos
