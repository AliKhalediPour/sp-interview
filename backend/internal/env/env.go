package env

import (
	"sp-interview/pkg/env"
)

// declare Config to hold the environment variables with default values
type Config struct {
	Mode     string `env:"MODE" envDefault:"prod"`
	Port     int    `env:"PORT" envDefault:"5001"`
	LogLevel string `env:"LOGLEVEL" envDefault:"debug"`

	RedisURI      string `env:"REDIS_URI" envDefault:"redis:6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisQueue    string `env:"REDIS_QUEUE" envDefault:"order"`
}

func ParseConfig() *Config {
	return env.ParseConfig[Config]()
}
