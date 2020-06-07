package editor

import (
	"net/http"

	"web20.tk/templates"
)

func Editor(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	return templates.PAGES_PATH + "editor/editor.tmpl", nil
}
