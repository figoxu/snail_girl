package ut

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsbedb346fe7b8c62c58fb30fd93c561c85b37a3de = "{% autoescape off %}\n{% for bo in bos %}\n\n    {% if bo.ReDeclareFlag %}\n# 重新清关\ncurl -X POST -H \"Content-Type: application/json\" -u admin:bar_e3cjs https://meishi.51baibao.com/api/admin/developer_operation/wx_custom_declare --data '{\"order_code\":\"{{ code }}\", \"declare\":false, \"is_modify\":false}'\n    {% elif bo.NewFlag %}\n# 创建清关订单\ncurl -X POST -H \"Content-Type: application/json\" -u admin:bar_e3cjs https://meishi.51baibao.com/api/admin/developer_operation/wx_custom_declare --data '{\"order_code\":\"{{ code }}\", \"declare\":true, \"is_modify\":false}'\n    {% else %}\n# 编辑清关订单\ncurl -X POST -H \"Content-Type: application/json\" -u admin:bar_e3cjs https://meishi.51baibao.com/api/admin/developer_operation/wx_custom_declare --data '{\"order_code\":\"{{ code }}\", \"declare\":true, \"is_modify\":true}'\n    {% endif %}\n\n{% endfor %}\n{% endautoescape %}"
var _Assets3fbf47a1740847049b1ecac743e781446c6467b7 = "{% autoescape off %}\npackage testdata\n\nimport (\n\t\"github.com/quexer/utee\"\n)\n\nfunc {{ content.Name }} () *domain.{{ content.Name }} {\n\tobj := &domain.{{ content.Name }} {\n        {% for field in content.Fields %}\n            {{ field.Name }}  : {{ field.MockVal }} , {% endfor %}\n\t}\n\terr := obj.Valid()\n\tutee.Chk(err)\n\treturn obj\n}\n\n{% endautoescape %}"
var _Assets34d1fc10068d7a91e960081240458413e83f099e = "{% autoescape off %}\npackage impl\n\ntype {{ content.Name }}RepoImpl struct {\n\tDs *database.Ds\n}\n\n{% endautoescape %}"
var _Assetsc5fa4fd3c30186118194a6d70943b5925e989834 = "{% autoescape off %}\n{% for code in codes %}\ncurl -X POST -H \"Content-Type: application/json\" -u admin:bar_e3cjs https://meishi.51baibao.com/api/admin/developer_operation/push_trade --data '{\"order_code\":\"{{ code }}\"}'\n{% endfor %}\n{% endautoescape %}"
var _Assets983c7fa7b5485f28a4c70e5c052a4140e09f0b62 = "{% autoescape off %}\npackage impl_test\n\nvar _ = Describe(\"{{ content.Name }}\", func() {\n\n\tvar sr *impl.{{ content.Name }}Impl\n\tvar {{ content.Name }}Repo *mrepo.{{ content.Name }}Repo\n\n\tBeforeEach(func() {\n\t\t{{ content.Name }}Repo = mrepo.New{{ content.Name }}(ctl)\n\t\tsr = &impl.BatchCouponSrvImpl{\n\t\t\t{{ content.Name }}Repo:     {{ content.Name }}Repo ,\n\t\t}\n\t})\n\n\tIt(\"Create\", func() {\n//\t\t{{ content.Name }}Repo.EXPECT()\n\n\t})\n\n})\n\n\n{% endautoescape %}"
var _Assets49b2ff61e8922d1ec75e37ede2cadcf00817e124 = "{% autoescape off %}\n{% for content in contents %}\n    {{ content }}\n{% endfor %}\n{% endautoescape %}"
var _Assets29a3b0fb9296755c557f1805fcccfb4692536da5 = "{% autoescape off %}\npackage impl_test\n\nimport (\n\t. \"github.com/onsi/ginkgo\"\n\t. \"github.com/onsi/gomega\"\n)\n\nvar _ = Describe(\"{{ content.Name }}\", func() {\n\n\tvar ro *impl.{{ content.Name }}RepoImpl\n\n\tBeforeEach(func() {\n\t\tro = &impl.{{ content.Name }}RepoImpl{\n\t\t\tDs: Ds,\n\t\t}\n\n\t\tf := func() {\n\t\t\treq := model.{{ content.Name }}{}.New(testdata.{{ content.Name }}())\n\t\t\terr := ro.Ds.Gdb().Create(req).Error\n\t\t\tΩ(err).ShouldNot(HaveOccurred())\n\t\t}\n\n\t\tf()\n\t})\n\n\tIt(\"Default\", func() {\n\n\t})\n\n})\n\n{% endautoescape %}"
var _Assetsaa90fe91593da4ed3e77feb783d4bb503d235241 = "{% autoescape off %}\npackage model\n\nimport (\n\t\"github.com/jinzhu/copier\"\n\t\"github.com/quexer/utee\"\n\n\t\"baibao/m2/pkg/domain\"\n)\n\ntype {{ content.Name }} struct {\n{% for field in content.Fields %}\n{{ field.Name }}    {{ field.Type }}    `{{ field.GormTag }}`  {% endfor %}\n}\n\nfunc ({{ content.Name }}) ModelToDomain(x *{{ content.Name }}) *domain.{{ content.Name }} {\n\treturn x.ToDomain()\n}\n\nfunc ({{ content.Name }}) New(x *domain.{{ content.Name }}) *{{ content.Name }} {\n\tout := &{{ content.Name }}{}\n\tutee.Chk(copier.Copy(out, x))\n\treturn out\n}\n\nfunc (p *{{ content.Name }}) ToDomain() *domain.{{ content.Name }} {\n\tout := &domain.{{ content.Name }}{}\n\tutee.Chk(copier.Copy(out, p))\n\treturn out\n}\n\n\n{% endautoescape %}"
var _Assetsa3993fef2ee05af4e4591868f81e4e0c18c0a4f6 = "{% autoescape off %}\npackage service\n\n//go:generate mockgen -destination=../mocks/mservice/{{ fileName }} -package=mservice . {{ content.Name }}Srv\ntype {{ content.Name }}Srv interface {\n\n}\n\n\n{% endautoescape %}"
var _Assets3972b2668dda5555d3a1429a108b0f8d9a67c695 = "{% autoescape off %}\npackage repo\n//go:generate mockgen -destination=../mocks/mrepo/{{ fileName }} -package=mrepo . {{ content.Name }}Repo\n\ntype {{ content.Name }}Repo interface {\n}\n\n{% endautoescape %}"
var _Assets33b074b5d52e69b2e2a40719e53010e54e3a70ff = "{% autoescape off %}\ncurl -X {{ content.Method }} {{ content.FullUrl }} \\\n{% if content.ContentType != '' %}\n-H \"Content-Type: {{ content.ContentType }}\" \\\n{% endif %}\n-H 'Cookie: {{ content.Cookie }}' \\\n{% if content.Body != '' %}\n    --data '{{ content.Body }}'\n{% endif %}\n{% endautoescape %}"
var _Assets13cee148fed685ecd1ea848a02afc39e28ba236c = "{% autoescape off %}\n\npackage impl\n\ntype {{ content.Name }}Impl struct {\n    {{ content.Name }}Repo repo.{{ content.Name }}Repo\n}\n\n{% endautoescape %}"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"assets"}, "/assets": []string{"domain_service_impl_test.tpl", "domain_repo.tpl", "declare.tpl", "content.tpl", "curl.tpl", "domain_data.tpl", "domain_repo_impl_test.tpl", "domain_repo_impl.tpl", "push_order.tpl", "domain_repo_internal_model.tpl", "domain_service.tpl", "domain_service_impl.tpl"}}, map[string]*assets.File{
	"/assets/curl.tpl": &assets.File{
		Path:     "/assets/curl.tpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1621931987, 1621931987484241360),
		Data:     []byte(_Assets33b074b5d52e69b2e2a40719e53010e54e3a70ff),
	}, "/assets/domain_service_impl.tpl": &assets.File{
		Path:     "/assets/domain_service_impl.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1624889372, 1624889372998244598),
		Data:     []byte(_Assets13cee148fed685ecd1ea848a02afc39e28ba236c),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1626484869, 1626484869597086630),
		Data:     nil,
	}, "/assets/declare.tpl": &assets.File{
		Path:     "/assets/declare.tpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1621931987, 1621931987484384845),
		Data:     []byte(_Assetsbedb346fe7b8c62c58fb30fd93c561c85b37a3de),
	}, "/assets/domain_data.tpl": &assets.File{
		Path:     "/assets/domain_data.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1626484505, 1626484505811344501),
		Data:     []byte(_Assets3fbf47a1740847049b1ecac743e781446c6467b7),
	}, "/assets/domain_repo_impl.tpl": &assets.File{
		Path:     "/assets/domain_repo_impl.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1625031249, 1625031249150616592),
		Data:     []byte(_Assets34d1fc10068d7a91e960081240458413e83f099e),
	}, "/assets": &assets.File{
		Path:     "/assets",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1626484505, 1626484505811424728),
		Data:     nil,
	}, "/assets/push_order.tpl": &assets.File{
		Path:     "/assets/push_order.tpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1621931987, 1621931987484503331),
		Data:     []byte(_Assetsc5fa4fd3c30186118194a6d70943b5925e989834),
	}, "/assets/domain_service_impl_test.tpl": &assets.File{
		Path:     "/assets/domain_service_impl_test.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1624889940, 1624889940401437419),
		Data:     []byte(_Assets983c7fa7b5485f28a4c70e5c052a4140e09f0b62),
	}, "/assets/content.tpl": &assets.File{
		Path:     "/assets/content.tpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1621931987, 1621931987484104355),
		Data:     []byte(_Assets49b2ff61e8922d1ec75e37ede2cadcf00817e124),
	}, "/assets/domain_repo_impl_test.tpl": &assets.File{
		Path:     "/assets/domain_repo_impl_test.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1625033727, 1625033727759041495),
		Data:     []byte(_Assets29a3b0fb9296755c557f1805fcccfb4692536da5),
	}, "/assets/domain_repo_internal_model.tpl": &assets.File{
		Path:     "/assets/domain_repo_internal_model.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1625031286, 1625031286841376327),
		Data:     []byte(_Assetsaa90fe91593da4ed3e77feb783d4bb503d235241),
	}, "/assets/domain_service.tpl": &assets.File{
		Path:     "/assets/domain_service.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1625031230, 1625031230425460366),
		Data:     []byte(_Assetsa3993fef2ee05af4e4591868f81e4e0c18c0a4f6),
	}, "/assets/domain_repo.tpl": &assets.File{
		Path:     "/assets/domain_repo.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1625031230, 1625031230428080543),
		Data:     []byte(_Assets3972b2668dda5555d3a1429a108b0f8d9a67c695),
	}}, "")
