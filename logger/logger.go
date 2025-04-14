package logger

import (
	"github.com/heyubani/go-template/interfaces"
)

// NewLogger will setup a logger to be used by any side of the application.
// All log implementations to be used must implement the LoggerInterface
func NewLogger(fields map[string]interface{}) interfaces.ILogger {
	return newLogrusLogger(fields)
}

// initialise the default logger that will take over the golangs default log
func init() {
	initLogrusLogger()
}
