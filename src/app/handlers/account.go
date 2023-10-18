package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"payments-api/src/domain/account"

	"github.com/go-chi/chi/v5"
)

func CreateAccount(repository *account.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a account.Account

		rb, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.Unmarshal(rb, &a)

		if a.Valid() == false {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := repository.Create(ctx, a.Document); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetAccount(rep *account.AccountRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a account.Account

		aId, err := strconv.Atoi(getURLParam(r, "accountId"))
		a.Id = aId

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		acc, err := rep.Get(ctx, a.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(acc)
	}
}

// getURLParam is used to read a value from the URL, as a string.
func getURLParam(r *http.Request, v string) string {
	return chi.URLParam(r, v)
}
