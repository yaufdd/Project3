package handler

import (
	"github.com/yaufdd/project3/internal/service"
)

// Repo struct for repository layer
type Handler struct {
	service *service.Servise
}

// NewHandler function for make Repository object
func NewHandler(service *service.Servise) *Handler {
	return &Handler{service: service}
}
