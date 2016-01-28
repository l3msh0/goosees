package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// "goosees dev groupA up" のように実行する

// global options. available to any subcommands.
var prmConfPath string
var prmGroup string

func parseConfPath(name string) (confPath string, err error) {
	if fileExists(name) {
		return name, nil
	}
	if fileExists(name + ".yml") {
		return name + ".yml", nil
	}
	if fileExists(name + ".yaml") {
		return name + ".yaml", nil
	}
	return "", errors.New(fmt.Sprintf("conf file \"%s\" not found", name))
}

var commands = []*Command{
	upCmd,
	downCmd,
	redoCmd,
	statusCmd,
	createCmd,
	dbVersionCmd,
}

func main() {

	args := os.Args
	if len(args) < 4 || args[0] == "-h" {
		usage()
		return
	}

	var err error
	prmConfPath, err = parseConfPath(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	prmGroup = args[2]

	var cmd *Command
	name := args[3]
	for _, c := range commands {
		if strings.HasPrefix(c.Name, name) {
			cmd = c
			break
		}
	}

	if cmd == nil {
		fmt.Printf("error: unknown command %q\n", name)
		usage()
		os.Exit(1)
	}

	cmd.Exec(args[4:])
}

func usage() {
	fmt.Print(usagePrefix)
	usageTmpl.Execute(os.Stdout, commands)
}

var usagePrefix = `
Goosees is a wrapper of goose for applying migration to multiple databases.

Usage:
    goosees <conf> <group> <subcommand> [subcommand options]

Options:
`
var usageTmpl = template.Must(template.New("usage").Parse(
	`
Commands:{{range .}}
    {{.Name | printf "%-10s"}} {{.Summary}}{{end}}
`))
