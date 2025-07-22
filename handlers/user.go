package handlers

import (
	"html/template"
	"net/http"
	"example.com/hello-app/data"
	"strconv"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
	tmpl := template.Must(template.ParseFiles("templates/signup.html"))
	tmpl.Execute(w, nil)
	return
	}

	//POST
	username := r.FormValue("username")
	password := r.FormValue("password")

	err := data.RegisterUser(username, password)
	if err != nil {
		http.Error(w, "Username already exists or invalid", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil)
	return
	}

	//POST
	username := r.FormValue("username")
	password := r.FormValue("password")

	userID, err := data.AuthenticateUser(username, password)
	if err != nil {
		http.Error(w, "Login failed", http.StatusUnauthorized)
		return
	}

	cookie := &http.Cookie{
		Name: "user_id",
		Value: strconv.Itoa(userID),
		Path: "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/jokes", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	expired := &http.Cookie{
		Name: "user_id",
		Value: "",
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(w, expired)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
