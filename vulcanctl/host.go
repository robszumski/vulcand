package main

import (
	"github.com/codegangsta/cli"
)

func NewHostCommand() cli.Command {
	return NewGroupCommand(GroupCommand{
		Name:  "host",
		Usage: "Operations with vulcan hosts",
		Flags: flags(),
		Subcommands: []cli.Command{
			{
				Name:   "add",
				Flags:  flags(),
				Usage:  "Add a new host to vulcan proxy",
				Action: addHostAction,
			},
			{
				Name:   "rm",
				Flags:  flags(),
				Usage:  "Remove a host from vulcan",
				Action: deleteHostAction,
			},
		},
	})
}

func addHostAction(c *cli.Context) {
	err := client(c).AddHost(c.Args().First())
	if err != nil {
		printError(err)
	} else {
		printOk("Host added")
	}
}

func deleteHostAction(c *cli.Context) {
	err := client(c).DeleteHost(c.Args().First())
	if err != nil {
		printError(err)
	} else {
		printOk("Host deleted")
	}
}
