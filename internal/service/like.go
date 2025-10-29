package service

import (
	"context"
)

func (ls *Servise) LikeProjectPost(ctx context.Context, postID int) error {
	return ls.repo.LikeProjectPost(ctx, postID)
}
