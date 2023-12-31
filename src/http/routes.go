package http

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"payments-api/internal/database"
	"payments-api/internal/logger"
	"payments-api/src/app/handlers"
	"payments-api/src/domain/account"
	"payments-api/src/domain/transaction"
)

func NewRouter(accRep *account.AccountRepository, trxRep *transaction.TransactionRepository, utils *database.Utils, log *logger.Logger) *chi.Mux {
	// Basic router
	router := chi.NewRouter()

	// go-chi useful middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentType("application/json"))
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(5 * time.Second))
	router.Use(appHeadersMiddleware)

	// Routes: healthcheck
	router.Get("/health", handlers.HealthCheck(utils, log))
	// Routes: Accounts
	router.Route("/accounts", func(r chi.Router) {
		r.Post("/", handlers.CreateAccount(accRep, log))
		r.Get("/{accountId}", handlers.GetAccount(accRep, log))
	})
	// Routes: Transactions
	router.Post("/transactions", handlers.CreateTransaction(trxRep, log))

	// 404
	router.NotFound(handlers.NotFound)
	// 405
	router.MethodNotAllowed(handlers.MethodNotAllowed)

	return router
}
