package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cedrickchee/snippetbox/pkg/models"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Because Pat matches the '/' path exactly, we can now remove the manual check
	// of r.URL.Path != '/' from this handler.

	// if r.URL.Path != "/" {
	// 	app.notFound(w) // Use the notFound() helper
	// 	return
	// }

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Create an instance of a templateData struct holding the slice of
	// snippets. Then, use the new render helper.
	app.render(w, r, "home.page.tmpl", &templateData{Snippets: s})
}

// showSnippet is a handler function.
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.
	//
	// Pat doesn't strip the colon from the named capture key, so we need to
	// get the value of ':id' from the query string instead of 'id'.
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w) // Use the notFound() helper.
		return
	}

	// Use the SnippetModel object's Get method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	// Create an instance of a templateData struct holding the snippet data.
	// Then, use the new render helper.
	app.render(w, r, "show.page.tmpl", &templateData{Snippet: s})
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet..."))
}

// createSnippet is a handler function.
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// The check of r.Method != 'POST' is now superfluous and can be removed.

	// if r.Method != "POST" {
	// 	// Use the Header().Set() method to add an 'Allow: POST' header to the
	// 	// response header map. The first parameter is the header name, and
	// 	// the second parameter is the header value.
	// 	w.Header().Set("Allow", "POST")

	// 	// Use the http.Error() function to send a 405 status code and "Method Not
	// 	// Allowed" string as the response body.
	// 	app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
	// 	return
	// }

	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
