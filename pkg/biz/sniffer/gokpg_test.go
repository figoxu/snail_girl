package sniffer_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"figoxu.me/snail_girl/pkg/biz/sniffer"
)

var _ = Describe("GoPkgSniffer", func() {
	gps := &sniffer.GoPkgSniffer{}

	It("AsPkg", func() {
		v, err := gps.AsPkg("/Users/xujianhui/xxbmm/projects/workspace_go/meishi/m2/pkg/domain/buy_group.go")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(v == "baibao/m2/pkg/domain/buy_group.go").Should(BeTrue())
	})
})
