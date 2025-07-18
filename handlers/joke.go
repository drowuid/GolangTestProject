package handlers

import (

"encoding/json"
"math/rand"
"net/http"
"strconv"
"strings"
"time"
"example.com/hello-app/data"
)

func JokeHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

if r.Method == http.MethodPost {
var newJoke data.Joke
err := json.NewDecoder(r.Body).Decode(&newJoke)
if err != nil {
http.Error(w, "Invalid JSON", http.StatusBadRequest)
return
}

newJoke.ID = len(data.Jokes) + 1
data.Jokes = append(data.Jokes, newJoke)
json.NewEncoder(w).Encode(newJoke)
return
}

rand.Seed(time.Now().UnixNano())
joke := data.Jokes[rand.Intn(len(data.Jokes))]
json.NewEncoder(w).Encode(joke)
}

func AllJokesHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

category := r.URL.Query().Get("category")
if category == "" {
json.NewEncoder(w).Encode(data.Jokes)
return
}
var filtered []data.Joke
for _, joke := range data.Jokes {
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

for _, joke := range data.Jokes {
if joke.ID == id {
json.NewEncoder(w).Encode(joke)
return
}
}

http.Error(w, "Joke not found", http.StatusNotFound)
}

func LikeJokeHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid joke ID", http.StatusBadRequest)
		return
}

	for i := range data.Jokes {
		if data.Jokes[i].ID == id {
			data.Jokes[i].Likes++
			json.NewEncoder(w).Encode(data.Jokes[i])
			return
		}
	}
	http.Error(w, "Joke not found", http.StatusNotFound)
}

func DislikeJokeHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid joke ID", http.StatusBadRequest)
		return
}

	for i := range data.Jokes {
		if data.Jokes[i].ID == id {
		data.Jokes[i].Dislikes++
		json.NewEncoder(w).Encode(data.Jokes[i])
		return
		}
	}
	http.Error(w, "Joke not found", http.StatusNotFound)
}


