package redis

import (
	"context"

	"time"

	"sp-consumer/internal/env"

	"github.com/go-redis/redis/v8"

	"github.com/rs/zerolog"
)

// RedisHandler expose `Consume` function
type RedisHandler interface {
	Consume(func(string))
}

// redisHandler is a implementation of RedisHandler with its variables
type redisHandler struct {
	client *redis.Client
	l      *zerolog.Logger
	ctx    context.Context
	queue  string
}

// NewRedisHandler create a new structure that implemented `ReditHandler` and pass the necessary arguments
func NewRedisHandler(cfg *env.Config, ctx context.Context, l *zerolog.Logger) RedisHandler {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisURI,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	return &redisHandler{
		client: client,
		l:      l,
		ctx:    ctx,
		queue:  cfg.RedisQueue,
	}
}

// Consume tries to act as a consumer with listening to the queue and block it until new message is recieved
// after each message we invoke callback function
func (r *redisHandler) Consume(callback func(string)) {
	r.l.Info().Msgf("Starting to consume on %s queue...", r.queue)

	for {
		// block the code for 1 second and wait for recieving new message from list of redis
		result, err := r.client.BLPop(r.ctx, 1*time.Second, r.queue).Result()

		if err != nil {
			continue
		}

		// invoke callback function and pass the message recieved
		callback(result[1])
	}
}
