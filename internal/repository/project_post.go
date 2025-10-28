package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

func (pstr *Repo) PublishProjectPost(ctx context.Context, post *models.ProjectPost) error {
	return post.Insert(ctx, pstr.DB, boil.Infer())
}

func (pstr *Repo) ReadProjectPost(ctx context.Context, id int) (*models.ProjectPost, error) {
	return models.FindProjectPost(ctx, pstr.DB, id)
}

func (pstr *Repo) UpdateProjectPostDescription(ctx context.Context, newDesc request.NewDescription) error {
	post, err := models.FindProjectPost(ctx, pstr.DB, newDesc.ID)
	if err != nil {
		return err
	}

	post.Description = newDesc.Description

	_, err = post.Update(ctx, pstr.DB, boil.Whitelist("description"))
	return err
}

func (pstr *Repo) DeleteProjectPost(ctx context.Context, id int) (int64, error) {
	post, err := models.FindProjectPost(ctx, pstr.DB, id)
	if err != nil {
		return 0, err
	}
	return post.Delete(ctx, pstr.DB)
}
