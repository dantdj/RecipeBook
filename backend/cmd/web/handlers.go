package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dantdj/RecipeBook/pkg/models"
)

func (app *application) ping(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "OK!")
}

func (app *application) showRecipe(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")

	recipe, err := app.recipes.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(writer)
		} else {
			app.serverError(writer, err)
		}

		return
	}
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	recipeJson, _ := json.Marshal(recipe)
	fmt.Fprintf(writer, string(recipeJson))
}

func (app *application) addRecipe(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		http.Error(writer, "Method Not Allowed", 405)
		return
	}

	writer.Header().Set("Access-Control-Allow-Origin", "*")

	var recipe models.AddRecipeRequest

	err := json.NewDecoder(request.Body).Decode(&recipe)
	if err != nil {
		app.clientError(writer, 400)
		return
	}

	ok, errors := app.validateRecipeInput(recipe.Name, recipe.Ingredients, recipe.Method)
	if !ok {
		fmt.Fprint(writer, errors)
		return
	}

	id, err := app.recipes.Insert("default", recipe.Name, recipe.Ingredients, recipe.Method)
	if err != nil {
		app.serverError(writer, err)
		return
	}
	returnObject := models.AddRecipeResponse{id}
	idJson, _ := json.Marshal(returnObject)

	fmt.Fprintf(writer, string(idJson))
}
