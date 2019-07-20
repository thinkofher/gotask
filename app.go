package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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
					Name:  "tags, t",
					Usage: "You can easly sort your tasks with tags",
				},
				cli.StringFlag{
					Name:        "sep, s",
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
					"Task \"%s\" added to your tasks list with id: %d.\n",
					task.Body, task.Id)

				return nil
			},
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Show tasks in your tasks list",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "full-info, f",
					Usage: "Check if you want to see full infor about tasks.",
				},
				cli.IntFlag{
					Name: "id, i",
					Usage: "Show task with given id, " +
						"leave id equal to 0 if you want to lists all tasks.",
					Value:       0,
					Destination: &taskId,
				},
			},
			Action: func(c *cli.Context) error {
				// 0 means to show all tasks
				if taskId == 0 {
					tasks, err := db.GetAllTasks()
					if err != nil {
						return err
					}
					if len(tasks) > 0 {
						stasks := make([]string, len(tasks))
						for i, val := range tasks {
							if c.Bool("full-info") {
								stasks[i] = fmt.Sprintf("Task no. %d)\n%v", i+1, val)
							} else {
								stasks[i] = fmt.Sprintf("Task no. %d) %s\n", i+1, val.Body)
							}
						}

						if c.Bool("full-info") {
							fmt.Println(strings.Join(stasks, "\n\n"))
						} else {
							fmt.Print(strings.Join(stasks, ""))
						}
					} else {
						fmt.Println("You have no tasks to show.\n" +
							"Try to add something with \"gotask add\".")
					}
					// Show only specific task
				} else {
					showTask, err := db.GetTask(taskId)
					if err != nil {
						fmt.Println("There is no such a task.")
						os.Exit(1)
					}

					if c.Bool("full-info") {
						fmt.Printf("%v\n", showTask)
					} else {
						fmt.Println(showTask.Body)
					}
				}
				return nil
			},
		},
		{
			Name:    "done",
			Aliases: []string{"d"},
			Usage:   "Complete task with given id.",
			Flags: []cli.Flag{
				cli.IntSliceFlag{
					Name:  "global-id, g",
					Usage: "Delete tasks with given global ids.",
				},
				cli.IntSliceFlag{
					Name:  "number, n",
					Usage: "Delete tasks with given local numbers.",
				},
				cli.BoolFlag{
					Name:  "verbose, v",
					Usage: "Print additional information.",
				},
			},
			Action: func(c *cli.Context) error {
				// global ids are use to place tasks
				// in a database
				globals := c.IntSlice("global-id")

				// by local ids we undersand ids
				// by which are tasks listed with
				// "gotask" show command
				locals := c.IntSlice("number")

				tasks, err := db.GetAllTasks()
				if err != nil {
					log.Fatal(err)
				}

				if len(globals) > 0 {
					for _, val := range tasks {
						if intInSlice(val.Id, globals) {
							err := db.DeleteTask(val.Id)
							if err != nil {
								fmt.Println("There is no such a task.")
							}
							if c.Bool("verbose") {
								fmt.Printf("Task with global id: %d is done.\n", val.Id)
							}
						}
					}
				}

				if len(locals) > 0 {
					for _, localId := range locals {
						// Check if there is a task with
						// given local number
						if localId <= len(tasks) {
							err := db.DeleteTask(tasks[localId-1].Id)
							if err != nil {
								return err
							}
							if c.Bool("verbose") {
								fmt.Printf("Task with number: %d is done.\n", localId)
							}
						} else {
							fmt.Printf(
								"There is no task with a number: %d.\n", localId)
						}
					}
				}

				return nil

			},
		},
	}
}

// Checks if given []int slice contains
// given integer.
func intInSlice(i int, list []int) bool {
	for _, val := range list {
		if val == i {
			return true
		}
	}
	return false
}
