package handlers

import (

"encoding/json"
"net/http"
"strings"
"example.com/hello-app/models"
)

func HelloPathHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

parts := strings.Split(r.URL.Path, "/")
var name string
if len(parts) >= 4 && parts[3] != "" {
name = parts[3]
} else{
name= "stranger"
}
response := models.Message{
Greeting: "Ola",
Name: name,
}

json.NewEncoder(w).Encode(response)
}
