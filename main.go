package main

import (
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/thinkofher/gotask/db"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := filepath.Join(home, ".tasks.db")

	err = db.InitDB(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	prepareApp()

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
