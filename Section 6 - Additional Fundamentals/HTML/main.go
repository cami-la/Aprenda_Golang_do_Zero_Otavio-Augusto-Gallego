package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
	Age   int
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := user{
			"Camila",
			"camila@email.com",
			31,
		}
		err := templates.ExecuteTemplate(w, "home.html", u)
		if err != nil {
			return
		}
	})

	fmt.Println("Listening on port 5000")
	log.Fatalln(http.ListenAndServe(":5000", nil))
}
