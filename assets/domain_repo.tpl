{% autoescape off %}
package repo
//go:generate mockgen -destination=../mocks/mrepo/{{ fileName }} -package=mrepo . {{ content.Name }}Repo

type {{ content.Name }}Repo interface {
    MustGet(ctx context.Context, id int) (*domain.{{ content.Name }}, error)
    MultiGet(ctx context.Context, id ...int) (domain.{{ content.Name }}List, error)
    Create(ctx context.Context, in *domain.{{ content.Name }}) error
    Update(ctx context.Context, in *domain.{{ content.Name }}) error
    DelById(ctx context.Context, id int) error
    List(ctx context.Context) (domain.{{ content.Name }}List, error)
}

{% endautoescape %}