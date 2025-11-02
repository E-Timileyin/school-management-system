package main

import (
	"fmt"
	"log"
	"net/http"
)

// Simple HTTP Server Example
func main() {
	// Define a handler function for the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! Welcome to my first Go server!")
	})

	// Define a handler for the "/about" path
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a simple Go web server!")
	})

	// Start the server on port 8080
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)

	// Start the server and log any errors
	log.Fatal(http.ListenAndServe(port, nil))
}
