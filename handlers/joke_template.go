package handlers

import (
	"html/template"
	"net/http"
)

type JokePageData struct {
	Name string
	Setup string
	Punchline string
}

func JokeHTMLHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "stranger"
}

joke := JokePageData{
	Name: name,
	Setup: "Why don't scientists trust atoms?",
	Punchline: "Because they make up everything!",
}

tmpl := template.Must(template.ParseFiles("templates/joke.html"))
	tmpl.Execute(w, joke)
}
