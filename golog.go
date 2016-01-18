package golog

import (
	"io"
	"time"

	log "github.com/Sirupsen/logrus"
)

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
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
	})
	return &Logger{
		app:     app,
		version: version,
	}
}

// SetOutput sets the output to be logged to.
func (l *Logger) SetOutput(out io.Writer) {
	log.SetOutput(out)
}

func (l *Logger) standardEntry(t time.Time) *log.Entry {
	return log.WithFields(log.Fields{
		"app": l.app,
		"v":   l.version,
	})
}

// Log writes the provided string to standard out with the proper logging
// format.
func (l *Logger) Log(s string) {
	l.standardEntry(time.Now()).Print(s)
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
