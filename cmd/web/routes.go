package main 
package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Status
	mux.HandleFunc("/ping", app.ping)

	// Recipe operations
	mux.HandleFunc("/recipe", app.showRecipe)
	return mux
}