package logger

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	"os"
)

func InitLogger() zerolog.Logger {
	// set up logger output and default logging level
	loggerLevel := viper.GetInt("logger.log_level")
	logger := zerolog.
		New(os.Stderr).
		With().
		Timestamp().
		Logger().
		Level(zerolog.Level(loggerLevel))

	return logger
}
