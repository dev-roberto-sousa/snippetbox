package main

import (
	"log"
	"net/http"
)

// home is an HTTP handler function that responds with a simple
// "It's alive!" message to indicate that the server is running.
// It writes the message as a plain text response to the client.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's alive!"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Starting on server localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
