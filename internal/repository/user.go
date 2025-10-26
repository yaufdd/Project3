// Package repository responsbile for functions that saves to db
package repository

import (
	"context"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

// CreateUser save data of new user into Users table
func (ur *Repo) CreateUser(ctx context.Context, user *models.User) error {
	return user.Insert(ctx, ur.db, boil.Infer())
}

// ReadUser return data of user by id
func (ur *Repo) ReadUser(ctx context.Context, id int) (*models.User, error) {
	return models.FindUser(ctx, ur.db, id)
}

// UpdateUser chages user. Need to send full json object of user.
func (ur *Repo) UpdateUsername(ctx context.Context, newUserInfo *models.User) error {
	user, err := models.FindUser(ctx, ur.db, newUserInfo.ID)
	if err != nil {
		return nil
	}

	user.Username = newUserInfo.Username
	_, err = user.Update(ctx, ur.db, boil.Whitelist("username"))

	return err
}

// DeleteUser delete user from db by id.
// If returns, 0 without error, that means user was not deleted
func (ur *Repo) DeleteUser(ctx context.Context, id int) (int64, error) {
	user, err := models.FindUser(ctx, ur.db, id)
	if err != nil {
		return 0, err
	}
	return user.Delete(ctx, ur.db)
}

func (ur *Repo) GetUserID(ctx context.Context, request request.GetUserID) (int, error) {
	user, err := models.Users(models.UserWhere.Username.EQ(request.Username)).One(ctx, ur.db)
	if err != nil {
		return 0, err
	}
	return user.ID, err
}
