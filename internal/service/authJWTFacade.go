package service

import (
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yaufdd/project3/internal/auth"
	"github.com/yaufdd/project3/internal/models"
	"golang.org/x/crypto/bcrypt"
)

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
	accessToken, refreshToken, jti, accessExpire, refreshExpire, err := f.tokenGenerator.generateTokens(userID, newUser.Role)
	if err != nil {
		return "", "", err
	}
	hashedToken := f.hasher.HashToken(refreshToken)
	err = f.database.saveRefreshToken(ctx, f.service, tx, userID, hashedToken, jti, refreshExpire)
	if err != nil {
		return "", "", err
	}
	if err := f.database.saveTokenToRedis(ctx, f.service, refreshToken, time.Until(refreshExpire)); err != nil {
		return "", "", err
	}
	if err = f.database.saveTokenToRedis(ctx, f.service, accessToken, time.Until(accessExpire)); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err
}

func (f *AuthJWTFacade) Authentication(ctx context.Context, reqUsername, reqPassword string) (string, string, error) {
	user, err := f.database.checkCredential(ctx, f.service, reqUsername, reqPassword)
	if err != nil {
		return "", "", err
	}
	accessToken, refreshToken, jti, accessExpire, refreshExpire, err := f.tokenGenerator.generateTokens(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}
	tx, err := f.service.repo.BeginTx(ctx)
	if err != nil {
		return "", "", err
	}
	hashedToken := f.hasher.HashToken(refreshToken)
	err = f.database.saveRefreshToken(ctx, f.service, tx, user.ID, hashedToken, jti, refreshExpire)
	if err != nil {
		return "", "", err
	}
	defer func() {
		f.database.rollBack(tx, err)
	}()
	if err = f.database.saveTokenToRedis(ctx, f.service, refreshToken, time.Until(refreshExpire)); err != nil {
		return "", "", err
	}
	if err = f.database.saveTokenToRedis(ctx, f.service, accessToken, time.Until(accessExpire)); err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, err

}

func (f *AuthJWTFacade) AuthByRefreshToken(ctx context.Context, refreshToken string) (string, string, error) {
	tokenHash := f.hasher.HashToken(refreshToken)
	refreshTokenModel, err := f.database.getRefreshTokenModel(ctx, f.service, tokenHash)
	if err != nil {
		return "", "", err
	}
	if refreshTokenModel.Revoked || refreshTokenModel.ExpiresAt.Before(time.Now()) {
		return "", "", errors.New("refresh token is old")
	}
	if err := f.database.makeRefreshTokenRevoked(ctx, f.service, tokenHash); err != nil {
		return "", "", err
	}
	user, err := f.service.repo.GetUserInfo(ctx, refreshTokenModel.UserID)
	if err != nil {
		return "", "", err
	}
	newAccessToken, newRefereshToken, jti, newAccessExpire, newRefreshExpire, err := f.tokenGenerator.generateTokens(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}
	tx, err := f.service.repo.BeginTx(ctx)
	if err != nil {
		return "", "", err
	}
	newHashedToken := f.hasher.HashToken(newRefereshToken)
	err = f.database.saveRefreshToken(ctx, f.service, tx, user.ID, newHashedToken, jti, newRefreshExpire)
	if err != nil {
		return "", "", err
	}
	defer func() {
		f.database.rollBack(tx, err)
	}()
	if err = f.database.saveTokenToRedis(ctx, f.service, refreshToken, time.Until(newRefreshExpire)); err != nil {
		return "", "", err
	}
	if err = f.database.saveTokenToRedis(ctx, f.service, newAccessToken, time.Until(newAccessExpire)); err != nil {
		return "", "", err
	}
	return newAccessToken, newRefereshToken, err
}

type Database struct{}

func (d *Database) saveTokenToRedis(ctx context.Context, s *Servise, token string, refreshExpire time.Duration) error {
	return s.rcache.SaveToken(ctx, token, refreshExpire)
}

func (d *Database) checkCredential(ctx context.Context, s *Servise, reqUsername, reqPassword string) (*models.User, error) {
	user, err := s.repo.GetUserInfoByUsername(ctx, reqUsername)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqPassword)); err != nil {
		return nil, err
	}
	return user, err
}

func (d *Database) saveNewUser(ctx context.Context, s *Servise, newUser *models.User) (int, *sql.Tx, error) {
	userID, tx, err := s.CreateUser(ctx, newUser)
	if err != nil {
		return 0, nil, err
	}
	return userID, tx, err
}

func (d *Database) saveRefreshToken(ctx context.Context, s *Servise, tx *sql.Tx, userID int, hashedToken, jti string, expiresAt time.Time) error {
	t := models.RefreshToken{
		Jti:       jti,
		UserID:    userID,
		TokenHash: hashedToken,
		ExpiresAt: expiresAt,
		Revoked:   false,
	}
	_, err := s.repo.InsertIntoRefreshTokenTable(ctx, &t, tx)
	if err != nil {
		return err
	}
	return err
}
func (d *Database) getRefreshTokenModel(ctx context.Context, s *Servise, tokenHash string) (*models.RefreshToken, error) {
	rTokenModel, err := s.repo.GetRefreshToken(ctx, tokenHash)
	if err != nil {
		return nil, err
	}
	return rTokenModel, err
}

func (d *Database) makeRefreshTokenRevoked(ctx context.Context, s *Servise, tokenHash string) error {
	affrow, err := s.repo.MakeRefreshTokenRevoked(ctx, tokenHash)
	if err != nil {
		return err
	}
	if affrow == 0 {
		return errors.New("no row was updated")
	}
	return err
}

func (d *Database) rollBack(tx *sql.Tx, registrationError error) {
	if registrationError != nil {
		fmt.Println("rolled back!")
		tx.Rollback()
	}
	fmt.Println("comitted")
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
	accessExpire time.Time,
	refreshExpire time.Time,
	err error,

) {
	privBytes, _ := os.ReadFile(os.Getenv("JWT_PRIVATE_PATH"))
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return "", "", "", time.Time{}, time.Time{}, err
	}
	accessToken, accessExpire, err = t.getAccessT(privKey, int64(userID), userRole)
	if err != nil {
		return "", "", "", time.Time{}, time.Time{}, err
	}
	refreshToken, jti, refreshExpire, err = t.getRefreshT(privKey, int64(userID))
	if err != nil {
		return "", "", "", time.Time{}, time.Time{}, err
	}

	return accessToken, refreshToken, jti, accessExpire, refreshExpire, err
}

func (t *TokenGenerator) getAccessT(privateKey *rsa.PrivateKey, userID int64, role string) (string, time.Time, error) {
	accessTokenDur := 15 * time.Minute
	auth := auth.NewAuthJWT(privateKey, "project3", accessTokenDur)

	accessToken, accessExpire, err := auth.IssueAccessToken(userID, []string{role})
	if err != nil {
		return "", time.Time{}, err
	}
	return accessToken, accessExpire, err
}

func (t *TokenGenerator) getRefreshT(privateKey *rsa.PrivateKey, userID int64) (refreshToken string, jti string, refreshExpire time.Time, err error) {
	refreshTokenDur := 7 * 24 * time.Hour
	auth := auth.NewAuthJWT(privateKey, "project3", refreshTokenDur)

	refreshToken, jti, refreshExpire, err = auth.IssueRefreshToken(userID)
	if err != nil {
		return "", "", time.Time{}, err
	}
	return refreshToken, jti, refreshExpire, err
}
