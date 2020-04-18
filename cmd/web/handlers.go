package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dantdj/RecipeBook/pkg/models"
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

	recipe, err := app.recipes.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(writer)
		} else {
			app.serverError(writer, err)
		}

		return
	}
	recipeJson, _ := json.Marshal(recipe)
	fmt.Fprintf(writer, string(recipeJson))
}

