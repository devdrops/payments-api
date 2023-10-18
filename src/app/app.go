package app

import (
	"net/http"

	"payments-api/internal/config"
	"payments-api/internal/database"
	"payments-api/internal/database/postgres"
	"payments-api/internal/logger"
	"payments-api/src/domain/account"
	"payments-api/src/domain/transaction"
	apphttp "payments-api/src/http"
)

type App struct {
	Cfg *config.Config
	Acc *account.AccountRepository
	Trx *transaction.TransactionRepository
	Utl *database.Utils
	Log *logger.Logger
}

func NewApp() *App {
	cfg := config.NewConfig()
	db, _ := postgres.NewAdapter(cfg)

	return &App{
		Cfg: cfg,
		Acc: account.NewRepository(db),
		Trx: transaction.NewRepository(db),
		Utl: database.NewUtils(db),
		Log: logger.NewLogger(),
	}
}

func (a *App) StartServer() {
	router := apphttp.NewRouter(a.Acc, a.Trx, a.Utl, a.Log)
	http.ListenAndServe(a.Cfg.Port, router)
}
