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

// TODO: Provide same cli interfaces for adding
//       and deleting tasks (tag, id)

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
				cli.IntSliceFlag{
					Name:  "id, i",
					Usage: "Show tasks with given ids.",
				},
				cli.StringSliceFlag{
					Name:  "tag, t",
					Usage: "Show tasks, which contains given tag. You choose multiple tags.",
				},
			},
			Action: func(c *cli.Context) error {
				var checkers []db.Checker

				// Add id checkers
				for _, id := range c.IntSlice("id") {
					checkers = append(checkers, db.IdChecker(id))
				}

				// Add tag checkers
				for _, tag := range c.StringSlice("tag") {
					checkers = append(checkers, db.TagChecker(tag))
				}

				tasks, err := db.GetAllTasks()
				if err != nil {
					return err
				}

				// If len of checkers is equal to 0
				// then show all tasks
				// TODO: instead of re-using lines of code
				//       create a functions for displaying tasks
				if len(checkers) == 0 {
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
						os.Exit(0)
					}
				} else {
					tasks = db.TaskSelection(tasks, checkers...)
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
						fmt.Println("There aren't such tasks with given conditions.")
						os.Exit(1)
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

				// TODO: make use of Checkers

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
