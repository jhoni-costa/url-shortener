package main

import (
	"log"
	"net/http"

	"github.com/jhoni-costa/url-shortener/handlers"
	"github.com/jhoni-costa/url-shortener/storage"
)

func main() {

	// 1. Start the memory storage (our "DB")
	db := storage.NewMapStorage()

	// 2. Start the handler, passing the DB to it
	urlHandler := &handlers.URLHandler{Storage: db}

	// 3. Define the routes (endpoints)
	http.HandleFunc("/shorten", urlHandler.Shorten)
	http.HandleFunc("/", urlHandler.Redirect)

	log.Println("Server running on http://localhost:8080")

	// 4. Start server and listen to requests
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
