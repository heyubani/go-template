package logger

import (
	"log"

	"github.com/sirupsen/logrus"

	"github.com/heyubani/go-template/config"
	"github.com/heyubani/go-template/interfaces"
)

// newLogrusLogger creates a new logrus logger. This implements interfaces.ILogger interface
func newLogrusLogger(fields map[string]interface{}) interfaces.ILogger {
	logrusFields := make(logrus.Fields, len(fields))

	if fields != nil {
		for key, val := range fields {
			logrusFields[key] = val
		}
	}

	logger := logrus.New()

	logEntry := setDefaultValues(logger)
	logEntry = logEntry.WithFields(logrusFields)

	return logEntry
}

// InitLogrusLogger will initialise logrus and make it the default logger used by
// golangs 'log' package
func initLogrusLogger() {

	logger := logrus.New()

	// set default configs and add fields that will be part of the logger
	logEntry := setDefaultValues(logger)

	// Use logrus for standard log output
	log.SetOutput(logEntry.Writer())
}

func setDefaultValues(logger *logrus.Logger) *logrus.Entry {
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	})

	logger.SetReportCaller(true)
	return logger.WithFields(logrus.Fields{
		"application": config.APP_NAME,
		"environment": config.AppConfig.Env,
	})
}
