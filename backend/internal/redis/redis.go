package redis

import (
	"context"

	"sp-interview/internal/env"

	"github.com/go-redis/redis/v8"

	"github.com/rs/zerolog"
)

// RedisHandler expose Push function
type RedisHandler interface {
	Push(message string, context context.Context) error
}

// redisHandler is a wrapper that implement the RedisHandler interface functions
type redisHandler struct {
	client *redis.Client
	logger *zerolog.Logger
	ctx    context.Context
	queue  string
}

// NewRedisHandler create a redis client structure and wrap it into redisHandler
func NewRedisHandler(cfg *env.Config, ctx context.Context, l *zerolog.Logger) RedisHandler {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURI,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	return &redisHandler{
		client: client,
		logger: l,
		ctx:    ctx,
		queue:  cfg.RedisQueue,
	}
}

// Add push the message into queue
func (r *redisHandler) Push(message string, context context.Context) error {
	result, err := r.client.RPush(context, r.queue, message).Result()

	if err != nil {
		r.logger.Warn().Msgf("error in pushing message to queue: %s", err.Error())
		return err
	}

	r.logger.Info().Msgf("message pushed to queue, length: %d", result)

	return nil
}
