package env

import (
	"sp-consumer/pkg/env"
)

// declare Config to hold the environment variables with default values
type Config struct {
	Mode     string `env:"MODE" envDefault:"prod"`
	LogLevel string `env:"LOGLEVEL" envDefault:"info"`

	RedisURI      string `env:"REDIS_URI" envDefault:"redis:6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisQueue    string `env:"REDIS_QUEUE" envDefault:"order"`

	MySqlUsername string `env:"MY_SQL_USERNAME" envDefault:"snappfood"`
	MySqlPassword string `env:"MY_SQL_PASSWORD" envDefault:"PassWorD"`
	MySqlHost     string `env:"MY_SQL_HOST" envDefault:"mysql"`
	MySqlPort     int    `env:"MY_SQL_PORT" envDefault:"3306"`
	MySqlDbName   string `env:"MY_SQL_DBNAME" envDefault:"snappfood"`
}

func ParseConfig() *Config {
	return env.ParseConfig[Config]()
}
