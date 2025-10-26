package service

import (
	"context"

	"github.com/yaufdd/project3/internal/request"
)

func (ls *Servise) LikeProjectPost(ctx context.Context, likeRequest request.LikePost) error {
	return ls.repo.LikeProjectPost(ctx, likeRequest)
}
