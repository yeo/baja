package baja_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/yeo/baja"

	"github.com/yeo/baja/utils"
)

var _ = Describe("Setup", func() {
	It("creates directory structure", func() {
		name := "testdir"

		Setup(name)

		Expect(utils.HasFile("./" + name + "/baja.yaml")).To(Equal(true))
		Expect(utils.HasFile("./" + name + "/content")).To(Equal(true))
		Expect(utils.HasFile("./" + name + "/theme/baja")).To(Equal(true))
		Expect(utils.HasFile("./" + name + "/public/asset")).To(Equal(true))

		os.RemoveAll(name)
	})

	It("Copy default theme", func() {
	})
})
