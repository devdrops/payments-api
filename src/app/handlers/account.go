package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"payments-api/src/domain/account"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
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

        if err := s.Create(a); err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

        w.WriteHeader(http.StatusCreated)
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
        var a account.Account

        aId, err := strconv.Atoi(getURLParam(r, "accountId"))
        a.Id = aId

        acc, err := s.Get(a)
        if err != nil {
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

        json.NewEncoder(w).Encode(acc)
}
