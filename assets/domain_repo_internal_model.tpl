{% autoescape off %}
package model

import (
	"github.com/jinzhu/copier"
	"github.com/quexer/utee"

	"baibao/m2/pkg/domain"
)

type {{ content.Name }} struct {
{% for field in content.Fields %}
{{ field.Name }}    {{ field.Type }}    `{{ field.GormTag }}`  {% endfor %}
}

func ({{ content.Name }}) ModelToDomain(x *{{ content.Name }}) *domain.{{ content.Name }} {
	return x.ToDomain()
}

func ({{ content.Name }}) New(x *domain.{{ content.Name }}) *{{ content.Name }} {
	out := &{{ content.Name }}{}
	utee.Chk(copier.Copy(out, x))
	return out
}

func (p *{{ content.Name }}) ToDomain() *domain.{{ content.Name }} {
	out := &domain.{{ content.Name }}{}
	utee.Chk(copier.Copy(out, p))
	return out
}


{% endautoescape %}