{% autoescape off %}
package impl_test

import (
    "baibao/meishi/pkg/next/repo/impl"
    "baibao/meishi/pkg/next/repo/internal/model"
    "baibao/meishi/pkg/next/testdata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("{{ content.Name }}", func() {

	var ro *impl.{{ content.Name }}RepoImpl

	BeforeEach(func() {
		ro = &impl.{{ content.Name }}RepoImpl{
			Db: db,
		}

		err := ro.Db.AutoMigrate(&model.{{ content.Name }}{})
        		Ω(err).NotTo(HaveOccurred())

		f := func() {
		    testData := testdata.{{ content.Name }}()
		    testData.ID = 1
			req := model.{{ content.Name }}{}.New(testData)
			err := ro.Db.Create(req).Error
			Ω(err).ShouldNot(HaveOccurred())
		}

		f()
	})

	It("Create", func() {
        err := ro.Create(ctx, testdata.{{ content.Name }}())
        Ω(err).ShouldNot(HaveOccurred())
    })

    It("Update", func() {
        err := ro.Update(ctx, testdata.{{ content.Name }}(1))
        Ω(err).ShouldNot(HaveOccurred())
    })

    It("MustGet", func() {
        _, err := ro.MustGet(ctx, 1)
        Ω(err).ShouldNot(HaveOccurred())
    })

    It("MultiGet", func() {
        l, err := ro.MultiGet(ctx, 1, 2, 3)
        Ω(err).NotTo(HaveOccurred())
        Ω(l).To(HaveLen(1))
    })

    It("List", func() {
        l, err := ro.List(ctx)
        Ω(err).NotTo(HaveOccurred())
        Ω(l).To(HaveLen(1))
    })

})

{% endautoescape %}