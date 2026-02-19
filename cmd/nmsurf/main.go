package main

import (
	"github.com/aayushkdev/nmsurf/internal/app"
	"github.com/aayushkdev/nmsurf/internal/core"
	"github.com/aayushkdev/nmsurf/internal/providers/wifi"
)

func main() {

	controller := app.NewController(
		[]core.Provider{
			wifi.New(),
		},
	)

	controller.Run()
}
