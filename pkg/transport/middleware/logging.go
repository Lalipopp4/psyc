package middleware

import (
	"net/http"
	"psyc/pkg/logger"
	"time"
)

// Logs requests
func Logging(logger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Info("\n-----New request-----\n, Method: %s\nPath: %s\nBenchmark time: %s\n---------------------\n",
				r.Method, r.URL.Path, time.Since(start).String())
		})
	}
}
