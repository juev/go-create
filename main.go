package main

import (
	"log"
	"os"
	"path/filepath"

	"text/template"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	pname    = kingpin.Arg("name", "Project name.").Required().String()
	pversion = kingpin.Flag("number", "Project version.").Short('n').Default("0.0.1").String()
	pauthor  = kingpin.Flag("author", "Author name.").Short('a').Default(os.Getenv("USER")).String()
)

// Tname type
type Tname struct {
	Name    string
	Version string
	Author  string
}

func check(msg string, e error) {
	if e != nil {
		log.Panicf("%s: %s", msg, e)
	}
}

func cName(dir, filename string) string {
	return dir + string(filepath.Separator) + filename
}

func main() {
	kingpin.Version("Version: " + version + "\nBuildTime: " + buildTime + "\nCommit: " + commit)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	project := Tname{*pname, *pversion, *pauthor}
	for _, path := range AssetNames() {
		data, err := Asset(path)
		check("Unable to parse: ", err)
		tmpl, err := template.New("").Parse(string(data))
		check("Error during creating new template: ", err)
		err = os.MkdirAll(cName(*pname, filepath.Dir(path)), os.FileMode(0755))
		check("Error during create directory: ", err)
		file, err := os.Create(cName(*pname, path))
		check("Error during creating file", err)
		err = tmpl.Execute(file, project)
		check("Error during executing template: ", err)
	}
}
