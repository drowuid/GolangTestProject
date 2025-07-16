package handlers

import (

"encoding/json"
"math/rand"
"net/http"
"strconv"
"strings"
"time"

)

type Joke struct {
ID int `json:"id"`
Setup string `json:"setup"`
Punchline string `json:"punchline"`
Category string `json:"category"`
}



var jokes = []Joke{

        {ID: 1, Setup: "Why don't scientists trust atoms?", Punchline: "Because they make up everything!", Category: "science"},
        {ID: 2, Setup: "How does a penguin build its house?", Punchline: "Igloos it together.", Category: "animals"},
        {ID: 3, Setup: "Why did the computer go to therapy?", Punchline: "It had too many processing issues.", Category: "tech"},
        {ID: 4, Setup: "What do you call fake spaghetti?", Punchline: "An impasta.", Category: "food"},
        {ID: 5, Setup: "Why was the math book sad?", Punchline: "It had too many problems.", Category:"math"},
}


func JokeHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

if r.Method == http.MethodPost {
var newJoke Joke
err := json.NewDecoder(r.Body).Decode(&newJoke)
if err != nil {
http.Error(w, "Invalid JSON", http.StatusBadRequest)
return
}

newJoke.ID = len(jokes) + 1
jokes = append(jokes, newJoke)
json.NewEncoder(w).Encode(newJoke)
return
}

rand.Seed(time.Now().UnixNano())
joke := jokes[rand.Intn(len(jokes))]
json.NewEncoder(w).Encode(joke)
}

func AllJokesHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

category := r.URL.Query().Get("category")
if category == "" {
json.NewEncoder(w).Encode(jokes)
return
}
var filtered []Joke
for _, joke := range jokes {
if strings.EqualFold(joke.Category, category) {
filtered = append(filtered, joke)
}
}
if len(filtered) == 0 {
http.Error(w, "No jokes found for category", http.StatusNotFound)
return
}


json.NewEncoder(w).Encode(filtered)
}

func JokeByIDHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

parts := strings.Split(r.URL.Path, "/")
if len(parts) < 4 {
http.Error(w, "Missing joke ID", http.StatusBadRequest)
return
}

idStr := parts[3]
id, err := strconv.Atoi(idStr)
if err != nil {
http.Error(w, "Invalid joke ID", http.StatusBadRequest)
return
}

for _, joke := range jokes {
if joke.ID == id {
json.NewEncoder(w).Encode(joke)
return
}
}

http.Error(w, "Joke not found", http.StatusNotFound)
}
