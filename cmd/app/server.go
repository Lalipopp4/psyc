package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	resultStorage "psyc/internal/controllers/db/result"
	userStorage "psyc/internal/controllers/db/user"
	"psyc/internal/service/result"
	"psyc/internal/service/user"
	app "psyc/pkg/transport/app"

	_ "github.com/lib/pq"
)

const (
	stop = "q"
)

func main() {

	connStr := "user=postgres password=postgres dbname=psyc sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	// Storages initialization
	userStorage := userStorage.New(db)
	resultStorage := resultStorage.New(db)

	// Services initialization
	userService := user.New(userStorage)
	resultService := result.New(resultStorage)

	// Logger initialization
	logger := &slog.Logger{}

	// Config initialization
	appConfig, err := app.InitConfig("config/config.yaml")

	if err != nil {
		panic(err)
	}

	// App initialization
	app := app.Init(userService, resultService, logger, appConfig)

	// Starting app
	go func() {
		if err := app.Start(); err != nil {
			panic(err)
		}
	}()

	logger.Info("Server is running on %s...", appConfig.Addr)
	logger.Info("To stop it enter %s", stop)

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		if err := app.Stop(ctx); err != nil {
			panic(err)
		}
	}(ctx)

	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
}
