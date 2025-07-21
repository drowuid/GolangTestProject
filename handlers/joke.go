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

        err = data.InsertJoke(newJoke.Setup, newJoke.Punchline, newJoke.Category)
        if err != nil {
            http.Error(w, "Failed to insert joke", http.StatusInternalServerError)
            return
        }

        // Get updated jokes and return the last one
        jokes, _ := data.GetAllJokes()
        json.NewEncoder(w).Encode(jokes[len(jokes)-1])
        return
    }

    jokes, err := data.GetAllJokes()
    if err != nil || len(jokes) == 0 {
        http.Error(w, "No jokes available", http.StatusNotFound)
        return
    }

    rand.Seed(time.Now().UnixNano())
    joke := jokes[rand.Intn(len(jokes))]
    json.NewEncoder(w).Encode(joke)
}

func AllJokesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    jokes, err := data.GetAllJokes()
    if err != nil {
        http.Error(w, "Failed to load jokes", http.StatusInternalServerError)
        return
    }

    category := r.URL.Query().Get("category")
    if category == "" {
        json.NewEncoder(w).Encode(jokes)
        return
    }

    var filtered []data.Joke
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

    joke, err := data.GetJokeByID(id)
    if err != nil {
        http.Error(w, "Joke not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(joke)
}

func LikeJokeHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid joke ID", http.StatusBadRequest)
        return
    }

    joke, err := data.GetJokeByID(id)
    if err != nil {
        http.Error(w, "Joke not found", http.StatusNotFound)
        return
    }

    joke.Likes++
    err = data.UpdateJoke(id, joke.Setup, joke.Punchline, joke.Category)
    if err != nil {
        http.Error(w, "Failed to update likes", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(joke)
}

func DislikeJokeHandler(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid joke ID", http.StatusBadRequest)
        return
    }

    joke, err := data.GetJokeByID(id)
    if err != nil {
        http.Error(w, "Joke not found", http.StatusNotFound)
        return
    }

    joke.Dislikes++
    err = data.UpdateJoke(id, joke.Setup, joke.Punchline, joke.Category)
    if err != nil {
        http.Error(w, "Failed to update dislikes", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(joke)
}
