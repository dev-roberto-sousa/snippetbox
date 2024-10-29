package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// home is an HTTP handler function that responds with a simple
// "It's alive!" message to indicate that the server is running.
// It writes the message as a plain text response to the client.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("It's alive!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new Snippet"))
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
