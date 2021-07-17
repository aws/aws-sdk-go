package awslog

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
)

// Error logs an error via the Config's ContextLogger or Logger.
func Error(ctx aws.Context, c *aws.Config, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Error(ctx, v...)
	} else if c.Logger != nil {
		c.Logger.Log(prefixLevel("ERROR:", v)...)
	} else {
		// no-op
	}
}

// Errorf logs an error via the Config's ContextLogger or Logger.
func Errorf(ctx aws.Context, c *aws.Config, format string, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Errorf(ctx, format, v...)
	} else if c.Logger != nil {
		c.Logger.Log(fmt.Sprintf("ERROR: "+format, v...))
	} else {
		// no-op
	}
}

// Warn logs a warning via the Config's ContextLogger or Logger.
func Warn(ctx aws.Context, c *aws.Config, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Warn(ctx, v...)
	} else if c.Logger != nil {
		c.Logger.Log(prefixLevel("WARNING:", v)...)
	} else {
		// no-op
	}
}

// Warnf logs a warning via the Config's ContextLogger or Logger.
func Warnf(ctx aws.Context, c *aws.Config, format string, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Warnf(ctx, format, v...)
	} else if c.Logger != nil {
		c.Logger.Log(fmt.Sprintf("WARNING: "+format, v...))
	} else {
		// no-op
	}
}

// Info logs an information message via the Config's ContextLogger or Logger.
func Info(ctx aws.Context, c *aws.Config, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Info(ctx, v...)
	} else if c.Logger != nil {
		c.Logger.Log(prefixLevel("INFO:", v)...)
	} else {
		// no-op
	}
}

// Infof logs an information message via the Config's ContextLogger or Logger.
func Infof(ctx aws.Context, c *aws.Config, format string, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Infof(ctx, format, v...)
	} else if c.Logger != nil {
		c.Logger.Log(fmt.Sprintf("INFO: "+format, v...))
	} else {
		// no-op
	}
}

// DebugError logs a debug error via the Config's ContextLogger or Logger.
func DebugError(ctx aws.Context, c *aws.Config, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Error(ctx, v...)
	} else if c.Logger != nil {
		c.Logger.Log(prefixLevel("DEBUG ERROR:", v)...)
	} else {
		// no-op
	}
}

// DebugErrorf logs a debug error via the Config's ContextLogger or Logger.
func DebugErrorf(ctx aws.Context, c *aws.Config, format string, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Errorf(ctx, format, v...)
	} else if c.Logger != nil {
		c.Logger.Log(fmt.Sprintf("DEBUG ERROR: "+format, v...))
	} else {
		// no-op
	}
}

// Debug logs a debug message via the Config's ContextLogger or Logger.
func Debug(ctx aws.Context, c *aws.Config, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Debug(ctx, v...)
	} else if c.Logger != nil {
		c.Logger.Log(prefixLevel("DEBUG:", v)...)
	} else {
		// no-op
	}
}

// Debugf logs a debug message via the Config's ContextLogger or Logger.
func Debugf(ctx aws.Context, c *aws.Config, format string, v ...interface{}) {
	if c.ContextLogger != nil {
		c.ContextLogger.Debugf(ctx, format, v...)
	} else if c.Logger != nil {
		c.Logger.Log(fmt.Sprintf("DEBUG: "+format, v...))
	} else {
		// no-op
	}
}

func prefixLevel(level string, v []interface{}) []interface{} {
	cpy := make([]interface{}, len(v)+1)
	cpy[0] = level
	copy(cpy[1:], v)
	return cpy
}
