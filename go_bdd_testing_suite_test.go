package go_bdd_testing_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoBddTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoBddTesting Suite")
}
