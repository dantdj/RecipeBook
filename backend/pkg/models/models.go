package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Recipe struct {
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Method      string `json:"method"`
}

type AddRecipeResponse struct {
	Id string `json:"id"`
}

type AddRecipeRequest struct {
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Method      string `json:"method"`
}
