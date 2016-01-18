package golog_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/crowdriff/golog"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeHTTPHandler struct {
	Code int
	Size int
}

func (h *fakeHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.Code != 200 {
		w.WriteHeader(h.Code)
	}
	w.Write(make([]byte, h.Size))
}

var _ = Describe("Middleware", func() {
	It("should log the request with middleware and the X-Forwarded-For ip address", func() {
		var buf bytes.Buffer
		l := NewLogger("golog", "v1")
		l.SetOutput(&buf)

		h := &fakeHTTPHandler{http.StatusOK, 100}
		f := LogRequestMiddleware(l)(h)

		wr := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "http://localhost/path?query=10", nil)
		Ω(err).ShouldNot(HaveOccurred())
		r.Header.Set("X-Forwarded-For", "127.0.0.1:62")
		f.ServeHTTP(wr, r)

		out := buf.String()
		Ω(strings.Contains(out, "level=info")).Should(BeTrue())
		Ω(strings.Contains(out, "app=golog")).Should(BeTrue())
		Ω(strings.Contains(out, "v=v1")).Should(BeTrue())
		Ω(strings.Contains(out, "time=\"")).Should(BeTrue())
		Ω(strings.Contains(out, "code=200")).Should(BeTrue())
		Ω(strings.Contains(out, "dur=")).Should(BeTrue())
		Ω(strings.Contains(out, "ip=\"127.0.0.1:62\"")).Should(BeTrue())
		Ω(strings.Contains(out, "method=GET")).Should(BeTrue())
		Ω(strings.Contains(out, "size=100")).Should(BeTrue())
		Ω(strings.Contains(out, "uri=\"/path?query=10\"")).Should(BeTrue())
	})

	It("should log the request with middleware and the X-Real-IP ip address", func() {
		var buf bytes.Buffer
		l := NewLogger("golog", "v1")
		l.SetOutput(&buf)

		h := &fakeHTTPHandler{http.StatusBadRequest, 2345}
		f := LogRequestMiddleware(l)(h)

		wr := httptest.NewRecorder()
		r, err := http.NewRequest("POST", "http://localhost/path?query=10#yolo", nil)
		Ω(err).ShouldNot(HaveOccurred())
		r.Header.Set("X-Real-IP", "127.0.0.1:62")
		f.ServeHTTP(wr, r)

		out := buf.String()
		Ω(strings.Contains(out, "level=info")).Should(BeTrue())
		Ω(strings.Contains(out, "app=golog")).Should(BeTrue())
		Ω(strings.Contains(out, "v=v1")).Should(BeTrue())
		Ω(strings.Contains(out, "time=\"")).Should(BeTrue())
		Ω(strings.Contains(out, "code=400")).Should(BeTrue())
		Ω(strings.Contains(out, "dur=")).Should(BeTrue())
		Ω(strings.Contains(out, "ip=\"127.0.0.1:62\"")).Should(BeTrue())
		Ω(strings.Contains(out, "method=POST")).Should(BeTrue())
		Ω(strings.Contains(out, "size=2345")).Should(BeTrue())
		Ω(strings.Contains(out, "uri=\"/path?query=10\"")).Should(BeTrue())
	})
})
