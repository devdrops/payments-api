package handlers

import (
	"context"
	"net/http"
	"time"

	"payments-api/internal/database"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	ctx, cancel := context.WithTimeout(ctx, 1 * time.Second)
	defer cancel()

	adapter := database.NewPostgresAdapter()
	if err, _ := adapter.Ping(ctx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

        w.WriteHeader(http.StatusOK)
}
