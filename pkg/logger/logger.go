package logger

import (
	"os"

	"github.com/danvergara/lazypodman/pkg/config"
	"github.com/sirupsen/logrus"
)

// NewLogger returns a new custom logger
func NewLogger(config *config.Config) *logrus.Entry {
	var log *logrus.Logger

	if config.Debug || os.Getenv("DEBUG") == "TRUE" {
		log = newDevelopmentLogger()
	} else {
		log = newProductionLogger()
	}

	return log.WithFields(logrus.Fields{
		"debug":   config.Debug,
		"version": config.Version,
	})
}

func newDevelopmentLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})
	return log
}

func newProductionLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.WarnLevel)
	log.SetOutput(os.Stdout)
	return log
}
