package transport

import (
	"context"
	"log/slog"
	"net/http"
	"psyc/internal/service/result"
	"psyc/internal/service/user"

	"psyc/pkg/logger"
	"psyc/pkg/transport/middleware"

	cache "psyc/internal/controllers/cache"

	"github.com/gorilla/mux"
)

type userHandler struct {
	service user.Service
	logger  logger.Logger
	cfg     *Config
}

type resultHandler struct {
	service result.Service
	logger  logger.Logger
	cfg     *Config
}

type appHTTP struct {
	user   *userHandler
	result *resultHandler
	server *http.Server
}

const addr = "localhost:8080"

// Inits app and handlers
func Init(user user.Service, result result.Service, logger logger.Logger, config *Config, cache cache.Cache) App {
	rtr := mux.NewRouter()

	userHandler := &userHandler{
		service: user,
		logger:  logger,
		cfg:     config,
	}

	resultHandler := &resultHandler{
		service: result,
		logger:  logger,
		cfg:     config,
	}

	app := &appHTTP{
		user:   userHandler,
		result: resultHandler,
		server: &http.Server{
			Addr:         addr,
			Handler:      rtr,
			WriteTimeout: config.Timeout * 10e9,
		},
	}

	rtr.Use(middleware.Logging(logger), middleware.PanicRecovery(logger))

	rtr.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/html/index.html")
	})

	rtr.HandleFunc("/reg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, `./static/html/reg.html`)
	}).Methods("GET")

	rtr.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, `./static/html/auth.html`)
	}).Methods("GET")

	rtr.HandleFunc("/auth", userHandler.auth).Methods("POST")

	rtr.HandleFunc("/reg", userHandler.register).Methods("POST")

	auth := rtr.PathPrefix("/user").Subrouter()

	auth.Use(middleware.AuthToken(logger, cache))

	rtr.HandleFunc("/test", resultHandler.getTest).Methods("GET")
	rtr.HandleFunc("", resultHandler.results).Methods("GET")
	rtr.HandleFunc("/result", resultHandler.addResult).Methods("POST")
	rtr.HandleFunc("/test", resultHandler.addTest).Methods("POST")
	rtr.HandleFunc("/review", resultHandler.addReview).Methods("POST")
	rtr.HandleFunc("/review", resultHandler.getReview).Methods("GET")

	auth.HandleFunc("/info", userHandler.info).Methods("POST")

	return app
}

// Starts app
func (app *appHTTP) Start() error {
	return app.server.ListenAndServe()
}

// Stops app from context
func (app *appHTTP) Stop(ctx context.Context) error {
	<-ctx.Done()
	slog.Info("Server stopped")
	return app.server.Shutdown(ctx)
}
