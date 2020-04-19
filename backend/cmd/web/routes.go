package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(secureHeaders)
	mux := http.NewServeMux()

	// Status
	mux.HandleFunc("/ping", app.ping)

	// Recipe operations
	mux.HandleFunc("/recipe", app.showRecipe)
	mux.HandleFunc("/recipe/add", app.addRecipe)

	return standardMiddleware.Then(mux)
}
