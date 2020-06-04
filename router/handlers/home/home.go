package home

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"web20.tk/router/common"
)

func Home(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
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
