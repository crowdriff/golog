package golog

import (
	"io"

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

// standardEntry returns an Entry with the app and version fields already set.
func (l *Logger) standardEntry() *log.Entry {
	return log.WithFields(log.Fields{
		"app": l.app,
		"v":   l.version,
	})
}

// Log writes the provided string to standard out with the proper logging
// format.
func (l *Logger) Log(msg string) {
	l.standardEntry().Print(msg)
}

// LogError logs the provided error to standard out with the proper logging
// format.
func (l *Logger) LogError(err error) {
	l.standardEntry().Error(err)
}

// LogWarning logs the provided warning message to standard out with the proper
// logging format.
func (l *Logger) LogWarning(msg string) {
	l.standardEntry().Warn(msg)
}
