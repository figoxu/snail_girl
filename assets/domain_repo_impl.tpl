{% autoescape off %}
package impl

type {{ content.Name }}RepoImpl struct {
	Ds *database.Ds
}

{% endautoescape %}