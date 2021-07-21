{% autoescape off %}

package impl

type {{ content.Name }}Impl struct {
    {{ content.Name }}Repo repo.{{ content.Name }}Repo
}

{% endautoescape %}