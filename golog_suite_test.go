package golog_test

import (
	"fmt"

	. "github.com/crowdriff/golog"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGolog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Golog Suite")
}

var app = "golog"
var appLog = fmt.Sprintf("app=%s", app)

var version = "v1"
var versionLog = fmt.Sprintf("v=%s", version)

var _ = BeforeSuite(func() {

	Ω(func() {
		Log("test")
	}).Should(Panic())

	Init(app, version)
})
