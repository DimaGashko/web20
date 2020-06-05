package editor

import (
	"net/http"

	"web20.tk/router/handlers/common"
)

func AddArticle(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	return common.PAGES_PATH + "add-article/add-article.tmpl", nil
}
