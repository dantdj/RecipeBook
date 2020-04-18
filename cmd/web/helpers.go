package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"unicode/utf8"
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

func (app *application) validateRecipeInput(name, ingredients, method string) (bool, map[string]string) {
	errors := make(map[string]string)

	if strings.TrimSpace(name) == "" {
		errors["name"] = "This field cannot be blank."
	} else if utf8.RuneCountInString(name) > 200 {
		errors["name"] = "This field is too long (maximum is 200 characters)"
	}

	if strings.TrimSpace(ingredients) == "" {
		errors["ingredients"] = "This field cannot be blank"
	}

	if len(errors) > 0 {
		return false, errors
	}

	return true, nil
}
