package transport

import (
	"context"
	"log/slog"
	"net/http"
	"psyc/internal/service/result"
	"psyc/internal/service/user"
	"psyc/internal/sessions"
	"time"

	"psyc/pkg/transport/middleware"

	cache "psyc/internal/controllers/cache"

	"github.com/gorilla/mux"
)

type userHandler struct {
	service user.Service
	logger  slog.Logger
}

type resultHandler struct {
	service result.Service
	logger  slog.Logger
}

type appHTTP struct {
	user   *userHandler
	result *resultHandler
	server *http.Server
}

// Inits app and handlers
func Init(user user.Service, result result.Service, logger slog.Logger, config *Config, cache cache.Cache) App {
	rtr := mux.NewRouter()

	userHandler := &userHandler{
		service: user,
		logger:  logger,
	}

	resultHandler := &resultHandler{
		service: result,
		logger:  logger,
	}

	app := &appHTTP{
		user:   userHandler,
		result: resultHandler,
		server: &http.Server{
			Addr:         config.Server.Addr,
			Handler:      rtr,
			WriteTimeout: time.Duration(config.Server.Timeout * 10e9),
		},
	}

	rtr.Use(middleware.Logging(logger), middleware.PanicRecovery(logger))

	fs := http.FileServer(http.Dir(`psyc/static/html`))
	rtr.Handle("/", fs)

	rtr.HandleFunc("/reg", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/static/html/reg.html`)
	}).Methods("GET")

	rtr.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/static/html/auth.html`)
	}).Methods("GET")

	rtr.HandleFunc("/auth", userHandler.login).Methods("POST")

	rtr.HandleFunc("/reg", userHandler.register).Methods("POST")

	auth := rtr.PathPrefix("/user").Subrouter()

	auth.Use(middleware.AuthToken(logger, cache))

	auth.HandleFunc("/keirsey", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/static/html/keirsey.html`)
	}).Methods("GET")

	auth.HandleFunc("/bass", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/static/html/bass.html`)
	}).Methods("GET")

	auth.HandleFunc("/eysenck", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/static/html/eysenck.html`)
	}).Methods("GET")

	auth.HandleFunc("/hall", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/static/html/hall.html`)
	}).Methods("GET")

	auth.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/static/html/info.html`)
	}).Methods("GET")

	auth.HandleFunc("/info", userHandler.info).Methods("POST")

	auth.HandleFunc("/keirsey", resultHandler.keirsey).Methods("POST")

	auth.HandleFunc("/hall", resultHandler.hall).Methods("POST")

	auth.HandleFunc("/bass", resultHandler.keirsey).Methods("POST")

	auth.HandleFunc("/eysenck", resultHandler.hall).Methods("POST")

	auth.HandleFunc("", resultHandler.account)

	auth.HandleFunc("/admin", resultHandler.admin).Methods("POST")

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
