package Ginkgo_test_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgoTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoTest Suite")
}
