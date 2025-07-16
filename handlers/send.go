package handlers

import(
"encoding/json"
"net/http"
"example.com/hello-app/models"
)

func SendHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

if r.Method != http.MethodPost {
http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
return
}

var input models.Message
err := json.NewDecoder(r.Body).Decode(&input)
if err != nil {
http.Error(w, "Invalid JSON", http.StatusBadRequest)
return
}

json.NewEncoder(w).Encode(input)
}
