package middleware

import (
	"context"
	"net/http"
	"psyc/internal/errors"
	"psyc/pkg/logger"
	"psyc/pkg/scripts"
	"strings"

	"psyc/internal/controllers/cache"
)

// Checks if token is valid
func AuthToken(log logger.Logger, cache cache.Cache) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			id, _, err := scripts.ParseJWTUserToken(r.Header.Get("Authorization"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Error(err.Error())
				return
			}
			var auth bool
			if strings.Contains(r.URL.Path, "admin") {
				auth = cache.Check(r.Context(), "admin", id)
			} else {
				auth = cache.Check(r.Context(), "user", id)
			}
			if !auth {
				http.Error(w, errors.ErrSessionNotAuthenticated, http.StatusUnauthorized)
				log.Error(errors.ErrSessionNotAuthenticated)
				return
			}
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "id", id)))
		})
	}
}
