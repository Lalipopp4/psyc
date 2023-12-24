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
func Init(user user.Service, result result.Service, logger slog.Logger, config *Config) App {
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

	rtr.HandleFunc("/user/keirsey", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/static/html/keirsey.html`)
	}).Methods("GET")

	rtr.HandleFunc("/user/hall", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/static/html/hall.html`)
	}).Methods("GET")

	rtr.HandleFunc("/user/keirsey", resultHandler.keirsey).Methods("POST")

	rtr.HandleFunc("/user/hall", resultHandler.hall).Methods("POST")

	rtr.HandleFunc("/user/bass", resultHandler.keirsey).Methods("POST")

	rtr.HandleFunc("/user/eysenck", resultHandler.hall).Methods("POST")

	rtr.HandleFunc("/user", resultHandler.account)

	rtr.HandleFunc("/auth", userHandler.login).Methods("POST")

	rtr.HandleFunc("/reg", userHandler.register).Methods("POST")

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
