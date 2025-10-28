// Package repository responsbile for functions that saves to db
package repository

import (
	"database/sql"
)

// Repo struct for repository layer
type Repo struct {
	DB *sql.DB
}

// NewRepository function for make Repository object
func NewRepository(DB *sql.DB) *Repo {
	return &Repo{DB: DB}
}
