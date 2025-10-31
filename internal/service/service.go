package service

import (
	"github.com/yaufdd/project3/internal/cache"
	"github.com/yaufdd/project3/internal/repository"
)

// Servise struct for repository layer
type Servise struct {
	repo   *repository.Repo
	rcache *cache.Redis
}

// NewHandler function for make Repository object
func NewServise(repo *repository.Repo, rcache *cache.Redis) *Servise {
	return &Servise{repo: repo, rcache: rcache}
}
