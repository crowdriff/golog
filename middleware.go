package golog

import (
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

// responseWriter represents an object that wraps an http.ResponseWriter to
// record the response code returned, as well as the size of the response body.
type responseWriter struct {
	http.ResponseWriter
	statusWritten bool
	Code          int
	Size          int
}

// WriteHeader records the code and calls the underlying ResponseWriter's
// WriteHeader method.
func (rw *responseWriter) WriteHeader(c int) {
	rw.Code = c
	rw.statusWritten = true
	rw.ResponseWriter.WriteHeader(c)
}

// Write records the number of bytes written and calls the underlying
// ResponseWriter's Write method.
func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.statusWritten {
		rw.WriteHeader(http.StatusOK)
	}
	n, err := rw.ResponseWriter.Write(b)
	rw.Size += n
	return n, err
}

// IP headers
var xForwardedFor = http.CanonicalHeaderKey("X-Forwarded-For")
var xRealIP = http.CanonicalHeaderKey("X-Real-IP")

// realIP returns the "real" IP address of the caller, or an empty string.
// Ported from https://github.com/pressly/chi/blob/master/middleware/realip.go
func realIP(r *http.Request) string {
	var ip string

	if xff := r.Header.Get(xForwardedFor); xff != "" {
		i := strings.Index(xff, ", ")
		if i == -1 {
			i = len(xff)
		}
		ip = xff[:i]
	} else if xrip := r.Header.Get(xRealIP); xrip != "" {
		ip = xrip
	}

	return ip
}

// logRequest logs the provided responseWriter, http request, and starting time
// to standard out in the proper format.
func (l *logger) logRequest(w *responseWriter, r *http.Request, start time.Time) {
	// set real IP
	if ip := realIP(r); ip != "" {
		r.RemoteAddr = ip
	}

	// log the request
	rootLogger.standardEntry().WithFields(log.Fields{
		"code":   w.Code,
		"dur":    int(time.Now().Sub(start)) / 1e3,
		"ip":     r.RemoteAddr,
		"method": r.Method,
		"size":   w.Size,
		"uri":    r.URL.RequestURI(),
	}).Print()
}

// LoggingMiddleware is a middleware function that logs all requests to standard
// out.
var LoggingMiddleware = func(h http.Handler) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		wr := &responseWriter{w, false, 0, 0}
		h.ServeHTTP(wr, r)
		rootLogger.logRequest(wr, r, t)
	}
	return http.HandlerFunc(f)
}
