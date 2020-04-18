package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) ping(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "OK!")
}

func (app *application) showRecipe(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.serverError(writer, err)
		return
	}
	fmt.Fprintf(writer, "Will display recipe with ID %d in future...", id)
}

