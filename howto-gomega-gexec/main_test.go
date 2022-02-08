package main_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"os/exec"
)

var _ = Describe("Main", func() {
	var (
		buildPath string
		cliPath   string
		err       error
	)

	BeforeEach(func() {
		buildPath, err = gexec.Build(".")
		Expect(err).ShouldNot(HaveOccurred())
		cliPath = fmt.Sprintf("%s/%s", buildPath, "howto-gomega-gexec")
	})

	AfterEach(func() {
		gexec.CleanupBuildArtifacts()
	})
	Context("with hoge set in first option", func() {
		It("return hoge in terminal", func() {
			command := exec.Command(cliPath, "-first hhhhh")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			session.Wait()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
