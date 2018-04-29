package counterfeiterexamplefakes_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCounterfeiterExamplefakes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CounterfeiterExamplefakes Suite")
}
