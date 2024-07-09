package test_test

import (
	"github.com/dustinspecker/ginkgo-gomega-custom-formatter-example/internal"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test", func() {
	It("custom reporter handles github.com/pkg/errors", func() {
		Expect(internal.ReturnsPkgError()).To(Succeed(), "example of github.com/pkg/errors")
	})

	It("custom reporter falls back to Gomega's default reporter for other types", func() {
		Expect(internal.ReturnsError()).To(Succeed(), "example of error")
	})
})
