package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	kingpin.Version("Version: " + version + " BuildTime: " + buildTime + " Commit: " + commit)
	kingpin.Parse()

	RestoreAssets("asd", "")
}
