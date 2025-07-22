package handlers

import (
	"net/http"
	"html/template"
	"strconv"
	"example.com/hello-app/data"
)

func MyJokesHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
	   http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	userID, _ := strconv.Atoi(cookie.Value)
	username := data.GetUsernameByID(userID)
	jokes, err := data.GetJokesByUserID(userID)
	if err != nil {
		http.Error(w, "Failed to load your jokes", http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/my_jokes.html"))

	viewData := struct {
		Username string
		Jokes []data.Joke
}{
		Username: username,
		Jokes: jokes,
}
		tmpl.Execute(w, viewData)
}


