package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
)

func (cmr *Repo) InsertCommentTable(ctx context.Context, comment *models.Comment) error {
	return comment.Insert(ctx, cmr.db, boil.Infer())
}
