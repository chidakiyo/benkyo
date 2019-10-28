package log

import (
	"context"
)

type CustomLogger struct {
	c context.Context
}

func NewLogger(c context.Context) *CustomLogger {
	return &CustomLogger{
		c:c,
	}
}

// Debug
func(u *CustomLogger) Debug(format string, a ...interface{}) {
	AppLogf(u.c, SeverityDebug, format, a...)
}

// Info
func(u *CustomLogger) Info(format string, a ...interface{}) {
	AppLogf(u.c, SeverityInfo, format, a...)
}

// Warn
func(u *CustomLogger) Warn(format string, a ...interface{}) {
	AppLogf(u.c, SeverityWarning, format, a...)
}

// Error
func(u *CustomLogger) Error(format string, a ...interface{}) {
	AppLogf(u.c, SeverityError, format, a...)
}

// Fatal
func(u *CustomLogger) Fatal(format string, a ...interface{}) {
	AppLogf(u.c, SeverityCritical, format, a...)
}

// Debug
func Debug(c context.Context, format string, a ...interface{}) {
	AppLogf(c, SeverityDebug, format, a...)
}

// Info
func Info(c context.Context, format string, a ...interface{}) {
	AppLogf(c, SeverityInfo, format, a...)
}

// Warn
func Warn(c context.Context, format string, a ...interface{}) {
	AppLogf(c, SeverityWarning, format, a...)
}

// Error
func Error(c context.Context, format string, a ...interface{}) {
	AppLogf(c, SeverityError, format, a...)
}

// Fatal
func Fatal(c context.Context, format string, a ...interface{}) {
	AppLogf(c, SeverityCritical, format, a...)
}
