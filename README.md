# Go-create

Binary files can be downloaded from this link:
https://github.com/Juev/go-create/releases/latest

## Using

	$ go-create --help
    usage: go-create-darwin-amd64 [<flags>] <name>

    Flags:
      -h, --help            Show context-sensitive help (also try --help-long and --help-man).
      -n, --number="0.0.1"  Project version.
          --version         Show application version.

    Args:
      <name>  Project name.

And you can provide new values from command line, please look to help message. Values from command line is higher priority.

By default, go-create will create new directory with your project name.

For directory structure used repo: [github.com/Juev/go-scratch](https://github.com/Juev/go-scratch) without LICENSE file.

## Updating templates

For updating templates you can use go-bindata%

    $ brew install go-bindata

After this update files in `templates` directory and run command:

    $ go-bindata -o src/data.go -prefix templates templates

File data.go will be updated, run `make` for creating new binary files for go-create. It's all.
