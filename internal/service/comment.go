package service

import (
	"context"

	"github.com/yaufdd/project3/internal/models"
)

func (cms *Servise) AddComment(ctx context.Context, comment *models.Comment) error {
	return cms.repo.InsertCommentTable(ctx, comment)
}
