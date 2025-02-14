package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", nil)
	})

	fmt.Println("Listening on port 5000")
	log.Fatalln(http.ListenAndServe(":5000", nil))
}
