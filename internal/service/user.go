// Package servies is layer between handler and reposirory.
// Resposbile for functions that bnot belong to handler and saving to db.
package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yaufdd/project3/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// CRUD for user
func (us *Servise) CreateUser(ctx context.Context, newUser *models.User) (int, *sql.Tx, error) {
	if newUser.Password == "" {
		return 0, nil, errors.New("invalid format of credential")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, nil, err
	}
	newUser.Password = string(hashedPassword)
	u := models.User{
		Username: newUser.Username,
		Password: newUser.Password,
		Email:    newUser.Email,
		Role:     newUser.Role,
	}
	return us.repo.InsertUserTable(ctx, &u)
}

func (us *Servise) ReadUser(ctx context.Context, id int) (*models.User, error) {
	return us.repo.GetUserInfo(ctx, id)
}

func (us *Servise) UpdateUsername(ctx context.Context, userID int, newUsername string) error {
	return us.repo.UpdateUsername(ctx, userID, newUsername)
}

func (us *Servise) DeleteUser(ctx context.Context, id int) (int64, error) {
	return us.repo.DeleteUser(ctx, id)
}

func (us *Servise) GetUserID(ctx context.Context, username string) (int, error) {
	user, err := us.repo.GetUserInfoByUsername(ctx, username)
	if err != nil {
		return 0, err
	}
	return user.ID, err

}

func (us *Servise) Registration(ctx context.Context, newUser *models.User) (string, string, error) {
	authJWTFacade := NewAuthJWTFacade(us)
	return authJWTFacade.Registration(ctx, newUser)
}

func (us *Servise) Authentication(ctx context.Context, reqUsername, reqPassword string) (string, string, error) {
	authJWTFacade := NewAuthJWTFacade(us)
	return authJWTFacade.Authentication(ctx, reqUsername, reqPassword)
}
