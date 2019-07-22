package main

import (
	"fmt"
	"os"

	"github.com/thinkofher/gotask/db"
	"github.com/urfave/cli"
)

var done = cli.Command{
	Name:    "done",
	Aliases: []string{"d"},
	Usage:   "Complete task with given id.",
	Flags: []cli.Flag{
		cli.IntSliceFlag{
			Name:  "id, i",
			Usage: "Delete tasks with given global ids.",
		},
		cli.StringSliceFlag{
			Name:  "tag, t",
			Usage: "Delete tasks with given tags.",
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "Print additional information.",
		},
	},
	Before: parseCheckers,
	Action: func(c *cli.Context) error {
		var err error

		if len(tasks) != 0 {
			for _, val := range tasks {
				err = db.DeleteTask(val.ID)
				if err != nil {
					return err
				}
				if c.Bool("verbose") {
					fmt.Printf("Task \"%s\" with global id: %d is done.\n",
						val.Body, val.ID)
				}
			}
		} else if len(checkers) != 0 {
			fmt.Println("There aren't such tasks with given conditions.")
			os.Exit(1)
		} else {
			fmt.Println("You have no tasks to done.\n" +
				"Try to add something with \"gotask add\".")
			os.Exit(0)
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
