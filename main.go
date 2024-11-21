package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve files from the "Static" directory, no automatic directory index handling.
	fs := http.StripPrefix("/", http.FileServer(http.Dir("./Static")))
	http.Handle("/", fs)

	// Start the server on port 8080 (use ":80" for production)
	log.Println("Server starting on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
