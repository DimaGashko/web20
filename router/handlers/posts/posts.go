package posts

import (
	"net/http"

	"web20.tk/core/db"
	"web20.tk/entries"
	"web20.tk/templates"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func List(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	conn := db.Get()
	posts := conn.Where("listed = false").Find(&[]entries.Post{})

	context["posts"] = posts.Value
	return templates.PATH + "posts-list.tmpl", nil
}

func Post(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	slug := mux.Vars(r)["slug"]

	conn := db.Get()
	post := conn.Where(`slug = ?`, slug).First(&entries.Post{})

	context["post"] = post.Value

	return templates.PATH + "post.tmpl", nil
}
