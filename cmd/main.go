package main

import (
	"payments-api/src/app"
)

func main() {
	a := app.New()
	a.StartServer()
}
