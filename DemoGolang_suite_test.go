package DemoGolang_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDemoGolang(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DemoGolang Suite")
}
