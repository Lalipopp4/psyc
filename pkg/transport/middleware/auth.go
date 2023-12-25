package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"psyc/internal/errors"
	"psyc/pkg/scripts"

	cache "psyc/internal/controllers/cache"
)

var (
	admin = []string{"1", "2"}
)

// Checks if token is valid
func AuthToken(logger slog.Logger, cache cache.Cache) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, _, err := scripts.ParseJWTUserToken(r.Header.Get("Authorization"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				slog.Error(err.Error())
				return
			}
			var auth bool
			switch r.URL.Path {
			case "/admin":
				auth = cache.Check(r.Context(), "admin", id)
			default:
				auth = cache.Check(r.Context(), "user", id)
			}

			if !auth {
				http.Error(w, errors.ErrSessionNotAuthenticated, http.StatusUnauthorized)
				slog.Error(errors.ErrSessionNotAuthenticated)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "id", id)))
		})
	}
}
