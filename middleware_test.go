package golog_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
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
	w.Write(make([]byte, h.Size))
}

var _ = Describe("Middleware", func() {
	It("should log the request with middleware", func() {
		var buf bytes.Buffer
		l := NewLogger("golog")
		l.Out = &buf

		h := &fakeHTTPHandler{http.StatusOK, 100}
		f := LogRequestMiddleware(l)(h)

		wr := httptest.NewRecorder()
		r, err := http.NewRequest("GET", "http://localhost/path?query=10", nil)
		Ω(err).ShouldNot(HaveOccurred())
		f.ServeHTTP(wr, r)

		out := buf.String()
		out = parseDate(out)
		out = parseServerName(out, "golog")
		reqStr := "200 GET /path?query=10 (100) "
		ok := strings.HasPrefix(out, reqStr)
		Ω(ok).Should(BeTrue())
		i, err := strconv.Atoi(out[len(reqStr) : len(out)-1])
		Ω(err).ShouldNot(HaveOccurred())
		Ω(i).Should(BeNumerically(">=", 0))
	})
})
