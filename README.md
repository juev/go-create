# Go-create

Binary files can be downloaded from this link:
https://github.com/Juev/go-create/releases/latest

## Using 

	$ go-create --help
	usage: go-create [<flags>] <name>

	Flags:
	-h, --help               Show context-sensitive help (also try --help-long and --help-man).
	-n, --number="0.0.1"     Project version.
	-g, --github=GITHUB      Github username.
	-u, --username=USERNAME  Username.
	-l, --local              Use current directory for creating.
		--version            Show application version.

	Args:
	<name>  Project name.

By default go-create used parameters from your environment: `~/.gitconfig` file and `$GOPATH` cariable. 

You can use `~/.gocreate` file for storing parameters, fo example:

	github = "Juev"
	username = "Denis Evsyukov"

And you can provide new values from command line, please look to help message. Values from command line is higher priority.

By default, go-create will create new directory with your project name on $GOPATH directory. But if you provide key `-l`, new directory will create in current local directory.

For directory structure used repo: [github.com/Juev/go-scratch](https://github.com/Juev/go-scratch) without LICENSE file.