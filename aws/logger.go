package aws

import (
	"fmt"
	"log"
	"os"
)

// A LogLevelType defines the level logging should be performed at. Used to instruct
// the SDK which statements should be logged.
type LogLevelType uint

// LogLevel returns the pointer to a LogLevel. Should be used to workaround
// not being able to take the address of a non-composite literal.
func LogLevel(l LogLevelType) *LogLevelType {
	return &l
}

// Value returns the LogLevel value or the default value LogOff if the LogLevel
// is nil. Safe to use on nil value LogLevelTypes.
func (l *LogLevelType) Value() LogLevelType {
	if l != nil {
		return *l
	}
	return LogOff
}

// Matches returns true if the v LogLevel is enabled by this LogLevel. Should be
// used with logging sub levels. Is safe to use on nil value LogLevelTypes. If
// LogLevel is nil, will default to LogOff comparison.
func (l *LogLevelType) Matches(v LogLevelType) bool {
	c := l.Value()
	return c&v == v
}

// AtLeast returns true if this LogLevel is at least high enough to satisfies v.
// Is safe to use on nil value LogLevelTypes. If LogLevel is nil, will default
// to LogOff comparison.
func (l *LogLevelType) AtLeast(v LogLevelType) bool {
	c := l.Value()
	return c >= v
}

const (
	// LogOff states that no logging should be performed by the SDK. This is the
	// default state of the SDK, and should be use to disable all logging.
	LogOff LogLevelType = iota * 0x1000

	// LogDebug state that debug output should be logged by the SDK. This should
	// be used to inspect request made and responses received.
	LogDebug
)

// Debug Logging Sub Levels
const (
	// LogDebugWithSigning states that the SDK should log request signing and
	// presigning events. This should be used to log the signing details of
	// requests for debugging. Will also enable LogDebug.
	LogDebugWithSigning LogLevelType = LogDebug | (1 << iota)

	// LogDebugWithHTTPBody states the SDK should log HTTP request and response
	// HTTP bodys in addition to the headers and path. This should be used to
	// see the body content of requests and responses made while using the SDK
	// Will also enable LogDebug.
	LogDebugWithHTTPBody

	// LogDebugWithRequestRetries states the SDK should log when service requests will
	// be retried. This should be used to log when you want to log when service
	// requests are being retried. Will also enable LogDebug.
	LogDebugWithRequestRetries

	// LogDebugWithRequestErrors states the SDK should log when service requests fail
	// to build, send, validate, or unmarshal.
	LogDebugWithRequestErrors

	// LogDebugWithEventStreamBody states the SDK should log EventStream
	// request and response bodys. This should be used to log the EventStream
	// wire unmarshaled message content of requests and responses made while
	// using the SDK Will also enable LogDebug.
	LogDebugWithEventStreamBody
)

// A Logger is a minimalistic interface for the SDK to log messages to. Should
// be used to provide custom logging writers for the SDK to use.
// Deprecated: Use ContextLogger instead.
type Logger interface {
	Log(...interface{})
}

// A ContextLogger is a minimalistic interface for the SDK to log messages to.
// Should be used to provide custom logging writers for the SDK to use.
type ContextLogger interface {
	Debug(Context, ...interface{})
	Debugf(Context, string, ...interface{})
	Info(Context, ...interface{})
	Infof(Context, string, ...interface{})
	Warn(Context, ...interface{})
	Warnf(Context, string, ...interface{})
	Error(Context, ...interface{})
	Errorf(Context, string, ...interface{})
}

// A LoggerFunc is a convenience type to convert a function taking a variadic
// list of arguments and wrap it so the Logger interface can be used.
// Deprecated: Use SimpleContextLoggerFunc instead.
//
// Example:
//     s3.New(sess, &aws.Config{Logger: aws.LoggerFunc(func(args ...interface{}) {
//         fmt.Fprintln(os.Stdout, args...)
//     })})
type LoggerFunc func(...interface{})

// Log calls the wrapped function with the arguments provided
func (f LoggerFunc) Log(args ...interface{}) {
	f(args...)
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

// A SimpleContextLoggerFunc is a convenience type to convert a function taking a
// logging level and message and wrap it so the ContextLogger interface can be used.
//
// Example:
//     f := func(ctx context.Context, level string, msg string) {
//         fmt.Fprint(os.Stderr, "[" + level + "] " + msg)
//     }
//     s3.New(sess, &aws.Config{
//         ContextLogger: aws.SimpleContextLoggerFunc(f),
//     })
type SimpleContextLoggerFunc func(ctx Context, level string, msg string)

// Debug implements the ContextLogger Debug method by calling the wrapped function.
func (f SimpleContextLoggerFunc) Debug(ctx Context, v ...interface{}) {
	f(ctx, "DEBUG", fmt.Sprintln(v...))
}

// Debugf implements the ContextLogger Debugf method by calling the wrapped function.
func (f SimpleContextLoggerFunc) Debugf(ctx Context, format string, v ...interface{}) {
	f(ctx, "DEBUG", fmt.Sprintf(format+"\n", v...))
}

// Info implements the ContextLogger Info method by calling the wrapped function.
func (f SimpleContextLoggerFunc) Info(ctx Context, v ...interface{}) {
	f(ctx, "INFO", fmt.Sprintln(v...))
}

// Infof implements the ContextLogger Infof method by calling the wrapped function.
func (f SimpleContextLoggerFunc) Infof(ctx Context, format string, v ...interface{}) {
	f(ctx, "INFO", fmt.Sprintf(format+"\n", v...))
}

// Warn implements the ContextLogger Warn method by calling the wrapped function.
func (f SimpleContextLoggerFunc) Warn(ctx Context, v ...interface{}) {
	f(ctx, "WARNING", fmt.Sprintln(v...))
}

// Warnf implements the ContextLogger Warnf method by calling the wrapped function.
func (f SimpleContextLoggerFunc) Warnf(ctx Context, format string, v ...interface{}) {
	f(ctx, "WARNING", fmt.Sprintf(format+"\n", v...))
}

// Error implements the ContextLogger Error method by calling the wrapped function.
func (f SimpleContextLoggerFunc) Error(ctx Context, v ...interface{}) {
	f(ctx, "ERROR", fmt.Sprintln(v...))
}

// Errorf implements the ContextLogger Errorf method by calling the wrapped function.
func (f SimpleContextLoggerFunc) Errorf(ctx Context, format string, v ...interface{}) {
	f(ctx, "ERROR", fmt.Sprintf(format+"\n", v...))
}
