package middleware

import (
	"log/slog"
	"net/http"
)

// Catches panics and recovers app
func PanicRecovery(logger slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					slog.Error("panic recovered on %s: %v", r.URL.Path, err)
					http.Error(w, err.(error).Error(), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
