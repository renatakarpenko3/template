package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Artic struct {
	Title   string
	Content string
}

func main() {

	fmt.Println("hello world")

	handler1 := func(w http.ResponseWriter, r *http.Request) {
		tmp1 := template.Must(template.ParseFiles("template/index.html"))
		articles := map[string][]Artic{
			"Articles": {
				{Title: "Article 1", Content: "Content of Article 1"},
				{Title: "Article 2", Content: "Content of Article 2"},
				{Title: "Article 3", Content: "Content of Article 3"},
			},
		}
		tmp1.Execute(w, articles)
	}

	handler2 := func(w http.ResponseWriter, r *http.Request) {
		title := r.PostFormValue("title")
		content := r.PostFormValue("content")
		tmp1 := template.Must(template.ParseFiles("template/index.html"))
		tmp1.ExecuteTemplate(w, "article-list-element", Artic{Title: title, Content: content})
	}

	http.HandleFunc("/", handler1)
	http.HandleFunc("/add-article/", handler2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
