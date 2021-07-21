{% autoescape off %}
package testdata

import (
    "baibao/meishi/pkg/next/domain"
	"github.com/quexer/utee"
)

func {{ content.Name }} (id ...int) *domain.{{ content.Name }} {
	obj := &domain.{{ content.Name }} {
        {% for field in content.Fields %}
            {{ field.Name }}  : {{ field.MockVal }} , {% endfor %}
	}
	err := obj.Valid()
	utee.Chk(err)
	if len(id) > 0 {
        obj.Id = id[0]
    }
	return obj
}

{% endautoescape %}