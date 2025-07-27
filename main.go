package main

import (
	"os"
)

const defaultFormat = "YYYYMMDD_HHmmss_*"
const outLine = "--------------------------------"
const evernightMoments = "EvernightMoments"
const evernightMomentsVersion = "1.0.0"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		runConfigMode()
		return
	}
	runRenameMode(args)
}
