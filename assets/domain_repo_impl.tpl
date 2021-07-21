{% autoescape off %}
package impl

import (
	"baibao/meishi/pkg/next/domain"
	"baibao/meishi/pkg/next/repo/internal/model"
	"context"
    "github.com/ahmetb/go-linq/v3"
    "github.com/pkg/errors"
)

type {{ content.Name }}RepoImpl struct {
	Db *gorm.DB
}

func (p *{{ content.Name }}RepoImpl) Update(ctx context.Context, in *domain.{{ content.Name }}) error {
	if err := in.Valid(); err != nil {
		return err
	}

	err := p.Db.Save(model.{{ content.Name }}{}.New(in)).Error

	return errors.WithStack(err)
}

func (p *{{ content.Name }}RepoImpl) Create(ctx context.Context, in *domain.{{ content.Name }}) error {
	if err := in.Valid(); err != nil {
		return err
	}
	obj := model.{{ content.Name }}{}.New(in)
	if err := p.Db.Create(obj).Error; err != nil {
		return errors.WithStack(err)
	}
	in.ID = obj.ID
	return nil
}

func (p *{{ content.Name }}RepoImpl) MustGet(ctx context.Context, id int) (*domain.{{ content.Name }}, error) {
	var o model.{{ content.Name }}
	if err := p.Db.Take(&o, id).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return o.ToDomain(), nil
}


func (p *{{ content.Name }}RepoImpl) MultiGet(ctx context.Context, id ...int) (domain.{{ content.Name }}List, error) {
	if len(id) == 0 {
		return domain.{{ content.Name }}List{}, nil
	}

	var l []*model.{{ content.Name }}
	if err := p.Db.
		Where("id in ?", id).
		Find(&l).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var out domain.{{ content.Name }}List
	linq.From(l).SelectT(model.{{ content.Name }}{}.ModelToDomain).ToSlice(&out)
	return out, nil
}

func (p *{{ content.Name }}RepoImpl) List(ctx context.Context) (domain.{{ content.Name }}List, error) {
	var l []*model.{{ content.Name }}
	db := p.Db

	if err := db.Find(&l).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var out domain.{{ content.Name }}List
	linq.From(l).SelectT(model.{{ content.Name }}{}.ModelToDomain).ToSlice(&out)
	return out, nil
}

{% endautoescape %}