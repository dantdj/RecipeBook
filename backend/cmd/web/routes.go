package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.logRequest, secureHeaders)
	mux := http.NewServeMux()

	// Status
	mux.HandleFunc("/ping", app.ping)

	// Recipe operations
	mux.HandleFunc("/recipe", app.showRecipe)
	mux.HandleFunc("/recipe/add", app.addRecipe)

	// Image operations
	mux.HandleFunc("/image", app.serveImage)
	mux.HandleFunc("/image/resize", app.resizeImage)

	return standardMiddleware.Then(mux)
}
