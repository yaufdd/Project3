package service

import (
	"github.com/yaufdd/project3/internal/repository"
)

// Servise struct for repository layer
type Servise struct {
	repo *repository.Repo
}

// NewHandler function for make Repository object
func NewServise(repo *repository.Repo) *Servise {
	return &Servise{repo: repo}
}
