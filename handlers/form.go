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
        category := r.FormValue("category") // 🛠️ corrected typo "categgory" ➡ "category"

        err := data.InsertJoke(setup, punchline, category)
        if err != nil {
            http.Error(w, "Failed to insert joke", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/joke/last?name=You", http.StatusSeeOther)
    }
}

