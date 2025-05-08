package xlogger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
	"xyz_golang/internal/config"
)

var (
	Logger *zerolog.Logger
)

func Setup(cfg config.Config) {
	if cfg.IsDevelopment {
		l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
		l.Level(zerolog.DebugLevel)
		Logger = &l
		return
	}

	l := zerolog.New(os.Stdout).With().Timestamp().Logger()
	Logger = &l
}
