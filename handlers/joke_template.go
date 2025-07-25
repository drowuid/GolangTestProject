package handlers

import (
    "html/template"
    "net/http"
    "time"
    "math/rand"
    "example.com/hello-app/data"
    "strings"
)

type JokePageData struct {
    Name      string
    Setup     string
    Punchline string
}

func JokeHTMLHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "stranger"
    }

    jokes, err := data.GetAllJokes()
    if err != nil || len(jokes) == 0 {
        http.Error(w, "No jokes available", http.StatusNotFound)
        return
    }

    rand.Seed(time.Now().UnixNano())
    joke := jokes[rand.Intn(len(jokes))]

    pageData := JokePageData{
        Name:      name,
        Setup:     joke.Setup,
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

    jokes, err := data.GetAllJokes()
    if err != nil || len(jokes) == 0 {
        http.Error(w, "No jokes available", http.StatusNotFound)
        return
    }

    joke := jokes[len(jokes)-1]

    pageData := JokePageData{
        Name:      name,
        Setup:     joke.Setup,
        Punchline: joke.Punchline,
    }

    tmpl := template.Must(template.ParseFiles("templates/joke.html"))
    tmpl.Execute(w, pageData)
}

func AllJokesPageHandler(w http.ResponseWriter, r *http.Request) {
    category := r.URL.Query().Get("category")
    keyword := r.URL.Query().Get("search")

    jokes, err := data.GetAllJokes()
    if err != nil {
        http.Error(w, "Error loading jokes", http.StatusInternalServerError)
        return
    }

    var filtered []data.Joke
    for _, joke := range jokes {
        if category != "" && !strings.EqualFold(joke.Category, category) {
            continue
        }
        if keyword != "" && !strings.Contains(strings.ToLower(joke.Setup+joke.Punchline), strings.ToLower(keyword)) {
            continue
        }
        filtered = append(filtered, joke)
    }

    tmpl := template.Must(template.ParseFiles("templates/all_jokes.html"))
    tmpl.Execute(w, filtered)
}
