package postsApi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
	"github.com/gosimple/slug"
	"web20.tk/core/db"
	"web20.tk/entries"
	"web20.tk/router/handlers/common"
)

func GetAll(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	conn := db.Get()
	resp := conn.Find(&[]entries.Post{})
	return resp, nil
}

func Get(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	resp := make(map[string]interface{})
	return resp, nil
}

func Create(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	conn := db.Get()

	var post entries.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		return nil, common.WrapError(err, http.StatusBadRequest)
	}

	now := fmt.Sprint((time.Now().Unix() % 1e4) + rand.Int63n(10000))
	post.Slug = now + "-" + slug.Make(post.Title)

	err = processImg(&post)
	if err != nil {
		return "", common.WrapError(err, http.StatusBadRequest)
	}

	if post.Secret == "" {
		post.Secret = "web20"
	}

	conn.Create(&post)

	return post, nil
}

func Update(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	slug := mux.Vars(r)["slug"]
	conn := db.Get()

	var curPost entries.Post
	res := conn.Where(`slug = ?`, slug).First(&curPost)

	if res.RowsAffected == 0 {
		return nil, common.New404()
	}

	var post entries.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		return nil, common.WrapError(err, http.StatusBadRequest)
	}

	if curPost.Secret != post.Secret {
		return nil, common.New403()
	}

	if post.Image != curPost.Image {
		err := processImg(&post)
		if err != nil {
			return "", common.WrapError(err, http.StatusBadRequest)
		}
	}

	conn.Save(&post)
	return post, nil
}

func Delete(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	resp := make(map[string]interface{})
	return resp, nil
}

func processImg(post *entries.Post) error {
	var err error
	if post.Image == "" {
		post.Image = "https://picsum.photos/1200/800"
	}

	post.Image, err = downloadImg(post.Image, post.Slug)
	if err != nil {
		return err
	}

	err = optimizeImg(post.Image)
	if err != nil {
		return err
	}

	return nil
}

func downloadImg(src, name string) (string, error) {
	response, err := http.Get(src)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	path := fmt.Sprintf("/static/posts/images/%s.jpg", name)
	sysPath := "./frontend/dist" + path

	file, err := os.Create(sysPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	return path, nil
}

func optimizeImg(path string) error {
	sysPath := "frontend/dist" + path
	cmd := exec.Command("convert", sysPath, "-quality", "79", sysPath)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
