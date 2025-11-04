// Package repository responsbile for functions that saves to db
package repository

import (
	"context"
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

func (r *Repo) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}
