package http

import (
	"net/http"
)

// appHeadersMiddleware adds a group of headers for every HTTP response.
func appHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Server", "transactions-api")
		next.ServeHTTP(w, r)
	})
}
