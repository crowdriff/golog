package golog_test

import (
	"bytes"
	"errors"
	"strings"

	. "github.com/crowdriff/golog"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Golog", func() {
	Context("NewLogger", func() {
		It("should create a Logger", func() {
			l := NewLogger("golog", "v1")
			Ω(l).ShouldNot(BeNil())
		})
	})

	Context("Log", func() {
		It("should log a message with the proper format", func() {
			var buf bytes.Buffer
			l := NewLogger("golog", "v1")
			l.SetOutput(&buf)
			l.Log("test message")
			out := buf.String()
			Ω(strings.Contains(out, "level=info")).Should(BeTrue())
			Ω(strings.Contains(out, "app=golog")).Should(BeTrue())
			Ω(strings.Contains(out, "v=v1")).Should(BeTrue())
			Ω(strings.Contains(out, "msg=\"test message\"")).Should(BeTrue())
			Ω(strings.Contains(out, "time=\"")).Should(BeTrue())
		})
	})

	Context("Log Error", func() {
		It("should log an error with the proper format", func() {
			var buf bytes.Buffer
			l := NewLogger("golog", "v1")
			l.SetOutput(&buf)
			l.LogError(errors.New("test error"))
			out := buf.String()
			Ω(strings.Contains(out, "level=error")).Should(BeTrue())
			Ω(strings.Contains(out, "app=golog")).Should(BeTrue())
			Ω(strings.Contains(out, "v=v1")).Should(BeTrue())
			Ω(strings.Contains(out, "msg=\"test error\"")).Should(BeTrue())
			Ω(strings.Contains(out, "time=\"")).Should(BeTrue())
		})
	})

	Context("Log Warning", func() {
		It("should log a warning with the proper format", func() {
			var buf bytes.Buffer
			l := NewLogger("golog", "v1")
			l.SetOutput(&buf)
			l.LogWarning("test warning")
			out := buf.String()
			Ω(strings.Contains(out, "level=warn")).Should(BeTrue())
			Ω(strings.Contains(out, "app=golog")).Should(BeTrue())
			Ω(strings.Contains(out, "v=v1")).Should(BeTrue())
			Ω(strings.Contains(out, "msg=\"test warning\"")).Should(BeTrue())
			Ω(strings.Contains(out, "time=\"")).Should(BeTrue())
		})
	})
})
