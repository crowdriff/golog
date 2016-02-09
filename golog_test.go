package golog_test

import (
	"bytes"
	"errors"

	. "github.com/crowdriff/golog"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Golog", func() {

	var buf bytes.Buffer

	BeforeEach(func() {
		buf.Reset()
		SetOutput(&buf)
	})

	Context("Log", func() {
		It("should log a message with the proper format", func() {
			Log("test message")
			out := buf.String()
			Ω(out).Should(ContainSubstring("level=info"))
			Ω(out).Should(ContainSubstring(appLog))
			Ω(out).Should(ContainSubstring(versionLog))
			Ω(out).Should(ContainSubstring("msg=\"test message\""))
			Ω(out).Should(ContainSubstring("time=\""))
		})
	})

	Context("Logf", func() {
		It("should log a message with a format string without placeholders", func() {
			Logf("test message", "blah")
			out := buf.String()
			Ω(out).Should(ContainSubstring("level=info"))
			Ω(out).Should(ContainSubstring(appLog))
			Ω(out).Should(ContainSubstring(versionLog))
			Ω(out).Should(ContainSubstring("msg=\"test message%!(EXTRA string=blah)\""))
			Ω(out).Should(ContainSubstring("time=\""))
		})

		It("should log a message with an invalid format parameter type", func() {
			Logf("test message %d", "blah")
			out := buf.String()
			Ω(out).Should(ContainSubstring("level=info"))
			Ω(out).Should(ContainSubstring(appLog))
			Ω(out).Should(ContainSubstring(versionLog))
			Ω(out).Should(ContainSubstring("msg=\"test message %!d(string=blah)\""))
			Ω(out).Should(ContainSubstring("time=\""))
		})

		It("should log a formatted message with the proper format", func() {
			Logf("%sst mess%s", "te", "age")
			out := buf.String()
			Ω(out).Should(ContainSubstring("level=info"))
			Ω(out).Should(ContainSubstring(appLog))
			Ω(out).Should(ContainSubstring(versionLog))
			Ω(out).Should(ContainSubstring("msg=\"test message\""))
			Ω(out).Should(ContainSubstring("time=\""))
		})
	})

	Context("Log Error", func() {
		It("should log an error with the proper format", func() {
			LogError(errors.New("test error"))
			out := buf.String()
			Ω(out).Should(ContainSubstring("level=error"))
			Ω(out).Should(ContainSubstring(appLog))
			Ω(out).Should(ContainSubstring(versionLog))
			Ω(out).Should(ContainSubstring("msg=\"test error\""))
			Ω(out).Should(ContainSubstring("time=\""))
			Ω(out).Should(ContainSubstring("file=\""))
			Ω(out).Should(ContainSubstring("line="))
		})
	})

	Context("Log Panic", func() {
		It("should log an error with the proper format & panic", func() {
			Ω(func() {
				LogPanic(errors.New("test panic"))
			}).Should(Panic())
			out := buf.String()
			Ω(out).Should(ContainSubstring("level=panic"))
			Ω(out).Should(ContainSubstring(appLog))
			Ω(out).Should(ContainSubstring(versionLog))
			Ω(out).Should(ContainSubstring("msg=\"test panic\""))
			Ω(out).Should(ContainSubstring("time=\""))
			Ω(out).Should(ContainSubstring("file=\""))
			Ω(out).Should(ContainSubstring("line="))
		})
	})

	Context("Log Warning", func() {
		It("should log a warning with the proper format", func() {
			LogWarning("test warning")
			out := buf.String()
			Ω(out).Should(ContainSubstring("level=warn"))
			Ω(out).Should(ContainSubstring(appLog))
			Ω(out).Should(ContainSubstring(versionLog))
			Ω(out).Should(ContainSubstring("msg=\"test warning\""))
			Ω(out).Should(ContainSubstring("time=\""))
		})
	})
})
