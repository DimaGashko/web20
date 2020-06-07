package posts

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gosimple/slug"
	"web20.tk/entries"
	"web20.tk/templates"

	_ "github.com/lib/pq"
)

func List(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	return templates.PATH + "posts-list.tmpl", nil
}

func Post(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	now := fmt.Sprint(time.Now().Unix())

	post := entries.Post{
		Slug:        now + "-" + slug.Make("SHA1 Collision"),
		Title:       "SHA1 Collision",
		Description: "When I was doing the DEF CON CTF Qualifier last weekend, I came across an interesting question where you need to create two pdf files with the same SHA1 hash.",
		Content:     md,
		Author:      "Alan Chang (https://github.com/tcode2k16)",
		Category: entries.Category{
			Slug: "backend",
			Name: "backend",
		},
		Tags: []entries.Tag{
			{Name: "security"},
			{Name: "ssh1"},
			{Name: "ctf"},
			{Name: "cyber-security"},
			{Name: "hash"},
		},
		Timestamp: time.Now(),
	}

	context["post"] = post

	return templates.PATH + "post.tmpl", nil
}

const md = `
### Research

I know SHA1 hash was already broken when google blogged about [creating the first SHA1 collision](https://security.googleblog.com/2017/02/announcing-first-sha1-collision.html), but I was not sure that I can reproduce the process with limited hardware.

### Result

In the end, I came across [this website](http://alf.nu/SHA1) that is able to generate two PDF files with the same SHA1 hash using two JPG images based on [this paper](https://shattered.io/). This helps demonstrate how SHA1 is no longer secure and developers should start using other hashing algori`
