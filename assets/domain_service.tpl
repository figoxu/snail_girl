{% autoescape off %}
package service

//go:generate mockgen -destination=../mocks/mservice/{{ fileName }} -package=mservice . {{ content.Name }}Srv
type {{ content.Name }}Srv interface {

}


{% endautoescape %}