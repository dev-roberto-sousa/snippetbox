package main

import (
	"log"
	"net/http"
)

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
