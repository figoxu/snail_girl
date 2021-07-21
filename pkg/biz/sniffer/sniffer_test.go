package sniffer_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestSniffer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "RepoImpl Suite", []Reporter{reporters.NewJUnitReporter("junit.xml")})
}
