package middleware

import (
	"net/http"
)

func JsonResponse(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		handler.ServeHTTP(w, r)
	}
}
