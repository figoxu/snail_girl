{% autoescape off %}
package repo
//go:generate mockgen -destination=../mocks/mrepo/{{ fileName }} -package=mrepo . {{ content.Name }}Repo

type {{ content.Name }}Repo interface {
}

{% endautoescape %}