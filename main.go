package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"text/template"

	"github.com/BurntSushi/toml"
	homedir "github.com/mitchellh/go-homedir"
	gitconfig "github.com/tcnksm/go-gitconfig"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	pname     = kingpin.Arg("name", "Project name.").Required().String()
	pversion  = kingpin.Flag("number", "Project version.").Short('n').Default("0.0.1").String()
	pgithub   = kingpin.Flag("github", "Github username.").Short('g').String()
	pusername = kingpin.Flag("username", "Username.").Short('u').String()
	plocal    = kingpin.Flag("local", "Use current directory for creating.").Short('l').Bool()
)

// ProjectConfig type
type ProjectConfig struct {
	Name     string
	Version  string
	Github   string
	Username string
}

func check(msg string, e error) {
	if e != nil {
		log.Fatalf("%s: %s", msg, e)
	}
}

func main() {
	kingpin.Version("Version: " + version + "\nBuildTime: " + buildTime + "\nCommit: " + commit)
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	var config ProjectConfig

	config.Name = *pname
	config.Version = *pversion

	home, err := homedir.Dir()
	check("Cannot get home directory: ", err)

	config.Github, _ = gitconfig.Global("github.user")
	config.Username, _ = gitconfig.Global("user.name")

	configFile := path.Join(home, ".gocreate")
	if _, err = os.Stat(configFile); err == nil {
		_, err = toml.DecodeFile(configFile, &config)
		check("Cannot decode config file: ", err)
	}

	if *pgithub != "" {
		config.Github = *pgithub
	}
	if *pusername != "" {
		config.Username = *pusername
	}

	if config.Github == "" && config.Username == "" {
		log.Fatal("Cannot get info about Github or Username. Pls use ~/.gitconfig file, or ~/.gocreate file for it.")
	}

	projectDir := config.Name
	if !*plocal {
		gopath := os.Getenv("GOPATH")
		projectDir = path.Join(gopath, "src", "github.com", config.Github, config.Name)
	}

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
	fmt.Printf("Directory was created: %s\nGithub username: %s\nVersion: %s", projectDir, config.Github, config.Version)
}
