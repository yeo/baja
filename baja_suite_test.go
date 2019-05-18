package baja_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBaja(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Baja Suite")
}
