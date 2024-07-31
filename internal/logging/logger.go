package logging

import (
	"github.com/rs/zerolog"
	"log"
	"os"
	"time"
)

func InitLogger(level zerolog.Level) {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339Nano,
	})
	zerolog.SetGlobalLevel(level)
}

func SetLogLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}
