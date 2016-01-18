package golog

import (
	"io"
	"time"

	log "github.com/Sirupsen/logrus"
)

const timeFormat = "2006/01/02 15:04:05"

// stackSize is the maximum number of bytes to log for a stack trace
const stackSize = 2 * 1024

// Logger represents the object used to log errors, panics, requests, etc.
// By default, Logger writes to standard out, however this can be changed
// for testing purposes.
type Logger struct {
	app     string // the server name
	version string // the server version
}

// NewLogger creates a new Logger object with the provided server name and
// returns the Logger object pointer.
func NewLogger(app, version string) *Logger {
	return &Logger{
		app:     app,
		version: version,
	}
}

// SetOutput sets the output to be logged to.
func SetOutput(out io.Writer) {
	log.SetOutput(out)
}

func formatTime(t time.Time) string {
	return t.UTC().Format(timeFormat)
}

func (l *Logger) standardEntry(t time.Time) *log.Entry {
	return log.WithFields(log.Fields{
		"app":  l.app,
		"v":    l.version,
		"time": formatTime(time.Now()),
	})
}

// LogInfo writes the provided string to standard out with the proper logging
// format.
func (l *Logger) LogInfo(s string) {
	l.standardEntry(time.Now()).Info(s)
}

// LogError logs the provided error to standard out with the proper logging
// format.
func (l *Logger) LogError(err error) {
	l.standardEntry(time.Now()).Error(err)
}

// LogWarning logs the provided warning message to standard out with the proper
// logging format.
func (l *Logger) LogWarning(s string) {
	l.standardEntry(time.Now()).Warn(s)
}
