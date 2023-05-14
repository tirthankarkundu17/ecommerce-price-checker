package middlewares

import (
	"errors"
	"net/http"

	"github.com/tirthankarkundu17/ecommerce-price-checker/auth"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			http.Error(w, errors.New("Unauthorized").Error(), http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
