package main

import (
	"log"
	"net/http"
	"example.com/hello-app/handlers"
	"example.com/hello-app/data"
)

// main is the entry point that registers both endpoints
func main() {
	err := data.InitDB()
	if err != nil {
	log.Fatal("Failed to connect to database:", err)
}

    http.HandleFunc("/api/hello", handlers.HelloHandler)
    http.HandleFunc("/api/send", handlers.SendHandler)
    http.HandleFunc("/api/goodbye", handlers.GoodbyeHandler)
    http.HandleFunc("/api/status", handlers.StatusHandler)
    http.Handle("/", http.FileServer(http.Dir("./html")))
    http.HandleFunc("/api/hello/", handlers.HelloPathHandler) 
    http.HandleFunc("/api/joke", handlers.JokeHandler)
    http.HandleFunc("/api/joke/", handlers.JokeByIDHandler)
    http.HandleFunc("/api/jokes", handlers.AllJokesHandler)
    http.HandleFunc("/joke", handlers.JokeHTMLHandler)
    http.HandleFunc("/submit", handlers.SubmitFormHandler)
    http.HandleFunc("/joke/last", handlers.LastJokeHandler)
    http.HandleFunc("/api/joke/like", handlers.LikeJokeHandler)
    http.HandleFunc("/api/joke/dislike", handlers.DislikeJokeHandler)
    http.HandleFunc("/jokes", handlers.AllJokesPageHandler) 
    http.ListenAndServe(":8080", nil)
}

