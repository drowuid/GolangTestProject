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

pageData := JokePageData{
        Name: name,
        Setup: joke.Setup,
        Punchline: joke.Punchline,
}

tmpl := template.Must(template.ParseFiles("templates/joke.html"))
        tmpl.Execute(w, pageData)
}


func LastJokeHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "stranger"
    }

if len (data.Jokes) == 0 {
	http.Error(w, "No jokes available", http.StatusNotFound)
	return
}

joke := data.Jokes[len(data.Jokes)-1]

    pageData := JokePageData{
        Name:      name,
        Setup:     joke.Setup,
        Punchline: joke.Punchline,
    }

    tmpl := template.Must(template.ParseFiles("templates/joke.html"))
    tmpl.Execute(w, pageData)
}
