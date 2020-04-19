package main

import (
	"encoding/json"
	"errors"
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

	err := request.ParseForm()
	if err != nil {
		app.clientError(writer, http.StatusBadRequest)
		return
	}

	name := request.PostForm.Get("name")
	ingredients := request.PostForm.Get("ingredients")
	method := request.PostForm.Get("method")

	ok, errors := app.validateRecipeInput(name, ingredients, method)
	if !ok {
		fmt.Fprint(writer, errors)
		return
	}

	id, err := app.recipes.Insert(name, ingredients, method)
	if err != nil {
		app.serverError(writer, err)
		return
	}

	http.Redirect(writer, request, fmt.Sprintf("/recipe?id=%d", id), http.StatusSeeOther)
}
