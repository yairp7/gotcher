package utils

import (
	"fmt"
	"log"
	"os"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	// DEBUG level for detailed information
	DEBUG LogLevel = iota
	// INFO level for general operational information
	INFO
	// WARN level for warning messages
	WARN
	// ERROR level for error messages
	ERROR
)

// Logger interface defines the methods that any logger implementation must provide
type Logger interface {
	// Debug logs a message at DEBUG level
	Debug(format string, args ...interface{})
	// Info logs a message at INFO level
	Info(format string, args ...interface{})
	// Warn logs a message at WARN level
	Warn(format string, args ...interface{})
	// Error logs a message at ERROR level
	Error(format string, args ...interface{})
	// SetLevel sets the minimum log level to output
	SetLevel(level LogLevel)
}

// StdoutLogger implements Logger interface with output to stdout
type StdoutLogger struct {
	level  LogLevel
	logger *log.Logger
}

// NewStdoutLogger creates a new StdoutLogger instance
func NewStdoutLogger() *StdoutLogger {
	return &StdoutLogger{
		level:  INFO,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *StdoutLogger) log(level LogLevel, prefix string, format string, args ...any) {
	if level >= l.level {
		msg := fmt.Sprintf(format, args...)
		l.logger.Printf("[%s] %s", prefix, msg)
	}
}

func (l *StdoutLogger) Debug(format string, args ...any) {
	l.log(DEBUG, "DEBUG", format, args...)
}

func (l *StdoutLogger) Info(format string, args ...any) {
	l.log(INFO, "INFO", format, args...)
}

func (l *StdoutLogger) Warn(format string, args ...any) {
	l.log(WARN, "WARN", format, args...)
}

func (l *StdoutLogger) Error(format string, args ...any) {
	l.log(ERROR, "ERROR", format, args...)
}

func (l *StdoutLogger) SetLevel(level LogLevel) {
	l.level = level
}

// String returns the string representation of a LogLevel
func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", l)
	}
}
