package home

import (
	"net/http"

	"web20.tk/core/db"
	"web20.tk/entries"
	"web20.tk/templates"

	_ "github.com/lib/pq"
)

func Home(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	conn := db.Get()
	posts := conn.Limit(10).Find(&[]entries.Post{})

	context["posts"] = posts.Value

	return templates.PATH + "home.tmpl", nil
}
