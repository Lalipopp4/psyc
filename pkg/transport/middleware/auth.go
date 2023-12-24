package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"psyc/internal/errors"
	"psyc/pkg/scripts"
	"slices"
)

var (
	admin = []string{"1", "2"}
)

// Checks if token is valid
func AuthToken(logger slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, _, err := scripts.ParseJWTUserToken(r.Header.Get("Authorization"))
			if err != nil || r.URL.Path == "admin" && !slices.Contains[[]string, string](admin, id) {
				http.Error(w, errors.ErrSessionNotAuthenticated, http.StatusUnauthorized)
				slog.Error(errors.ErrSessionNotAuthenticated)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "id", id)))
		})
	}
}
