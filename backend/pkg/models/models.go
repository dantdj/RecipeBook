package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Recipe struct {
	ID          int
	Name        string
	Ingredients string
	Method      string
	Created     time.Time
}
