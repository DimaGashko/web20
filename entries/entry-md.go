package entries

import (
	"html/template"
	"regexp"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

type Md string

func (md Md) Html() template.HTML {
	html := parseMd([]byte(md))
	return template.HTML(html)
}

func parseMd(input []byte) []byte {
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code")
	return p.SanitizeBytes(blackfriday.Run(input))
}
