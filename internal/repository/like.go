package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

func (lr *Repo) LikeProjectPost(ctx context.Context, likeRequest request.LikePost) error {
	post, err := models.FindProjectPost(ctx, lr.DB, likeRequest.PostID)
	if err != nil {
		return err
	}
	post.LikeCount += 1

	_, err = post.Update(ctx, lr.DB, boil.Whitelist("like_count"))
	return err

}
