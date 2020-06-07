package entries

import (
	"html/template"
	"regexp"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

type Post struct {
	gorm.Model
	Slug        string    `gorm:"slug;unique_index"`
	Title       string    `gorm:"title"`
	Description string    `gorm:"description"`
	Content     Md        `gorm:"content"`
	Author      string    `gorm:"author"`
	Timestamp   time.Time `gorm:"timestamp"`
	Listed      bool      `gorm:"listed"`
	Category    Category  `gorm:"category;foreignkey:Category"`
	Tags        []Tag
}

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
