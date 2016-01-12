package golog_test

import (
	"bytes"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	. "github.com/crowdriff/golog"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Golog", func() {
	Context("NewLogger", func() {
		It("should create a Logger with standard out as default", func() {
			l := NewLogger("golog")
			l.Log("test")
			Ω(l).ShouldNot(BeNil())
			Ω(l.Out).Should(Equal(os.Stdout))
		})
	})

	Context("Log", func() {
		It("should log a message with the proper format", func() {
			var buf bytes.Buffer
			l := NewLogger("golog")
			l.Out = &buf
			Ω(l.Log("test message")).ShouldNot(HaveOccurred())
			out := buf.String()
			out = parseDate(out)
			out = parseServerName(out, "golog")
			Ω(out).Should(Equal("test message\n"))
		})
	})

	Context("Log Error", func() {
		It("should log an error with the proper format", func() {
			var buf bytes.Buffer
			l := NewLogger("golog")
			l.Out = &buf
			Ω(l.LogError(errors.New("error text"))).ShouldNot(HaveOccurred())
			out := buf.String()
			out = parseDate(out)
			out = parseServerName(out, "golog")
			Ω(out).Should(Equal("error: error text\n"))
		})
	})

	Context("Log Panic", func() {
		It("should recover from a panic and log the message and stack trace with the proper format", func() {
			var buf bytes.Buffer
			l := NewLogger("golog")
			l.Out = &buf

			go func() {
				defer l.LogPanic()
				panic("message")
			}()
			time.Sleep(10 * time.Millisecond)
			out := buf.String()
			out = parseDate(out)
			out = parseServerName(out, "golog")
			Ω(strings.HasPrefix(out, "panic: message\n")).Should(BeTrue())
			// ensure that stack trace exists
			Ω(len(out) - len("panic: message\n")).Should(BeNumerically(">", 0))
		})
	})
})

// parseDate parses the date (and trailing space) from the provided string and
// returns the remaining string.
func parseDate(s string) string {
	// parse year
	year, err := strconv.Atoi(s[:4])
	Ω(err).ShouldNot(HaveOccurred())
	Ω(year).Should(BeNumerically(">", 2015))
	Ω(year).Should(BeNumerically("<", 9999))

	Ω(s[4]).Should(Equal(uint8('/')))

	// parse month
	month, err := strconv.Atoi(s[5:7])
	Ω(err).ShouldNot(HaveOccurred())
	Ω(month).Should(BeNumerically(">", 0))
	Ω(month).Should(BeNumerically("<=", 12))

	Ω(s[7]).Should(Equal(uint8('/')))

	// parse day
	day, err := strconv.Atoi(s[8:10])
	Ω(err).ShouldNot(HaveOccurred())
	Ω(day).Should(BeNumerically(">", 0))
	Ω(day).Should(BeNumerically("<=", 31))

	Ω(s[10]).Should(Equal(uint8(' ')))

	// parse hour
	hour, err := strconv.Atoi(s[11:13])
	Ω(err).ShouldNot(HaveOccurred())
	Ω(hour).Should(BeNumerically(">=", 0))
	Ω(hour).Should(BeNumerically("<", 24))

	Ω(s[13]).Should(Equal(uint8(':')))

	// parse minute
	min, err := strconv.Atoi(s[14:16])
	Ω(err).ShouldNot(HaveOccurred())
	Ω(min).Should(BeNumerically(">=", 0))
	Ω(min).Should(BeNumerically("<", 60))

	Ω(s[16]).Should(Equal(uint8(':')))

	// parse second
	sec, err := strconv.Atoi(s[17:19])
	Ω(err).ShouldNot(HaveOccurred())
	Ω(sec).Should(BeNumerically(">=", 0))
	Ω(sec).Should(BeNumerically("<", 60))

	Ω(s[19]).Should(Equal(uint8(' ')))

	return s[20:]
}

// parseServerName parses the server name (and trailing space) from the provided
// string and returns the remaining string.
func parseServerName(s, name string) string {
	ok := strings.HasPrefix(s, "["+name+"]")
	Ω(ok).Should(BeTrue())
	Ω(s[len(name)+2]).Should(Equal(uint8(' ')))
	return s[len(name)+3:]
}
