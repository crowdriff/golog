package golog_test

import (
	. "github.com/crowdriff/golog"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGolog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Golog Suite")
}

var _ = BeforeSuite(func() {

	Ω(func() {
		Log("test")
	}).Should(Panic())

	Init("golog", "v1")
})
