package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// define log levels into a map
var levels map[string]zerolog.Level = map[string]zerolog.Level{
	"debug":    zerolog.DebugLevel,
	"info":     zerolog.InfoLevel,
	"warn":     zerolog.WarnLevel,
	"error":    zerolog.ErrorLevel,
	"fatal":    zerolog.FatalLevel,
	"panic":    zerolog.PanicLevel,
	"no":       zerolog.NoLevel,
	"disabled": zerolog.Disabled,
}

// NewLogger create a logger structure based on log level
func NewLogger(levelStr string, consoleLog bool) *zerolog.Logger {
	level, ok := levels[levelStr]
	if !ok {
		level = zerolog.WarnLevel
	}

	zerolog.SetGlobalLevel(level)
	if consoleLog {
		l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger().With().Caller().Logger()
		return &l
	}

	l := zerolog.New(os.Stdout).With().Timestamp().Logger().With().Caller().Logger()
	return &l
}
