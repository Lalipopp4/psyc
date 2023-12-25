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
	stop     = "q"
	filepath = "config/config.yaml"
)

func main() {

	cfgPostgres, err := resultStorage.InitConfig(filepath)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", cfgPostgres.URL)
	ctx, cancel := context.WithCancel(context.Background())

	cfgRedis, err := cache.InitConfig(filepath)
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfgRedis.Redis.Addr,
		Password: cfgRedis.Redis.Password,
		DB:       cfgRedis.Redis.Database,
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
	appConfig, err := app.InitConfig(filepath)

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
