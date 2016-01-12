package golog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// stackSize is the maximum number of bytes to log for a stack trace
const stackSize = 2 * 1024

// Logger represents the object used to log errors, panics, requests, etc.
// By default, Logger writes to standard out, however this can be changed
// for testing purposes.
type Logger struct {
	mu      sync.Mutex // mutex that protects everything below
	Out     io.Writer  // where the logs go
	svrName string     // the server name prefix to use
	buf     []byte     // buffer that accumulates the log message
}

// NewLogger creates a new Logger object with the provided server name and
// returns the Logger object pointer.
func NewLogger(svr string) *Logger {
	return &Logger{
		Out:     os.Stdout,
		svrName: svr,
	}
}

// formatHeader writes the provided time to the logging buffer slice as well as
// the server name prefix with the following format:
// '<year>/<month>/<day> <hour>:<min>:<sec> [<server name>] '
func (l *Logger) formatHeader(t time.Time) {
	l.buf = l.buf[:0]

	year, month, day := t.Date()
	l.buf = append(l.buf, strconv.Itoa(year)...)
	l.buf = append(l.buf, '/')
	l.buf = append(l.buf, strconv.Itoa(int(month))...)
	l.buf = append(l.buf, '/')
	l.buf = append(l.buf, strconv.Itoa(day)...)
	l.buf = append(l.buf, ' ')

	hour, min, sec := t.Clock()
	l.buf = append(l.buf, strconv.Itoa(hour)...)
	l.buf = append(l.buf, ':')
	l.buf = append(l.buf, strconv.Itoa(min)...)
	l.buf = append(l.buf, ':')
	l.buf = append(l.buf, strconv.Itoa(sec)...)
	l.buf = append(l.buf, ' ')

	l.buf = append(l.buf, '[')
	l.buf = append(l.buf, l.svrName...)
	l.buf = append(l.buf, ']', ' ')
}

// log is the function that actually logs the provided string and log time. The
// date and server name and written first, followed by the message.
// An example is:
// 2016/01/12 10:21:38 [golog] message goes here
// <time> [<server name>] <message>
//
// Note: the provided time should be UTC.
func (l *Logger) log(s string, t time.Time) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.formatHeader(t)
	l.buf = append(l.buf, s...)
	l.buf = append(l.buf, '\n')

	_, err := l.Out.Write(l.buf)
	return err
}

// Log writes the provided string to standard out with the proper logging
// format.
func (l *Logger) Log(s string) error {
	t := time.Now().UTC()
	return l.log(s, t)
}

// LogPanic recovers from any panic and logs the panic message and the stack
// trace to standard out with the proper logging format.
func (l *Logger) LogPanic() {
	defer func() {
		if err := recover(); err != nil {
			stack := make([]byte, stackSize)
			stack = stack[:runtime.Stack(stack, true)]
			l.Log(fmt.Sprintf("panic: %s\n%s", err, stack))
		}
	}()
}

// LogError logs the provided error to standard out with the proper logging
// format.
func (l *Logger) LogError(err error) error {
	return l.Log("error: " + err.Error())
}
