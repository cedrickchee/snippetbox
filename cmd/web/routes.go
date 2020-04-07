package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our application receives.
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()
	// Important to note that Pat matches patterns in the order that they are
	// registered.
	mux.Get("/", http.HandlerFunc(app.home)) // the pattern '/' is a special case.
	// To ensure that the exact match takes preference, we need to register the
	// exact match routes before any wildcard routes.
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	// Wildcard routes.
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	// Create a file server which serves files out of the './ui/static' directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// Return the 'standard' middleware chain followed by the servemux.
	return standardMiddleware.Then(mux)
}
