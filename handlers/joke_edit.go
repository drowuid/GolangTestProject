package handlers

import (
	"net/http"
	"html/template"
	"strconv"
	"example.com/hello-app/data"
)

func ShowEditFormHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	joke, err := data.GetJokeByID(id)
	if err != nil {
		http.Error(w, "Joke not found", http.StatusNotFound)
	return
	}
	tmpl := template.Must(template.ParseFiles("templates/edit_joke.html"))
	tmpl.Execute(w, joke)
}

func HandleEditSubmission(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))
	setup := r.FormValue("setup")
	punchline := r.FormValue("punchline")
	category := r.FormValue("category")

	err := data.UpdateJoke(id, setup, punchline, category)
	if err != nil {
		http.Error(w, "Failed to update joke", http.StatusInternalServerError)
	return
	}
	http.Redirect(w, r, "/jokes", http.StatusSeeOther)
}

func HandleDeleteJoke(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, _ := strconv.Atoi(r.FormValue("id"))

	err := data.DeleteJoke(id)
	if err != nil {
		http.Error(w, "Failed to delete joke", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/jokes", http.StatusSeeOther)
}

func EditJokeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ShowEditFormHandler(w, r)
}	else if r.Method == http.MethodPost {
	HandleEditSubmission(w, r)
} else {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
