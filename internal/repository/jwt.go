package repository

import (
	"context"
	"database/sql"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
)

func (tr *Repo) InsertIntoRefreshTokenTable(ctx context.Context, newToken *models.RefreshToken, tx *sql.Tx) (int, error) {
	err := newToken.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return 0, err
	}
	return newToken.ID, err
}

func (tr *Repo) GetRefreshToken(ctx context.Context, tokenHash string) (*models.RefreshToken, error) {
	return models.RefreshTokens(models.RefreshTokenWhere.TokenHash.EQ(tokenHash)).One(ctx, tr.db)
}

func (tr *Repo) MakeRefreshTokenRevoked(ctx context.Context, tokenHash string) (int64, error) {
	token, err := models.RefreshTokens(models.RefreshTokenWhere.TokenHash.EQ(tokenHash)).One(ctx, tr.db)
	if err != nil {
		return 0, err
	}
	token.Revoked = true
	return token.Update(ctx, tr.db, boil.Whitelist("revoked"))

}
