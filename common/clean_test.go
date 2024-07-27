package baja_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/yeo/baja"
)

var _ = Describe("Clean", func() {
	It("Delete public directory", func() {
		os.MkdirAll("public", os.ModePerm)

		Clean()

		_, err := os.Stat("./public")
		Expect(err).ToNot(Equal(nil))
	})
})
