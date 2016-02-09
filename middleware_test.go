package golog_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

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

	var buf bytes.Buffer

	BeforeEach(func() {
		buf.Reset()
		SetOutput(&buf)
	})

	It("should log the request with middleware and the X-Forwarded-For ip address", func() {
		h := &fakeHTTPHandler{http.StatusOK, 100}
		f := LoggingMiddleware(h)

		wr := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "http://localhost/path?query=10", nil)
		Ω(err).ShouldNot(HaveOccurred())
		r.Header.Set("X-Forwarded-For", "127.0.0.1:62")
		f.ServeHTTP(wr, r)

		out := buf.String()
		Ω(out).Should(ContainSubstring("level=info"))
		Ω(out).Should(ContainSubstring(appLog))
		Ω(out).Should(ContainSubstring(versionLog))
		Ω(out).Should(ContainSubstring("time=\""))
		Ω(out).Should(ContainSubstring("code=200"))
		Ω(out).Should(ContainSubstring("dur="))
		Ω(out).Should(ContainSubstring("ip=\"127.0.0.1:62\""))
		Ω(out).Should(ContainSubstring("method=GET"))
		Ω(out).Should(ContainSubstring("size=100"))
		Ω(out).Should(ContainSubstring("uri=\"/path?query=10\""))
	})

	It("should log the request with middleware and the X-Real-IP ip address", func() {
		h := &fakeHTTPHandler{http.StatusBadRequest, 2345}
		f := LoggingMiddleware(h)

		wr := httptest.NewRecorder()
		r, err := http.NewRequest("POST", "http://localhost/path?query=10#yolo", nil)
		Ω(err).ShouldNot(HaveOccurred())
		r.Header.Set("X-Real-IP", "127.0.0.1:62")
		f.ServeHTTP(wr, r)

		out := buf.String()
		Ω(out).Should(ContainSubstring("level=info"))
		Ω(out).Should(ContainSubstring(appLog))
		Ω(out).Should(ContainSubstring(versionLog))
		Ω(out).Should(ContainSubstring("time=\""))
		Ω(out).Should(ContainSubstring("code=400"))
		Ω(out).Should(ContainSubstring("dur="))
		Ω(out).Should(ContainSubstring("ip=\"127.0.0.1:62\""))
		Ω(out).Should(ContainSubstring("method=POST"))
		Ω(out).Should(ContainSubstring("size=2345"))
		Ω(out).Should(ContainSubstring("uri=\"/path?query=10\""))
	})
})
