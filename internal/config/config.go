package config

import (
	"os"

	"github.com/rs/zerolog"
)

type Config struct {
	DBConnString string
	LogLevel     zerolog.Level
}

func LoadConfig() *Config {
	dbConnString := os.Getenv("DB_CONN_STRING")
	logLevel := zerolog.InfoLevel

	if os.Getenv("LOG_LEVEL") == "debug" {
		logLevel = zerolog.DebugLevel
	}

	return &Config{
		DBConnString: dbConnString,
		LogLevel:     logLevel,
	}
}
