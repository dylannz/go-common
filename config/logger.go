package config

import (
	"errors"
	"os"

	"github.com/HomesNZ/go-common/env"
	"github.com/Sirupsen/logrus"
	bugsnag "github.com/bugsnag/bugsnag-go"
)

// InitLogger initializes the logger by setting the log level to the env var LOG_LEVEL, or defaulting to `info`.
func InitLogger() {
	// If running in the production environment, output the logs as JSON format for parsing by Logstash.
	if env.IsProd() {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}

	logrus.SetOutput(os.Stdout)

	level, err := logrus.ParseLevel(env.GetString("LOG_LEVEL", "info"))
	// No need to handle the error here, just don't update the log level
	if err == nil {
		logrus.SetLevel(level)
	}

	// Hooks
	logrus.AddHook(bugsnagHook{})

	logrus.Infof("Log level: %s", logrus.GetLevel().String())
}

type bugsnagHook struct{}

// Levels returns the logging levels that this hook will be fired for.
func (b bugsnagHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.PanicLevel,
		logrus.WarnLevel,
	}
}

// Fire sends the logrus entry to bugsnag.
func (b bugsnagHook) Fire(entry *logrus.Entry) error {
	return bugsnag.Notify(errors.New(entry.Message), entry.Data)
}
