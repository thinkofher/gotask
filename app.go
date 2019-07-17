package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func prepareApp() {
	app.Action = func(c *cli.Context) error {
		fmt.Println("Type \"gotask --help\" or \"gotask help\" to get info about possible usage.")
		return nil
	}
	appInfo()
	appCommands()
}

func appInfo() {
	app.Name = "gotask"
	app.Usage = "Add, remove and edit tasks in your local database."
	app.Author = "Beniamin Dudek"
	app.Email = "beniamin.dudek@yahoo.com"
	app.Version = "0.0.1"
}

func appCommands() {
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add task to your tasks list",
			Action: func(c *cli.Context) error {
				task := strings.Join(c.Args(), " ")
				fmt.Printf(
					"Task %q added to your tasks list.\n",
					task)
				return nil
			},
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Show tasks in your tasks list.",
			Action: func(c *cli.Context) error {
				fmt.Println("Testing...")
				return nil
			},
		},
		{
			Name:    "done",
			Aliases: []string{"s"},
			Usage:   "Complete task with given id.",
			Action: func(c *cli.Context) error {
				id, err := strconv.Atoi(c.Args().Get(0))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Done task with id: \"%d\".\n", id)
				return nil
			},
		},
	}
}
