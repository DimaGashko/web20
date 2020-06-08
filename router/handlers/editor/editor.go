package editor

import (
	"net/http"

	"github.com/gorilla/mux"
	"web20.tk/core/db"
	"web20.tk/entries"
	"web20.tk/router/handlers/common"
	"web20.tk/templates"
)

func New(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	return templates.PATH + "editor.tmpl", nil
}

func Edit(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	slug := mux.Vars(r)["slug"]

	conn := db.Get()
	res := conn.Where(`slug = ?`, slug).First(&entries.Post{})

	if res.RowsAffected == 0 {
		return "", common.New404()
	}

	context["post"] = res.Value

	return templates.PATH + "editor.tmpl", nil
}
