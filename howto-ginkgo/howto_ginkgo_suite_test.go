package howto_ginkgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHowtoGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HowtoGinkgo Suite")
}
