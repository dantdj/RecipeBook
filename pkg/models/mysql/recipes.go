package mysql

import (
	"database/sql"

	"github.com/dantdj/RecipeBook/pkg/models"
)

type RecipeModel struct {
	DB *sql.DB
}

