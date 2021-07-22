{% autoescape off %}

func (p *{{ content.Name }}) Valid() error {
	if err := validator.New().Struct(p); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

type {{ content.Name }}List []*{{ content.Name }}

func (p {{ content.Name }}List) FindById(id int) *{{ content.Name }} {
	for _, v := range p {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func (p {{ content.Name }}List) ListById(id int) {{ content.Name }}List {
	fn := func(x *{{ content.Name }}) bool {
		return x.ID == id
	}

	var l {{ content.Name }}List
	linq.From(p).WhereT(fn).ToSlice(&l)
	return l
}

{% endautoescape %}