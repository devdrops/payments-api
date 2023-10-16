package app

import (
	"net/http"

	"payments-api/internal/config"
	"payments-api/src/http"
)

type App struct {
	c *config.Config
}

func NewApp() *App {
	cfg := config.NewConfig()

	return &App{
		c: cfg,
	}
}

func (a *App) StartServer() {
	router := http.NewRouter()
	http.ListenAndServe(a.c.Port, router)
}
