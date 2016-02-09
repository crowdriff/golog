package golog

import (
	"io"
	"runtime"
	"strings"

	log "github.com/Sirupsen/logrus"
)

// rootLogger represents the global logger object
var rootLogger *logger

// Logger represents the object used to log errors, panics, requests, etc.
// By default, Logger writes to standard out, however this can be changed
// for testing purposes.
type logger struct {
	app     string // the server name
	version string // the server version
}

// newLogger creates a new Logger object with the provided server name and
// returns the Logger object pointer.
func newLogger(app, version string) *logger {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
	})
	return &logger{
		app:     app,
		version: version,
	}
}

// Init sets the application name and version. This is required before using
// any logging functionality.
func Init(app, version string) {
	if rootLogger == nil {
		rootLogger = newLogger(app, version)
	}
}

// SetOutput sets the output to be logged to.
func SetOutput(out io.Writer) {
	log.SetOutput(out)
}

// standardEntry returns an Entry with the app and version fields already set.
func (l *logger) standardEntry() *log.Entry {
	return log.WithFields(log.Fields{
		"app": l.app,
		"v":   l.version,
	})
}

// Log writes the provided string to standard out with the proper logging
// format.
func Log(msg string) {
	rootLogger.standardEntry().Info(msg)
}

// addFileLine is a helper that adds the file & line of the original
// caller of the Logging function if possible.
func addFileLine(entry *log.Entry) *log.Entry {
	if _, file, line, ok := runtime.Caller(2); ok {
		// cut all of the filepath before the "src" folder
		if idx := strings.Index(file, "/src/"); idx > -1 {
			file = file[idx+5:]
		}
		entry = entry.WithFields(log.Fields{
			"file": file,
			"line": line,
		})
	}
	return entry
}

// LogError logs the provided error to standard out with the proper logging
// format.
func LogError(err error) {
	entry := rootLogger.standardEntry()
	addFileLine(entry).Error(err)
}

// LogFatal logs the provided error to standard out with the proper logging
// format, and then exits.
func LogFatal(v ...interface{}) {
	entry := rootLogger.standardEntry()
	addFileLine(entry).Fatal(v...)
}

// LogPanic logs the provided error to standard out with the proper logging
// format, and then panics.
func LogPanic(v interface{}) {
	entry := rootLogger.standardEntry()
	addFileLine(entry).Panic(v)
}

// LogWarning logs the provided warning message to standard out with the proper
// logging format.
func LogWarning(msg string) {
	rootLogger.standardEntry().Warn(msg)
}
