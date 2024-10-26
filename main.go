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

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux,
	// then register the home, snippetView and snippetCreate functions
	// as the handlers for the "/", "/snippet/view",
	// and "/snippet/create" URL patterns.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting on server localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
