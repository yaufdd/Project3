package service

import (
	"context"

	"github.com/yaufdd/project3/internal/models"
)

func (psts *Servise) PublishProjectPost(ctx context.Context, post *models.ProjectPost) error {
	return psts.repo.InsertProjectPostTable(ctx, post)
}

func (psts *Servise) ReadProjectPost(ctx context.Context, id int) (*models.ProjectPost, error) {
	return psts.repo.GetProjectPostInfo(ctx, id)
}

func (psts *Servise) UpdateProjectPostDescription(ctx context.Context, postID int, newDescriptin string) error {
	return psts.repo.UpdateProjectPostDescription(ctx, postID, newDescriptin)
}

func (psts *Servise) DeleteProjectPost(ctx context.Context, id int) (int64, error) {
	return psts.repo.DeleteProjectPost(ctx, id)
}
