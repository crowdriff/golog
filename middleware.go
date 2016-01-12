package golog

import (
	"net/http"
	"strconv"
	"time"
)

// responseWriter represents an object that wraps an http.ResponseWriter to
// record the response code returned, as well as the size of the response body.
type responseWriter struct {
	http.ResponseWriter
	written bool
	Code    int
	Size    int
}

// WriteHeader records the code and calls the underlying ResponseWriter's
// WriteHeader method.
func (rw *responseWriter) WriteHeader(c int) {
	rw.Code = c
	rw.written = true
	rw.ResponseWriter.WriteHeader(c)
}

// Write records the number of bytes written and calls the underlying
// ResponseWriter's Write method.
func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.written {
		rw.WriteHeader(http.StatusOK)
	}
	n, err := rw.ResponseWriter.Write(b)
	rw.Size += n
	return n, err
}

// Flush calls the underlying ResponseWriter's Flush method, if it conforms to
// the http.Flusher interface.
func (rw *responseWriter) Flush() {
	if f, ok := rw.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

// logRequest logs a request to standard out with the provided resonse writer,
// request, and starting time.
// An example to show the format is:
//
// 2016/01/12 10:21:38 [golog] 200 GET /path?query=10 (1024) 4580
// <time> [<server name>] <status> <method> <path> (<bytes written>) <microseconds>
//
// Note: the provided time should be UTC.
func (l *Logger) logRequest(w *responseWriter, r *http.Request, start time.Time) error {
	since := time.Since(start)

	l.mu.Lock()
	defer l.mu.Unlock()

	l.formatHeader(start)
	l.buf = append(l.buf, strconv.Itoa(w.Code)...)
	l.buf = append(l.buf, ' ')
	l.buf = append(l.buf, r.Method...)
	l.buf = append(l.buf, ' ')
	l.buf = append(l.buf, r.URL.RequestURI()...)
	l.buf = append(l.buf, ' ', '(')
	l.buf = append(l.buf, strconv.Itoa(w.Size)...)
	l.buf = append(l.buf, ')', ' ')
	l.buf = append(l.buf, strconv.Itoa(int(since)/1e3)...)
	l.buf = append(l.buf, '\n')

	_, err := l.Out.Write(l.buf)
	return err
}

// LogRequestMiddleware returns a middleware function that logs all requests
// with the provided Logger pointer.
func LogRequestMiddleware(l *Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		f := func(w http.ResponseWriter, r *http.Request) {
			t := time.Now().UTC()
			wr := &responseWriter{w, false, 0, 0}
			h.ServeHTTP(wr, r)
			l.logRequest(wr, r, t)
		}
		return http.HandlerFunc(f)
	}
}
