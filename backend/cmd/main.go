package main

import (
	"context"

	"github.com/rs/zerolog"

	"sp-interview/internal/env"

	"sp-interview/internal/fiber/routers"

	"sp-interview/pkg/logger"

	redis_handler "sp-interview/internal/redis"
)

var (
	cfg *env.Config

	l *zerolog.Logger
)

func main() {
	// create context based on app life
	ctx := context.Background()

	// define config structure and fill it with environment variables
	cfg = env.ParseConfig()

	// create logger with specific log level
	l = logger.NewLogger(cfg.LogLevel, true)

	// create a redis client instance
	redis := redis_handler.NewRedisHandler(cfg, ctx, l)

	// initiallize the fiber
	r := routers.NewRouter(redis, l, cfg)

	r.Init()
}
