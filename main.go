package main

import (
    "net/http"
"example.com/hello-app/handlers"
)

// main is the entry point that registers both endpoints
func main() {
    http.HandleFunc("/api/hello", handlers.HelloHandler)
    http.HandleFunc("/api/send", handlers.SendHandler)
    http.HandleFunc("/api/goodbye", handlers.GoodbyeHandler)
    http.HandleFunc("/api/status", handlers.StatusHandler)
    http.Handle("/", http.FileServer(http.Dir("./html")))
    http.HandleFunc("/api/hello/", handlers.HelloPathHandler) 
    http.HandleFunc("/api/joke", handlers.JokeHandler)
    http.HandleFunc("/api/joke/", handlers.JokeByIDHandler)
    http.HandleFunc("/api/jokes", handlers.AllJokesHandler)
 
    http.ListenAndServe(":8080", nil)
}

