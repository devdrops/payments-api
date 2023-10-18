package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"payments-api/internal/logger"
	"payments-api/src/domain/transaction"
)

func CreateTransaction(rep *transaction.TransactionRepository, log *logger.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t transaction.Transaction

		rb, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.Unmarshal(rb, &t)

		if t.Valid() == false {
			log.Error(errors.New("Invalid input"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := rep.Create(ctx, t.AccountId, t.OperationId, t.Amount); err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
