package service

import (
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yaufdd/project3/internal/auth"
	"github.com/yaufdd/project3/internal/models"
)

type Database struct{}

func (d *Database) saveNewUser(ctx context.Context, s *Servise, newUser *models.User) (int, *sql.Tx, error) {
	userID, tx, err := s.CreateUser(ctx, newUser)
	if err != nil {
		return 0, tx, err
	}
	return userID, tx, err
}

func (d *Database) saveRefreshToken(ctx context.Context, s *Servise, userID int, hashedToken, jti string, expiresAt time.Time) error {
	t := models.RefreshToken{
		Jti:       jti,
		UserID:    userID,
		TokenHash: hashedToken,
		ExpiresAt: expiresAt,
		Revoked:   false,
	}
	_, err := s.repo.InsertIntoRefreshTokenTable(ctx, &t)
	if err != nil {
		return err
	}
	return err
}

func (d *Database) rollBack(tx *sql.Tx, registrationError error) {
	if registrationError != nil {
		tx.Rollback()
	}
	tx.Commit()
}

type Hasher struct{}

func (h *Hasher) HashToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

type TokenGenerator struct{}

func (t *TokenGenerator) generateTokens(userID int, userRole string) (
	accessToken, refreshToken string,
	jti string,
	refreshExpire time.Time,
	err error,

) {
	privBytes, _ := os.ReadFile(os.Getenv("JWT_PRIVATE_PATH"))
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return "", "", "", time.Time{}, err
	}
	accessToken, err = t.getAccessT(privKey, int64(userID), userRole)
	if err != nil {
		return "", "", "", time.Time{}, err
	}
	refreshToken, jti, refreshExpire, err = t.getRefreshT(privKey, userID)
	if err != nil {
		return "", "", "", time.Time{}, err
	}

	return accessToken, refreshToken, jti, refreshExpire, err
}

func (t *TokenGenerator) getAccessT(privateKey *rsa.PrivateKey, userID int64, role string) (string, error) {
	accessTokenDur := 15 * time.Minute
	auth := auth.NewAuthJWT(privateKey, "project3_access", accessTokenDur)

	accessToken, err := auth.IssueAccessToken(int64(userID), []string{role})
	if err != nil {
		return "", err
	}
	return accessToken, err
}

func (t *TokenGenerator) getRefreshT(privateKey *rsa.PrivateKey, userID int) (hashedToken string, jti string, refreshExpire time.Time, err error) {
	refreshTokenDur := 7 * (24 * time.Hour)
	auth := auth.NewAuthJWT(privateKey, "project3_refreshT", refreshTokenDur)

	refreshToken, jti, refreshExpire, err := auth.IssueRefreshToken(int64(userID))
	if err != nil {
		return "", "", time.Time{}, err
	}
	return refreshToken, jti, refreshExpire, err
}

type AuthJWTFacade struct {
	database       *Database
	tokenGenerator *TokenGenerator
	hasher         *Hasher
	service        *Servise
}

func NewAuthJWTFacade(service *Servise) *AuthJWTFacade {
	return &AuthJWTFacade{database: &Database{}, tokenGenerator: &TokenGenerator{}, hasher: &Hasher{}, service: service}
}

func (f *AuthJWTFacade) Registration(ctx context.Context, newUser *models.User) (string, string, error) {
	userID, tx, err := f.database.saveNewUser(ctx, f.service, newUser)
	if err != nil {
		return "", "", err

	}
	defer func() {
		f.database.rollBack(tx, err)
	}()

	accessToken, refreshToken, jti, refreshExpire, err := f.tokenGenerator.generateTokens(userID, newUser.Role)
	if err != nil {
		return "", "", err
	}
	hashedToken := f.hasher.HashToken(refreshToken)
	if err = f.database.saveRefreshToken(ctx, f.service, userID, hashedToken, jti, refreshExpire); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err

}
