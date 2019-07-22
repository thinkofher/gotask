package main

import (
	"fmt"
	"log"

	"github.com/thinkofher/gotask/db"
	"github.com/urfave/cli"
)

var task db.Task

var add = cli.Command{
	Name:    "add",
	Aliases: []string{"a"},
	Usage:   "Add task to your tasks list",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "body",
			Usage:       "Fill it with something you want or must do",
			Value:       "Sample task content.",
			Destination: &task.Body,
		},
		cli.StringSliceFlag{
			Name:  "tag, t",
			Usage: "Tag your task with any category you like (you can add multiple tags)",
		},
	},
	Action: func(c *cli.Context) error {
		if c.NArg() > 0 {
			task.Body = c.Args().Get(0)
		}
		task.Tags = c.StringSlice("tag")
		task.SetCurrDate()

		err := db.AddTask(&task)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"Task \"%s\" added to your tasks list with id: %d.\n",
			task.Body, task.Id)

		return nil
	},
}
