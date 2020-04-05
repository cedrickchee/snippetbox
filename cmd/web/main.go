package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files out of the './ui/static' directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with '/static/'. For matching paths, we strip the
	// '/static' prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
