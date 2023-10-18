package handlers

import (
	"context"
	"net/http"
	"time"

	"payments-api/internal/database"
	"payments-api/internal/logger"
)

var (
	Version   string
	BuildDate string
)

func HealthCheck(u *database.Utils, log *logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		_, err := u.PingDatabase(ctx)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
