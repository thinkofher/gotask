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
