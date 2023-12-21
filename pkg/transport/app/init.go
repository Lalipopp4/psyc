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

type appHTTP struct {
	user   user.Service
	result result.Service
	logger *slog.Logger
	server *http.Server
}

// Inits app and handlers
func Init(user user.Service, result result.Service, logger *slog.Logger, config *Config) App {
	rtr := mux.NewRouter()

	app := &appHTTP{
		user:   user,
		result: result,
		logger: logger,
		server: &http.Server{
			Addr:         config.Addr,
			Handler:      rtr,
			WriteTimeout: time.Duration(config.Timeout * 10e9),
		},
	}

	rtr.Use(middleware.Logging(app.logger), middleware.PanicRecovery(app.logger))

	fs := http.FileServer(http.Dir(`C:\Users\anton\Go\src\github.com\Lalipopp4\test_server\ui\html`))
	rtr.Handle("/", fs)

	rtr.HandleFunc("/reg", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/ui/html/reg.html`)
	}).Methods("GET")

	rtr.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/ui/html/auth.html`)
	}).Methods("GET")

	rtr.HandleFunc("/user/keirsey", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/ui/html/keirsey.html`)
	}).Methods("GET")

	rtr.HandleFunc("/user/hall", func(w http.ResponseWriter, r *http.Request) {
		if !sessions.Check(w, r) {
			w.WriteHeader(http.StatusForbidden)
		}
		http.ServeFile(w, r, `psyc/ui/html/hall.html`)
	}).Methods("GET")

	rtr.HandleFunc("/user/keirsey", app.keirsey).Methods("POST")

	rtr.HandleFunc("/user/hall", app.hall).Methods("POST")

	// rtr.HandleFunc("/user", app.account).Methods("GET")

	rtr.HandleFunc("/auth", app.login).Methods("POST")

	rtr.HandleFunc("/reg", app.register).Methods("POST")

	return app
}

// Starts app
func (app *appHTTP) Start() error {
	return app.server.ListenAndServe()
}

// Stops app from context
func (app *appHTTP) Stop(ctx context.Context) error {
	<-ctx.Done()
	app.logger.Info("Server stopped")
	return app.server.Shutdown(ctx)
}
