package main

import (
	"context"
	"database/sql"
	"flag"
	"os"
	"os/signal"
	"syscall"

	resultStorage "psyc/internal/controllers/db/result"
	userStorage "psyc/internal/controllers/db/user"
	"psyc/internal/controllers/mail"
	"psyc/internal/service/result"
	"psyc/internal/service/user"
	"psyc/pkg/logger"
	app "psyc/pkg/transport/app"

	"psyc/internal/controllers/cache"

	_ "github.com/lib/pq"
	redis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	filepath := *flag.String("config", "config/config.yml", "Defines path to config file")

	cfgPostgres, err := resultStorage.InitConfig(filepath)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", cfgPostgres.URL)
	if err != nil {
		panic(err)
	}

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

	if _, err = redisClient.Ping(ctx).Result(); err != nil {
		panic("Redis connection error: " + err.Error())
	}

	cache := cache.New(redisClient)

	cfgMail, err := mail.InitConfig(filepath)
	if err != nil {
		panic(err)
	}

	mail := mail.New(cfgMail)

	// Storages initialization
	userStorage := userStorage.New(db)
	resultStorage := resultStorage.New(db)

	// Services initialization
	userService := user.New(userStorage, cache)
	resultService := result.New(resultStorage, mail)

	// Logger initialization
	zerolog := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger := logger.New(zerolog)

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

	log.Info().Msg("Server is running on " + appConfig.Server.Addr)

	go func(ctx context.Context) {
		if err := app.Stop(ctx); err != nil {
			panic(err)
		}
	}(ctx)

	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancel()
	log.Info().Msg("Server stopped")
}
