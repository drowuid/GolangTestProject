package handlers

import (
	"html/template"
	"net/http"
	"example.com/hello-app/data"
)

func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
	tmpl := template.Must(template.ParseFiles("templates/submit.html"))
	tmpl.Execute(w, nil)
return
}

	if r.Method == http.MethodPost {
	setup := r.FormValue("setup")
	punchline := r.FormValue("punchline")
	category := r.FormValue("categgory")
	
	newJoke := data.Joke{
	ID: len(data.Jokes) + 1,
	Setup: setup,
	Punchline: punchline,
	Category: category,
	}
	
	data.Jokes = append(data.Jokes, newJoke)
	http.Redirect(w, r, "/joke/last?name=You", http.StatusSeeOther)
	}
}
