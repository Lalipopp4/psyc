package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"psyc/internal/errors"
	"psyc/pkg/scripts"
)

// Checks if token is valid
func AuthToken(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if user, _, err := scripts.ParseJWTUserToken(r.Header.Get("Authorization")); err == nil {
				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "id", user)))
				return
			}
			http.Error(w, errors.ErrSessionNotAuthenticated, http.StatusUnauthorized)
			slog.Error(errors.ErrSessionNotAuthenticated)
		})
	}
}
