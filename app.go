package main

import (
	"fmt"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/thinkofher/gotask/db"
	"github.com/urfave/cli"
)

// Path where database will be stored
var taskdbPath = ".tasks.db"

// Returns pointer to main cli application
func prepareApp() *cli.App {
	var app = cli.NewApp()

	app.Before = func(c *cli.Context) error {
		// Acquire the home directory
		// independently of OS
		home, err := homedir.Dir()
		if err != nil {
			return err
		}
		// Initalize the database if not exist
		dbPath := filepath.Join(home, taskdbPath)
		err = db.InitDB(dbPath)
		if err != nil {
			return err
		}
		return nil
	}
	app.Action = func(c *cli.Context) error {
		// App prompt information about possile usage
		// when user gives undefined flags or commands
		if len(c.Args()) > 0 {
			fmt.Println(
				"Type \"gotask\", \"gotask --help\"" +
					" or \"gotask help\" to get info about possible usage.")
		} else {
			// When typing just name of the app
			// show help
			err := cli.ShowAppHelp(c)
			if err != nil {
				return err
			}

		}
		return nil
	}

	app.Name = "gotask"
	app.Usage = "Add, remove and edit tasks in your local database."
	app.Author = "Beniamin Dudek"
	app.Email = "beniamin.dudek@yahoo.com"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		add,
		show,
		done,
	}

	return app
}
