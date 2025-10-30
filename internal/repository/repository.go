// Package repository responsbile for functions that saves to db
package repository

import (
	"database/sql"
)

// Repo struct for repository layer
type Repo struct {
	db *sql.DB
}

// NewRepository function for make Repository object
func NewRepository(db *sql.DB) *Repo {
	return &Repo{db: db}
}
