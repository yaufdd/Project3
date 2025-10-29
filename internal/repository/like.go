package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
)

func (lr *Repo) LikeProjectPost(ctx context.Context, postID int) error {
	post, err := models.FindProjectPost(ctx, lr.DB, postID)
	if err != nil {
		return err
	}
	post.LikeCount += 1

	_, err = post.Update(ctx, lr.DB, boil.Whitelist("like_count"))
	return err

}
