package main

import (
	"log"
	"github.com/aayushkdev/nmsurf/internal"
)

func main() {

	app := internal.NewController()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
