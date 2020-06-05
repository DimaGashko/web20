package editor

import (
	"net/http"
)

func AddArticle(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	return "frontend/dist/templates/add-article.tmpl", nil
}
