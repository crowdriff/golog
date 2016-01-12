package golog_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGolog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Golog Suite")
}
