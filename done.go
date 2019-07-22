package main

import (
	"fmt"
	"log"

	"github.com/thinkofher/gotask/db"
	"github.com/urfave/cli"
)

var done = cli.Command{
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
