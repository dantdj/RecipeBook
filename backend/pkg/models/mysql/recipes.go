package mysql

import (
	"database/sql"
	"errors"

	"github.com/dantdj/RecipeBook/pkg/models"
)

type RecipeModel struct {
	DB *sql.DB
}

func (m *RecipeModel) Insert(name, ingredients, method string) (int, error) {
	statement := `INSERT INTO recipes (name, ingredients, method, created) VALUES (?, ?, ?, UTC_TIMESTAMP())`

	result, err := m.DB.Exec(statement, name, ingredients, method)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
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
