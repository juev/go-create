package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	pname    = kingpin.Arg("name", "Project name.").Required().String()
	pversion = kingpin.Arg("version", "Project version.").Default("0.0.1").String()
	pauthor  = kingpin.Arg("author", "Author name.").Default(os.Getenv("USER")).String()
)

func main() {
	kingpin.Version("Version: " + version + "\nBuildTime: " + buildTime + "\nCommit: " + commit)
	kingpin.CommandLine.VersionFlag.Short('v')
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	RestoreAssets(*pname, "")
	fmt.Println("Name: ", *pname, " Version: ", *pversion, " Author: ", *pauthor)
}
