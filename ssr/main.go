package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

const PORT = 8000

type User struct {
	Name  string
	Email string
}

var gopher = User{"Gopher", "gopher@google.com"}

func main() {
	out := exec.Command("ls").String()
	fmt.Print(out)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/user", userHandlers)

	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	handleError(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.tmpl")
	handleError(err)

	err = t.Execute(w, gopher)
	handleError(err)
}

func userHandlers(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/user.tmpl")
	handleError(err)
	err = t.Execute(w, gopher)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
