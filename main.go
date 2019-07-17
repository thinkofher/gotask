package main

import (
	"log"
	"os"
)

func main() {
	prepareApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
