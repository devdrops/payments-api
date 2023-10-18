package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"payments-api/internal/database"
	"payments-api/internal/logger"
)

// TODO: add these values to the body response
var (
	Version   string
	BuildDate string
)

func HealthCheck(u *database.Utils, log *logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		ok, err := u.PingDatabase(ctx)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !ok {
			log.Error(errors.New("Database connection unavailable"))
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
