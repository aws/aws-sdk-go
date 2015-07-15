// Package awslog provides a minimalistic interface for the SDK to log messages to.
package awslog

import (
	"log"
	"os"
)

// A Logger is a minimalistic interface for the SDK to log messages to. Should
// be used to provide custom logging writers for the SDK to use.
type Logger interface {
	Log(...interface{})
}

// NewDefaultLogger returns a Logger which will write log messages to stdout, and
// use same formatting runes as the stdlib log.Logger
func NewDefaultLogger() Logger {
	return &defaultLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// A defaultLogger provides a minimalistic logger satisfying the Logger interface.
type defaultLogger struct {
	logger *log.Logger
}

// Log logs the parameters to the stdlib logger. See log.Println.
func (l defaultLogger) Log(args ...interface{}) {
	l.logger.Println(args...)
}
