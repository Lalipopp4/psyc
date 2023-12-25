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

	cache "psyc/internal/controllers/cache"

	_ "github.com/lib/pq"
	redis "github.com/redis/go-redis/v9"
)

const (
	stop = "q"
)

func main() {

	// cfgdb := &resultStorage.InitConfig{"psyc/config/config.yaml"}

	connStr := "user=postgres password=postgres dbname=psyc sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	ctx, cancel := context.WithCancel(context.Background())

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		panic("Redis connection error: " + err.Error())
	}

	cache := cache.New(redisClient)

	// Storages initialization
	userStorage := userStorage.New(db)
	resultStorage := resultStorage.New(db)

	// Services initialization
	userService := user.New(userStorage, cache)
	resultService := result.New(resultStorage)

	// Logger initialization
	logger := slog.Logger{}

	// Config initialization
	appConfig, err := app.InitConfig("config/config.yaml")

	if err != nil {
		panic(err)
	}

	// App initialization
	app := app.Init(userService, resultService, logger, appConfig, cache)

	// Starting app
	go func() {
		if err := app.Start(); err != nil {
			panic(err)
		}
	}()

	slog.Info("Server is running on " + appConfig.Server.Addr)
	slog.Info("To stop it enter " + stop)

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
