package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"payments-api/src/domain/transaction"
)

func CreateTransaction(rep *transaction.TransactionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t transaction.Transaction

		rb, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.Unmarshal(rb, &t)

		if t.Valid() == false {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := rep.Create(ctx, t.AccountId, t.OperationId, t.Amount); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
