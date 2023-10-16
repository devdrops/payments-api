package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"payments-api/src/app/handlers"
)

func NewRouter() *chi.Mux {
	// Basic router
	router := chi.NewRouter()

	// go-chi useful middlewares
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(appHeadersMiddleware)

	// Routes: healthcheck
	router.Get("/health", handlers.HealthCheck)
	// Routes: Accounts
	router.Route("/accounts", func(r chi.Router) {
		r.Post("/", handlers.CreateAccount)
		r.Get("/{accountId}", handlers.GetAccount)
	})
	// Routes: Transactions
	router.Post("/transactions", handlers.CreateTransaction)

	return router
}

// getURLParam is used to read a value from the URL, as a string.
func getURLParam(r *http.Request, v string) string {
	return chi.URLParam(r, v)
}
