package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"text/template"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	pname    = kingpin.Arg("name", "Specify project name.").String()
	pversion = kingpin.Flag("v", "Specify project version.").Default("0.0.1").String()

	buildTime = "unset"
	commit    = "unset"
	version   = "unset"
	branch    = "unset"
)

// ProjectConfig type
type ProjectConfig struct {
	Name    string
	Version string
}

func check(msg string, e error) {
	if e != nil {
		log.Fatalf("%s: %s", msg, e)
	}
}

func main() {
	kingpin.Version("Version: " + version + "\nBuildTime: " + buildTime + "\nCommit: " + commit + "\n")
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	if strings.Compare(*pname, "") == 0 {
		kingpin.Usage()
		os.Exit(0)
	}

	var config ProjectConfig

	config.Name = *pname
	config.Version = *pversion

	projectDir := config.Name

	for _, el := range AssetNames() {
		data, err := Asset(el)
		check("Unable to parse: ", err)
		tmpl, err := template.New("").Parse(string(data))
		check("Error during creating new template: ", err)
		err = os.MkdirAll(path.Join(projectDir, filepath.Dir(el)), os.FileMode(0755))
		check("Error during create directory: ", err)
		file, err := os.Create(path.Join(projectDir, el))
		check("Error during creating file: ", err)
		err = tmpl.Execute(file, config)
		check("Error during executing template: ", err)
	}
	fmt.Printf("Directory was created: %s\nVersion: %s\n", projectDir, config.Version)
}
