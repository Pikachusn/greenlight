package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthchechHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
	// Create a fixed-format JSON response from a string. Notice how we're using a raw
	// string literal (enclosed with backticks) so that we can include double-quote
	// characters in the JSON without needing to escape them? We also use the %q verb to
	// wrap the inerpolated values in double-quotes.

	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	// Set the "Content-Type: application/json" header on the response. If you forget to
	// this, Go will default ti sending a "Content-Type: text/plain; charset=utf-8"
	// header instead.
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON as then HTTP response body.
	w.Write([]byte(js))
}
