package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
)

func (tr *Repo) InsertIntoRefreshTokenTable(ctx context.Context, newToken *models.RefreshToken) (int, error) {
	err := newToken.Insert(ctx, tr.db, boil.Infer())
	if err != nil {
		return 0, err
	}
	return newToken.ID, err
}
