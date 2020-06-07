package home

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/gosimple/slug"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"web20.tk/core/db"
	"web20.tk/entries"
	"web20.tk/templates"

	_ "github.com/lib/pq"
)

func Home(w http.ResponseWriter, r *http.Request, context map[string]interface{}) (string, error) {
	conn := db.Get()

	now := fmt.Sprint(time.Now().Unix())

	article := entries.Article{
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
		},
	}

	conn.First(&article)
	context["art"] = article

	// results, err := db.Query("SELECT id, slug, title, description, content, author, timestamp from articles")
	// if err != nil {
	// 	return "", err
	// }
	// defer results.Close()

	// articles := []entries.Article{}

	// for results.Next() {
	// 	var a entries.Article
	// 	err = results.Scan(&a.Id, &a.Slug, &a.Title, &a.Description, &a.Content, &a.Author, &a.Timestamp)
	// 	if err != nil {
	// 		return "", err
	// 	}

	// 	articles = append(articles, a)
	// }

	// context["articles"] = articles

	// file, err := os.Open("md.md")
	// if err != nil {
	// 	return "", err
	// }
	// defer file.Close()

	// b, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	return "", err
	// }

	// html := parseMd(b)
	// context["out"] = template.HTML(html)

	// context["home-data"] = "something"
	// context["app-config"] = common.AppConfig

	// 	article := entries.Article{
	// 	Slug:        now + "-" + slug.Make("SHA1 Collision"),
	// 	Title:       "SHA1 Collision",
	// 	Description: "When I was doing the DEF CON CTF Qualifier last weekend, I came across an interesting question where you need to create two pdf files with the same SHA1 hash.",
	// 	Content:     md,
	// 	Author:      "Alan Chang (https://github.com/tcode2k16)",
	// 	Category: entries.Category{
	// 		Slug: "backend",
	// 		Name: "backend",
	// 	},
	// 	Tags: []entries.Tag{
	// 		{Name: "security"},
	// 		{Name: "ssh1"},
	// 	},
	// }

	return templates.PAGES_PATH + "home/home.tmpl", nil
}

func parseMd(input []byte) []byte {
	p := bluemonday.UGCPolicy()
	p.AllowAttrs("class").Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code")
	return p.SanitizeBytes(blackfriday.Run(input))
}

const md = `---
title: "SHA1 Collision"
date: 2018-05-16T11:32:16+08:00
draft: false
tags: [
  "ctf",
  "cyber-security",
  "hash"
]
description: Creating a SHA1 collision with PDF files
---

### Introduction

When I was doing the [DEF CON CTF Qualifier](https://www.oooverflow.io/) last weekend, I came across an interesting question where you need to create two pdf files with the same SHA1 hash.

### Research

I know SHA1 hash was already broken when google blogged about [creating the first SHA1 collision](https://security.googleblog.com/2017/02/announcing-first-sha1-collision.html), but I was not sure that I can reproduce the process with limited hardware.

### Result

In the end, I came across [this website](http://alf.nu/SHA1) that is able to generate two PDF files with the same SHA1 hash using two JPG images based on [this paper](https://shattered.io/). This helps demonstrate how SHA1 is no longer secure and developers should start using other hashing algori`
