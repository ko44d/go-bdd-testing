package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"os/exec"
)

var _ = Describe("Main", func() {
	var (
		cli string
		err error
	)

	BeforeEach(func() {
		cli, err = gexec.Build(".")
		Expect(err).ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		gexec.CleanupBuildArtifacts()
	})
	Context("with hoge set in first option", func() {
		It("return hoge in terminal", func() {
			command := exec.Command(cli+"/howto-gomega-gexec", "-first hhhhh")
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			session.Wait()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
