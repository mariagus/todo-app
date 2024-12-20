package main

import (
	"log"
	"net/http"
)

func main() {
	initDB()
	defer DB.Close()
	// Route the handlers for each URL path
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/delete", deleteHandler)

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server on port 8080
}
