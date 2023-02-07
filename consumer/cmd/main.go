package main

import (
	"context"

	"encoding/json"

	"github.com/rs/zerolog"

	env "sp-consumer/internal/env"

	models "sp-consumer/internal/gorm/models"

	logger "sp-consumer/pkg/logger"

	redis_handler "sp-consumer/internal/redis"

	gorm_handler "sp-consumer/internal/gorm"
)

var (
	cfg *env.Config

	l *zerolog.Logger

	db *gorm_handler.DBHandler

	redis redis_handler.RedisHandler
)

func processMessage(message string) {
	l.Info().Msgf("message recieved from redis: %s\n", message)

	// try to unmarshal recieved message to Order model
	var order models.Order
	err := json.Unmarshal([]byte(message), &order)

	if err != nil {
		l.Warn().Msgf("error:", err.Error())
		return
	}

	// add order to database
	err = db.OrderRepository.Add(&order)

	if err != nil {
		l.Warn().Msgf("error:", err.Error())
		return
	}

	l.Info().Msgf("message added to mysql database")
}

func main() {
	// create context based on app life
	ctx := context.Background()

	// define config structure and fill it with environment variables
	cfg = env.ParseConfig()

	// create logger with specific log level
	l = logger.NewLogger(cfg.LogLevel, true)

	// create a redis client instance
	redis = redis_handler.NewRedisHandler(cfg, ctx, l)

	// create db handler and migrate database
	db = gorm_handler.NewDbHandler(cfg, l)
	err := db.Init()

	if err != nil {
		l.Fatal().Msgf("error in initiallizing db: ", err.Error())
	}

	db.InitRepositories()
	db.Migrate()

	// start consuming and pass the callback
	redis.Consume(processMessage)
}
