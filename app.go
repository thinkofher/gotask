package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/thinkofher/gotask/db"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func prepareApp() {
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) > 0 {
			fmt.Println(
				"Type \"gotask\", \"gotask --help\"" +
					" or \"gotask help\" to get info about possible usage.")
		} else {
			err := cli.ShowAppHelp(c)
			if err != nil {
				log.Fatal(err)
			}

		}
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

var task db.Task
var tagSeparator string
var taskId int

func appCommands() {
	app.Commands = []cli.Command{
		{
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
				cli.StringFlag{
					Name:  "tags",
					Usage: "You can easly sort your tasks with tags",
				},
				cli.StringFlag{
					Name:        "sep",
					Usage:       "Character to separate tags",
					Value:       ",",
					Destination: &tagSeparator,
				},
			},
			Action: func(c *cli.Context) error {
				if c.NArg() > 0 {
					task.Body = c.Args().Get(0)
				}
				task.ParseTags(c.String("tags"), tagSeparator)
				task.SetCurrDate()

				err := db.AddTask(&task)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf(
					"Task %q added to your tasks list.\n",
					task)

				return nil
			},
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Show tasks in your tasks list",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "id",
					Usage:       "Show task with given id",
					Value:       1,
					Destination: &taskId,
				},
			},
			Action: func(c *cli.Context) error {
				showTask, err := db.GetTask(taskId)
				if err != nil {
					fmt.Println("There is no such a task.")
					os.Exit(1)
				}
				fmt.Printf("%v\n", showTask)
				return nil
			},
		},
		{
			Name:    "done",
			Aliases: []string{"d"},
			Usage:   "Complete task with given id.",
			Action: func(c *cli.Context) error {
				arg := c.Args().Get(0)
				id, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Printf("Invalid input: %q.\n", arg)
					os.Exit(1)
				}
				fmt.Printf("Done task with id: \"%d\".\n", id)
				return nil
			},
		},
	}
}
