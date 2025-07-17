package handlers

import (
	"html/template"
	"net/http"
	"time"
	"math/rand"
	"example.com/hello-app/data"
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

rand.Seed(time.Now().UnixNano())
joke := data.Jokes[rand.Intn(len(data.Jokes))]

data := JokePageData{
	Name: name,
	Setup: joke.Setup,
	Punchline: joke.Punchline,
}

tmpl := template.Must(template.ParseFiles("templates/joke.html"))
	tmpl.Execute(w, data)
}
