package DemoGolang

import (
	"DemoGolang/requests"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Verify an existing name", func() {
	It("Verify an existing name", func() {
		result, _ := requests.GetResult("", "?name=pepe", "GET", "")
		Expect(result.StatusCode).To(Equal(200))
	})
})