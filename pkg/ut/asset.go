package ut

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets3fbf47a1740847049b1ecac743e781446c6467b7 = "{% autoescape off %}\npackage testdata\n\nimport (\n    \"baibao/meishi/pkg/next/domain\"\n\t\"github.com/quexer/utee\"\n)\n\nfunc {{ content.Name }} (id ...int) *domain.{{ content.Name }} {\n\tobj := &domain.{{ content.Name }} {\n        {% for field in content.Fields %}\n            {{ field.Name }}  : {{ field.MockVal }} , {% endfor %}\n\t}\n\terr := obj.Valid()\n\tutee.Chk(err)\n\tif len(id) > 0 {\n        obj.Id = id[0]\n    }\n\treturn obj\n}\n\n{% endautoescape %}"
var _Assets29a3b0fb9296755c557f1805fcccfb4692536da5 = "{% autoescape off %}\npackage impl_test\n\nimport (\n    \"baibao/meishi/pkg/next/repo/impl\"\n    \"baibao/meishi/pkg/next/repo/internal/model\"\n    \"baibao/meishi/pkg/next/testdata\"\n\t. \"github.com/onsi/ginkgo\"\n\t. \"github.com/onsi/gomega\"\n)\n\nvar _ = Describe(\"{{ content.Name }}\", func() {\n\n\tvar ro *impl.{{ content.Name }}RepoImpl\n\n\tBeforeEach(func() {\n\t\tro = &impl.{{ content.Name }}RepoImpl{\n\t\t\tDb: db,\n\t\t}\n\n\t\terr := ro.Db.AutoMigrate(&model.{{ content.Name }}{})\n        \t\tΩ(err).NotTo(HaveOccurred())\n\n\t\tf := func() {\n\t\t    testData := testdata.{{ content.Name }}()\n\t\t    testData.ID = 1\n\t\t\treq := model.{{ content.Name }}{}.New(testData)\n\t\t\terr := ro.Db.Create(req).Error\n\t\t\tΩ(err).ShouldNot(HaveOccurred())\n\t\t}\n\n\t\tf()\n\t})\n\n\tIt(\"Create\", func() {\n        err := ro.Create(ctx, testdata.{{ content.Name }}())\n        Ω(err).ShouldNot(HaveOccurred())\n    })\n\n    It(\"Update\", func() {\n        err := ro.Update(ctx, testdata.{{ content.Name }}(1))\n        Ω(err).ShouldNot(HaveOccurred())\n    })\n\n    It(\"MustGet\", func() {\n        _, err := ro.MustGet(ctx, 1)\n        Ω(err).ShouldNot(HaveOccurred())\n    })\n\n    It(\"MultiGet\", func() {\n        l, err := ro.MultiGet(ctx, 1, 2, 3)\n        Ω(err).NotTo(HaveOccurred())\n        Ω(l).To(HaveLen(1))\n    })\n\n    It(\"List\", func() {\n        l, err := ro.List(ctx)\n        Ω(err).NotTo(HaveOccurred())\n        Ω(l).To(HaveLen(1))\n    })\n\n})\n\n{% endautoescape %}"
var _Assetsa3993fef2ee05af4e4591868f81e4e0c18c0a4f6 = "{% autoescape off %}\npackage service\n\n//go:generate mockgen -destination=../mocks/mservice/{{ fileName }} -package=mservice . {{ content.Name }}Srv\ntype {{ content.Name }}Srv interface {\n\n}\n\n\n{% endautoescape %}"
var _Assetsaa90fe91593da4ed3e77feb783d4bb503d235241 = "{% autoescape off %}\npackage model\n\nimport (\n\t\"github.com/jinzhu/copier\"\n\t\"github.com/quexer/utee\"\n\n\t\"baibao/m2/pkg/domain\"\n)\n\ntype {{ content.Name }} struct {\n{% for field in content.Fields %}\n{{ field.Name }}    {{ field.Type }}    `{{ field.GormTag }}`  {% endfor %}\n}\n\nfunc ({{ content.Name }}) ModelToDomain(x *{{ content.Name }}) *domain.{{ content.Name }} {\n\treturn x.ToDomain()\n}\n\nfunc ({{ content.Name }}) New(x *domain.{{ content.Name }}) *{{ content.Name }} {\n\tout := &{{ content.Name }}{}\n\tutee.Chk(copier.Copy(out, x))\n\treturn out\n}\n\nfunc (p *{{ content.Name }}) ToDomain() *domain.{{ content.Name }} {\n\tout := &domain.{{ content.Name }}{}\n\tutee.Chk(copier.Copy(out, p))\n\treturn out\n}\n\n\n{% endautoescape %}"
var _Assets13cee148fed685ecd1ea848a02afc39e28ba236c = "{% autoescape off %}\n\npackage impl\n\ntype {{ content.Name }}Impl struct {\n    {{ content.Name }}Repo repo.{{ content.Name }}Repo\n}\n\n{% endautoescape %}"
var _Assets983c7fa7b5485f28a4c70e5c052a4140e09f0b62 = "{% autoescape off %}\npackage impl_test\n\nvar _ = Describe(\"{{ content.Name }}\", func() {\n\n\tvar sr *impl.{{ content.Name }}Impl\n\tvar {{ content.Name }}Repo *mrepo.{{ content.Name }}Repo\n\n\tBeforeEach(func() {\n\t\t{{ content.Name }}Repo = mrepo.New{{ content.Name }}(ctl)\n\t\tsr = &impl.BatchCouponSrvImpl{\n\t\t\t{{ content.Name }}Repo:     {{ content.Name }}Repo ,\n\t\t}\n\t})\n\n\tIt(\"Create\", func() {\n//\t\t{{ content.Name }}Repo.EXPECT()\n\n\t})\n\n})\n\n\n{% endautoescape %}"
var _Assets3972b2668dda5555d3a1429a108b0f8d9a67c695 = "{% autoescape off %}\npackage repo\n//go:generate mockgen -destination=../mocks/mrepo/{{ fileName }} -package=mrepo . {{ content.Name }}Repo\n\ntype {{ content.Name }}Repo interface {\n    MustGet(ctx context.Context, id int) (*domain.{{ content.Name }}, error)\n    MultiGet(ctx context.Context, id ...int) (domain.{{ content.Name }}List, error)\n    Create(ctx context.Context, in *domain.{{ content.Name }}) error\n    Update(ctx context.Context, in *domain.{{ content.Name }}) error\n    List(ctx context.Context) (domain.{{ content.Name }}List, error)\n}\n\n{% endautoescape %}"
var _Assets34d1fc10068d7a91e960081240458413e83f099e = "{% autoescape off %}\npackage impl\n\nimport (\n\t\"baibao/meishi/pkg/next/domain\"\n\t\"baibao/meishi/pkg/next/repo/internal/model\"\n\t\"context\"\n    \"github.com/ahmetb/go-linq/v3\"\n    \"github.com/pkg/errors\"\n)\n\ntype {{ content.Name }}RepoImpl struct {\n\tDb *gorm.DB\n}\n\nfunc (p *{{ content.Name }}RepoImpl) Update(ctx context.Context, in *domain.{{ content.Name }}) error {\n\tif err := in.Valid(); err != nil {\n\t\treturn err\n\t}\n\n\terr := p.Db.Save(model.{{ content.Name }}{}.New(in)).Error\n\n\treturn errors.WithStack(err)\n}\n\nfunc (p *{{ content.Name }}RepoImpl) Create(ctx context.Context, in *domain.{{ content.Name }}) error {\n\tif err := in.Valid(); err != nil {\n\t\treturn err\n\t}\n\tobj := model.{{ content.Name }}{}.New(in)\n\tif err := p.Db.Create(obj).Error; err != nil {\n\t\treturn errors.WithStack(err)\n\t}\n\tin.ID = obj.ID\n\treturn nil\n}\n\nfunc (p *{{ content.Name }}RepoImpl) MustGet(ctx context.Context, id int) (*domain.{{ content.Name }}, error) {\n\tvar o model.{{ content.Name }}\n\tif err := p.Db.Take(&o, id).Error; err != nil {\n\t\treturn nil, errors.WithStack(err)\n\t}\n\treturn o.ToDomain(), nil\n}\n\n\nfunc (p *{{ content.Name }}RepoImpl) MultiGet(ctx context.Context, id ...int) (domain.{{ content.Name }}List, error) {\n\tif len(id) == 0 {\n\t\treturn domain.{{ content.Name }}List{}, nil\n\t}\n\n\tvar l []*model.{{ content.Name }}\n\tif err := p.Db.\n\t\tWhere(\"id in ?\", id).\n\t\tFind(&l).Error; err != nil {\n\t\treturn nil, errors.WithStack(err)\n\t}\n\n\tvar out domain.{{ content.Name }}List\n\tlinq.From(l).SelectT(model.{{ content.Name }}{}.ModelToDomain).ToSlice(&out)\n\treturn out, nil\n}\n\nfunc (p *{{ content.Name }}RepoImpl) List(ctx context.Context) (domain.{{ content.Name }}List, error) {\n\tvar l []*model.{{ content.Name }}\n\tdb := p.Db\n\n\tif err := db.Find(&l).Error; err != nil {\n\t\treturn nil, errors.WithStack(err)\n\t}\n\n\tvar out domain.{{ content.Name }}List\n\tlinq.From(l).SelectT(model.{{ content.Name }}{}.ModelToDomain).ToSlice(&out)\n\treturn out, nil\n}\n\n{% endautoescape %}"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"assets"}, "/assets": []string{"domain_service_impl_test.tpl", "domain_repo.tpl", "domain_data.tpl", "domain_repo_impl_test.tpl", "domain_repo_impl.tpl", "domain_repo_internal_model.tpl", "domain_service.tpl", "domain_service_impl.tpl"}}, map[string]*assets.File{
	"/assets": &assets.File{
		Path:     "/assets",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1626879170, 1626879170728953805),
		Data:     nil,
	}, "/assets/domain_service_impl_test.tpl": &assets.File{
		Path:     "/assets/domain_service_impl_test.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1626859602, 1626859602220194823),
		Data:     []byte(_Assets983c7fa7b5485f28a4c70e5c052a4140e09f0b62),
	}, "/assets/domain_repo.tpl": &assets.File{
		Path:     "/assets/domain_repo.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1626864609, 1626864609355768862),
		Data:     []byte(_Assets3972b2668dda5555d3a1429a108b0f8d9a67c695),
	}, "/assets/domain_repo_impl.tpl": &assets.File{
		Path:     "/assets/domain_repo_impl.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1626878875, 1626878875834079858),
		Data:     []byte(_Assets34d1fc10068d7a91e960081240458413e83f099e),
	}, "/assets/domain_repo_internal_model.tpl": &assets.File{
		Path:     "/assets/domain_repo_internal_model.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1626859602, 1626859602218746511),
		Data:     []byte(_Assetsaa90fe91593da4ed3e77feb783d4bb503d235241),
	}, "/assets/domain_service_impl.tpl": &assets.File{
		Path:     "/assets/domain_service_impl.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1626859602, 1626859602219855051),
		Data:     []byte(_Assets13cee148fed685ecd1ea848a02afc39e28ba236c),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1626862851, 1626862851747105037),
		Data:     nil,
	}, "/assets/domain_data.tpl": &assets.File{
		Path:     "/assets/domain_data.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1626879170, 1626879170728786840),
		Data:     []byte(_Assets3fbf47a1740847049b1ecac743e781446c6467b7),
	}, "/assets/domain_repo_impl_test.tpl": &assets.File{
		Path:     "/assets/domain_repo_impl_test.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1626878963, 1626878963906515518),
		Data:     []byte(_Assets29a3b0fb9296755c557f1805fcccfb4692536da5),
	}, "/assets/domain_service.tpl": &assets.File{
		Path:     "/assets/domain_service.tpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1626859602, 1626859602219518540),
		Data:     []byte(_Assetsa3993fef2ee05af4e4591868f81e4e0c18c0a4f6),
	}}, "")
