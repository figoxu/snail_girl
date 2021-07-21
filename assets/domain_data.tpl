{% autoescape off %}
package testdata

import (
	"github.com/quexer/utee"
)

func {{ content.Name }} () *domain.{{ content.Name }} {
	obj := &domain.{{ content.Name }} {
        {% for field in content.Fields %}
            {{ field.Name }}  : {{ field.MockVal }} , {% endfor %}
	}
	err := obj.Valid()
	utee.Chk(err)
	return obj
}

{% endautoescape %}