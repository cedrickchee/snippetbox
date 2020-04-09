package main

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	// We can then check the value of the response status code and body using
	// the same code as before.
	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	// And we can check that the response body written by the ping handler
	// equals 'OK'.
	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}
