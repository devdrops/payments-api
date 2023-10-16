package handlers

import (
	"net/http"
	"encoding/json"
	"io"

	"payments-api/src/domain/transaction"
)

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
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

        if err := s.Create(t); err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

        w.WriteHeader(http.StatusCreated)
}
