{% autoescape off %}
package impl

type {{ content.Name }}RepoImpl struct {
	Ds *database.Ds
}

func (p *{{ content.Name }}RepoImpl) Update(ctx context.Context, in *domain.{{ content.Name }}) error {
	if err := in.Valid(); err != nil {
		return err
	}

	err := p.Ds.Gdb().Save(model.{{ content.Name }}{}.New(in)).Error

	return errors.WithStack(err)
}

func (p *{{ content.Name }}RepoImpl) Create(ctx context.Context, in *domain.{{ content.Name }}) error {
	if err := in.Valid(); err != nil {
		return err
	}
	obj := model.{{ content.Name }}{}.New(in)
	if err := p.Ds.Gdb().Create(obj).Error; err != nil {
		return errors.WithStack(err)
	}
	in.ID = obj.ID
	return nil
}

func (p *{{ content.Name }}RepoImpl) MustGet(ctx context.Context, id int) (*domain.{{ content.Name }}, error) {
	var o model.{{ content.Name }}
	if err := p.Ds.Gdb().Take(&o, id).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return o.ToDomain(), nil
}


func (p *{{ content.Name }}RepoImpl) MultiGet(ctx context.Context, id ...int) (domain.{{ content.Name }}List, error) {
	if len(id) == 0 {
		return domain.{{ content.Name }}List{}, nil
	}

	var l []*model.Order
	if err := p.Ds.Gdb().
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
	db := p.Ds.Gdb()

	if err := db.Find(&l).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	var out domain.{{ content.Name }}List
	linq.From(l).SelectT(model.{{ content.Name }}{}.ModelToDomain).ToSlice(&out)
	return out, nil
}

{% endautoescape %}