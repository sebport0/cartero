package main

import (
	"log"

	"github.com/sebport0/cartero/cmd/cartero/app"
)

func main() {
	app := app.NewApp()
	err := app.Run()
	if err != nil {
		log.Panicln(err)
	}
}
