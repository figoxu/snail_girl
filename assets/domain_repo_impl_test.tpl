{% autoescape off %}
package impl_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("{{ content.Name }}", func() {

	var ro *impl.{{ content.Name }}RepoImpl

	BeforeEach(func() {
		ro = &impl.{{ content.Name }}RepoImpl{
			Ds: Ds,
		}

		f := func() {
			req := model.{{ content.Name }}{}.New(testdata.{{ content.Name }}())
			err := ro.Ds.Gdb().Create(req).Error
			Ω(err).ShouldNot(HaveOccurred())
		}

		f()
	})

	It("Default", func() {

	})

})

{% endautoescape %}