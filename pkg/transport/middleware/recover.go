package middleware

import (
	"net/http"
	"psyc/pkg/logger"
)

// Catches panics and recovers app
func PanicRecovery(log logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Error("panic recovered on %s: %v", r.URL.Path, err)
					http.Error(w, err.(error).Error(), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
