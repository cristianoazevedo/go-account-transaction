package main

import (
	"github.com/csazevedo/go-account-transaction/app"
	"github.com/csazevedo/go-account-transaction/config"
)

func main() {
	config := config.GetConfig()
	app := app.New(config)

	app.Run(config.App.Host)
}
