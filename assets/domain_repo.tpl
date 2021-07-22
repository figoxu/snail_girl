{% autoescape off %}
package repo
import (
	"baibao/meishi/pkg/next/domain"
	"context"
)
//go:generate mockgen -destination=../mocks/mrepo/{{ fileName }} -package=mrepo . {{ content.Name }}Repo

type {{ content.Name }}Repo interface {
    MustGet(ctx context.Context, id int) (*domain.{{ content.Name }}, error)
    MultiGet(ctx context.Context, id ...int) (domain.{{ content.Name }}List, error)
    Create(ctx context.Context, in *domain.{{ content.Name }}) error
    Update(ctx context.Context, in *domain.{{ content.Name }}) error
    List(ctx context.Context) (domain.{{ content.Name }}List, error)
}

{% endautoescape %}