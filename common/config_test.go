package baja_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/yeo/baja"
)

var _ = Describe("Config", func() {
	Describe("NodeDB", func() {
		Describe("Append", func() {
			It("appends node to db", func() {
				db := &NodeDB{}

				n := &Node{}

				db.Append(n)

				Expect(db.Total).To(Equal(1))
				Expect(db.NodeList[0]).To(Equal(n))
			})
		})
	})
})
