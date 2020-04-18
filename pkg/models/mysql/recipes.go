package mysql

import (
	"database/sql"
	"errors"

	"github.com/dantdj/RecipeBook/pkg/models"
)

type RecipeModel struct {
	DB *sql.DB
}

func (m *RecipeModel) Get(id int) (*models.Recipe, error) {
	statement := `SELECT id, name, ingredients, method, created FROM recipes WHERE id = ?`
	recipe := &models.Recipe{}

	err := m.DB.QueryRow(statement, id).Scan(&recipe.ID, &recipe.Name, &recipe.Ingredients, &recipe.Method, &recipe.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return recipe, nil
}
