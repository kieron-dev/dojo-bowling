package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestBowling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bowling Suite")
}

var (
	cmd string
	err error
)

var _ = BeforeSuite(func() {
	cmd, err = gexec.Build("./main.go")
	Expect(err).NotTo(HaveOccurred())
})
