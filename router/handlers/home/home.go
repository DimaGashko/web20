package home

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"web20.tk/entries"
	"web20.tk/router/common"

	_ "github.com/lib/pq"
)

func Home(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	//iDb := common.AppConfig.Db
	psqlInfo := fmt.Sprintf("host=localhost port=5432 user=postgres password=14159265 dbname=web20 sslmode=disable")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return "", err
	}
	defer db.Close()

	// results, err := db.Query(`
	// 	SELECT a.*, c.name category, c.slug category_slug, ta.tag
	// 	FROM article a
	// 		JOIN category c on a.category_id = c.id
	// 		JOIN article_tag ta on ta.article_id = a.id;
	// `)
	results, err := db.Query("SELECT id title from article")
	if err != nil {
		return "", err
	}
	defer results.Close()

	for results.Next() {
		var a entries.Article
		err = results.Scan(&a.Id, &a.Title)
		if err != nil {
			return "", err
		}

		fmt.Print(a.Title)
	}

	file, err := os.Open("md.md")
	if err != nil {
		return "", err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	html := parseMd(b)
	context["out"] = template.HTML(html)

	context["home-data"] = "something"
	context["app-config"] = common.AppConfig
	return "frontend/dist/templates/home.tmpl", nil
}

func parseMd(input []byte) []byte {
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code")
	return p.SanitizeBytes(blackfriday.Run(input))
}
