package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Recipe struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Ingredients string    `json:"ingredients"`
	Method      string    `json:"method"`
	Created     time.Time `json:"created"`
}
