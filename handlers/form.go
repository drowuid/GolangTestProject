package handlers

import (
    "html/template"
    "net/http"
    "strconv"
    "example.com/hello-app/data"
)

func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	userID, _ := strconv.Atoi(cookie.Value)
	username := data.GetUsernameByID(userID)



    if r.Method == http.MethodGet {
        tmpl := template.Must(template.ParseFiles("templates/submit.html"))
        tmpl.Execute(w, struct {
		Username string
	}{Username: username})
        return
    }
   
	if r.Method == http.MethodPost {
        setup := r.FormValue("setup")
        punchline := r.FormValue("punchline")
        category := r.FormValue("category")

        err = data.InsertJoke(userID, setup, punchline, category)
        if err != nil {
            http.Error(w, "Failed to insert joke", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/joke/last?name="+username, http.StatusSeeOther)
    }
}

