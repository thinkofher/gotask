package main

import (
	"github.com/thinkofher/gotask/db"
	"github.com/urfave/cli"
)

var tasks []db.Task
var checkers []db.Checker

// Use it in Before field in Command struct
// You have to provide flags with id and tag names
var parseCheckers = func(c *cli.Context) error {
	var tasksp = &tasks
	var checkersp = &checkers
	var err error

	// Add id checkers
	for _, id := range c.IntSlice("id") {
		*checkersp = append(*checkersp, db.IdChecker(id))
	}

	// Add tag checkers
	for _, tag := range c.StringSlice("tag") {
		*checkersp = append(*checkersp, db.TagChecker(tag))
	}

	*tasksp, err = db.GetAllTasks()
	if err != nil {
		return err
	}

	if len(*checkersp) != 0 {
		*tasksp = db.TaskSelection(*tasksp, *checkersp...)
	}

	return nil
}
