// Package repository responsbile for functions that saves to db
package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser save data of new user into Users table
func (ur *Repo) CreateUser(ctx context.Context, newUser *request.CreateUserRequest) (int, error) {
	if newUser.Password == "" {
		return 0, errors.New("invalid format of credential")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	newUser.Password = string(hashedPassword)
	u := models.User{
		Username: newUser.Username,
		Password: newUser.Password,
		Email:    newUser.Email,
		Role:     newUser.Role,
	}
	if err := u.Insert(ctx, ur.DB, boil.Infer()); err != nil {
		return 0, err
	}
	return u.ID, err

}

// ReadUser return data of user by id
func (ur *Repo) ReadUser(ctx context.Context, userID int) (*models.User, error) {
	return models.FindUser(ctx, ur.DB, userID)
}

// UpdateUser chages user. Need to send full json object of user.
func (ur *Repo) UpdateUsername(ctx context.Context, newUserInfo *models.User) error {
	user, err := models.FindUser(ctx, ur.DB, newUserInfo.ID)
	if err != nil {
		return err
	}

	user.Username = newUserInfo.Username
	_, err = user.Update(ctx, ur.DB, boil.Whitelist("username"))

	return err
}

// DeleteUser delete user from db by id.
// If returns, 0 without error, that means user was not deleted
func (ur *Repo) DeleteUser(ctx context.Context, id int) (int64, error) {
	user, err := models.FindUser(ctx, ur.DB, id)
	if err != nil {
		return 0, err
	}
	return user.Delete(ctx, ur.DB)
}

func (ur *Repo) GetUserID(ctx context.Context, request request.GetUserID) (int, error) {
	user, err := models.Users(models.UserWhere.Username.EQ(request.Username)).One(ctx, ur.DB)
	if err != nil {
		return 0, err
	}
	return user.ID, err
}

func (ur *Repo) AuthenticateUser(ctx context.Context, credential *request.AuthCredential) (*models.User, error) {
	user, err := models.Users(models.UserWhere.Username.EQ(credential.Username)).One(ctx, ur.DB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("invalid username or password")
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password)); err != nil {
		return nil, errors.New("invalid username or password")
	}
	return user, err

}
