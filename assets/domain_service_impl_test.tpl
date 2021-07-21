{% autoescape off %}
package impl_test

var _ = Describe("{{ content.Name }}", func() {

	var sr *impl.{{ content.Name }}Impl
	var {{ content.Name }}Repo *mrepo.{{ content.Name }}Repo

	BeforeEach(func() {
		{{ content.Name }}Repo = mrepo.New{{ content.Name }}(ctl)
		sr = &impl.BatchCouponSrvImpl{
			{{ content.Name }}Repo:     {{ content.Name }}Repo ,
		}
	})

	It("Create", func() {
//		{{ content.Name }}Repo.EXPECT()

	})

})


{% endautoescape %}