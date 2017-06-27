package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/uraura/karabiner-elements-profile-switcher/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "use",
		Usage:  "",
		Action: command.CmdUse,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
