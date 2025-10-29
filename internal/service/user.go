// Package servies is layer between handler and reposirory.
// Resposbile for functions that bnot belong to handler and saving to db.
package service

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yaufdd/project3/internal/auth"
	"github.com/yaufdd/project3/internal/models"
)

// CRUD for user
func (us *Servise) CreateUser(ctx context.Context, newUser *models.User) (int, error) {
	return us.repo.CreateUser(ctx, newUser)
}

func (us *Servise) ReadUser(ctx context.Context, id int) (*models.User, error) {
	return us.repo.ReadUser(ctx, id)
}

func (us *Servise) UpdateUsername(ctx context.Context, userID int, newUsername string) error {
	return us.repo.UpdateUsername(ctx, userID, newUsername)
}

func (us *Servise) DeleteUser(ctx context.Context, id int) (int64, error) {
	return us.repo.DeleteUser(ctx, id)
}

func (us *Servise) GetUserID(ctx context.Context, username string) (int, error) {
	return us.repo.GetUserID(ctx, username)
}

func (us *Servise) Registration(ctx context.Context, newUser *models.User) (string, error) {
	tx, err := us.repo.DB.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	userID, err := us.repo.CreateUser(ctx, newUser)
	if err != nil {
		return "", err
	}

	privBytes, _ := os.ReadFile(os.Getenv("JWT_PRIVATE_PATH"))
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return "", err
	}
	accessTokenDur := 15 * time.Minute
	auth := auth.NewAuthJWT(privKey, "project3", accessTokenDur)

	accessToken, err := auth.IssueAccessToken(int64(userID), []string{newUser.Role})
	if err != nil {
		return "", err
	}
	return accessToken, err

}

func (us *Servise) Authentication(ctx context.Context, username, password string) (string, error) {
	user, err := us.repo.AuthenticateUser(ctx, username, password)
	if err != nil {
		return "", err
	}

	privBytes, _ := os.ReadFile(os.Getenv("JWT_PRIVATE_PATH"))
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privBytes)
	if err != nil {
		return "", err
	}
	accessTokenDur := 15 * time.Minute
	auth := auth.NewAuthJWT(privKey, "project3", accessTokenDur)

	accessToken, err := auth.IssueAccessToken(int64(user.ID), []string{user.Role})
	if err != nil {
		return "", err
	}
	return accessToken, err
}
