package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
)

func (pstr *Repo) InsertProjectPostTable(ctx context.Context, post *models.ProjectPost) error {
	return post.Insert(ctx, pstr.db, boil.Infer())
}

func (pstr *Repo) GetProjectPostInfo(ctx context.Context, id int) (*models.ProjectPost, error) {
	return models.FindProjectPost(ctx, pstr.db, id)
}

func (pstr *Repo) UpdateProjectPostDescription(ctx context.Context, postID int, newDescription string) error {
	post, err := models.FindProjectPost(ctx, pstr.db, postID, newDescription)
	if err != nil {
		return err
	}

	post.Description = newDescription

	_, err = post.Update(ctx, pstr.db, boil.Whitelist("description"))
	return err
}

func (pstr *Repo) DeleteProjectPost(ctx context.Context, id int) (int64, error) {
	post, err := models.FindProjectPost(ctx, pstr.db, id)
	if err != nil {
		return 0, err
	}
	return post.Delete(ctx, pstr.db)
}

func (pstr *Repo) UpdateLikeCount(ctx context.Context, id int) error {
	post, err := models.FindProjectPost(ctx, pstr.db, id)
	if err != nil {
		return err
	}
	post.LikeCount += 1

	_, err = post.Update(ctx, pstr.db, boil.Whitelist("like_count"))
	return err
}
