package handlers

import (
"encoding/json"
"net/http"
"example.com/hello-app/models"
)

func GoodbyeHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

name := r.URL.Query().Get("name")
if name == "" {
name = "stranger"
}

response := models.Message{
Greeting: "Adeus",
Name: name,
}

json.NewEncoder(w).Encode(response)
}
