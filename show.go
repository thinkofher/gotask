package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/thinkofher/gotask/db"
	"github.com/urfave/cli"
)

var show = cli.Command{
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
	Before: parseCheckers,
	Action: func(c *cli.Context) error {
		if len(tasks) != 0 {
			visTasks(tasks, c.Bool("full-info"))
			os.Exit(0)
		} else if len(checkers) != 0 {
			fmt.Println("There aren't such tasks with given conditions.")
			os.Exit(1)
		} else {
			fmt.Println("You have no tasks to show.\n" +
				"Try to add something with \"gotask add\".")
			os.Exit(0)
		}
		return nil
	},
}

// Visualize given slice of Tasks with
// short or full information
func visTasks(tasks []db.Task, full bool) {
	stasks := make([]string, len(tasks))
	for i, val := range tasks {
		if full {
			stasks[i] = fmt.Sprintf("Task no. %d)\n%v", i+1, val)
		} else {
			stasks[i] = fmt.Sprintf("Task no. %d) %s\n", i+1, val.Body)
		}
	}
	if full {
		fmt.Println(strings.Join(stasks, "\n\n"))
	} else {
		fmt.Print(strings.Join(stasks, ""))
	}
}
