package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Writes an error message and the stack trace to errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(writer http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Sends a specific status code and its corresponding description
// to the user.
func (app *application) clientError(writer http.ResponseWriter, status int) {
	http.Error(writer, http.StatusText(status), status)
}

// Convenience wrapper around `clientError` which sends 404
// Not Found to the user.
func (app *application) notFound(writer http.ResponseWriter) {
	app.clientError(writer, http.StatusNotFound)
}
