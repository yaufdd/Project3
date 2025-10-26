package service

import (
	"context"

	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

func (psts *Servise) PublishProjectPost(ctx context.Context, post *models.ProjectPost) error {
	return psts.repo.PublishProjectPost(ctx, post)
}

func (psts *Servise) ReadProjectPost(ctx context.Context, id int) (*models.ProjectPost, error) {
	return psts.repo.ReadProjectPost(ctx, id)
}

func (psts *Servise) UpdateProjectPostDescription(ctx context.Context, newDesc request.NewDescription) error {
	return psts.repo.UpdateProjectPostDescription(ctx, newDesc)
}

func (psts *Servise) DeleteProjectPost(ctx context.Context, id int) (int64, error) {
	return psts.repo.DeleteProjectPost(ctx, id)
}
