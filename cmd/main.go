package main

import (
	"payments-api/src/app"
)

func main() {
	api := app.NewApp()
	api.StartServer()
}
