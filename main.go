package main

import (
	"github.com/csazevedo/go-account-transaction/app"
	"github.com/csazevedo/go-account-transaction/config"
)

func main() {
	app := app.New(config.GetConfig())

	app.Run(":3001")
}
