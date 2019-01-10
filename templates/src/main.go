package main

import (
	"fmt"
)

var (
	buildTime = "unset"
	commit    = "unset"
	version   = "unset"
	branch    = "unset"
)

func main() {
	fmt.Println("Version: ", version, " BuildTime: ", buildTime, " Commit: ", commit)
}
