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
			Init("golog", "v1")
			Ω(func() {
				Log("test")
			}).ShouldNot(Panic())
		})
	})

	Context("Log", func() {
		It("should log a message with the proper format", func() {
			var buf bytes.Buffer
			SetOutput(&buf)
			Log("test message")
			out := buf.String()
			Ω(strings.Contains(out, "level=info")).Should(BeTrue())
			Ω(strings.Contains(out, appLog)).Should(BeTrue())
			Ω(strings.Contains(out, versionLog)).Should(BeTrue())
			Ω(strings.Contains(out, "msg=\"test message\"")).Should(BeTrue())
			Ω(strings.Contains(out, "time=\"")).Should(BeTrue())
		})
	})

	Context("Log Error", func() {
		It("should log an error with the proper format", func() {
			var buf bytes.Buffer
			SetOutput(&buf)
			LogError(errors.New("test error"))
			out := buf.String()
			Ω(strings.Contains(out, "level=error")).Should(BeTrue())
			Ω(strings.Contains(out, appLog)).Should(BeTrue())
			Ω(strings.Contains(out, versionLog)).Should(BeTrue())
			Ω(strings.Contains(out, "msg=\"test error\"")).Should(BeTrue())
			Ω(strings.Contains(out, "time=\"")).Should(BeTrue())
			Ω(strings.Contains(out, "file=\"")).Should(BeTrue())
			Ω(strings.Contains(out, "line=")).Should(BeTrue())
		})
	})

	Context("Log Panic", func() {
		It("should log an error with the proper format & panic", func() {
			var buf bytes.Buffer
			SetOutput(&buf)
			Ω(func() {
				LogPanic(errors.New("test panic"))
			}).Should(Panic())
			out := buf.String()
			Ω(strings.Contains(out, "level=panic")).Should(BeTrue())
			Ω(strings.Contains(out, appLog)).Should(BeTrue())
			Ω(strings.Contains(out, versionLog)).Should(BeTrue())
			Ω(strings.Contains(out, "msg=\"test panic\"")).Should(BeTrue())
			Ω(strings.Contains(out, "time=\"")).Should(BeTrue())
			Ω(strings.Contains(out, "file=\"")).Should(BeTrue())
			Ω(strings.Contains(out, "line=")).Should(BeTrue())
		})
	})

	Context("Log Warning", func() {
		It("should log a warning with the proper format", func() {
			var buf bytes.Buffer
			SetOutput(&buf)
			LogWarning("test warning")
			out := buf.String()
			Ω(strings.Contains(out, "level=warn")).Should(BeTrue())
			Ω(strings.Contains(out, appLog)).Should(BeTrue())
			Ω(strings.Contains(out, versionLog)).Should(BeTrue())
			Ω(strings.Contains(out, "msg=\"test warning\"")).Should(BeTrue())
			Ω(strings.Contains(out, "time=\"")).Should(BeTrue())
		})
	})
})
