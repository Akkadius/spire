package boot

import (
	"github.com/Akkadius/spire/internal/env"
	"github.com/go-errors/errors"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

const (
	LoggingFormatText = "text"
	LoggingFormatJson = "json"
)

// wire set for loading the stores.
var loggerSet = wire.NewSet(
	provideLogger,
)

// logging provider
func provideLogger() (*logrus.Logger, error) {
	baseLogger := logrus.New()
	cnf, err := getLoggerConfig()
	if err != nil {
		return nil, err
	}

	baseLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableLevelTruncation: true,
	})

	if cnf.Formatter == LoggingFormatJson {
		baseLogger.SetFormatter(&logrus.JSONFormatter{})
	}

	baseLogger.WithFields(logrus.Fields{
		"application": cnf.Application,
		"environment": cnf.Environment,
	})

	return baseLogger, nil
}

type LoggerConfig struct {
	Application string
	Environment string
	Formatter   string // json or text
}

// return logger config
func getLoggerConfig() (*LoggerConfig, error) {
	c := &LoggerConfig{
		Application: env.Get("APP_NAME", "local"),
		Environment: env.Get("APP_ENV", "local"),
		Formatter:   env.Get("LOGGING_FORMAT", "text"),
	}

	if len(c.Application) == 0 {
		return &LoggerConfig{}, errors.New("application must be present")
	}

	return c, nil
}
