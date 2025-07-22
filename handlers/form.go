package handlers

import (
    "html/template"
    "net/http"
    "strconv"
    "example.com/hello-app/data"
)

func SubmitFormHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        tmpl := template.Must(template.ParseFiles("templates/submit.html"))
        tmpl.Execute(w, nil)
        return
    }

    if r.Method == http.MethodPost {
        cookie, err := r.Cookie("user_id")
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        userID, _ := strconv.Atoi(cookie.Value)

        setup := r.FormValue("setup")
        punchline := r.FormValue("punchline")
        category := r.FormValue("category")

        err = data.InsertJoke(userID, setup, punchline, category)
        if err != nil {
            http.Error(w, "Failed to insert joke", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/joke/last?name=You", http.StatusSeeOther)
    }
}

