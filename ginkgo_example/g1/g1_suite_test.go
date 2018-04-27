package g1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestG1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "G1 Suite")
}
