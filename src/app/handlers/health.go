package handlers

import (
	"context"
	"net/http"
	"time"

	"payments-api/internal/database"
)

var (
	Version   string
	BuildDate string
)

func HealthCheck(u *database.Utils) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		if ok, err := u.PingDatabase(ctx); err != nil || ok == false {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
