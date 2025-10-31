// Package repository responsbile for functions that saves to db
package repository

import (
	"context"
	"database/sql"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/yaufdd/project3/internal/models"
)

// InsertUserTable save data of new user into Users table
func (ur *Repo) InsertUserTable(ctx context.Context, newUser *models.User) (int, *sql.Tx, error) {
	tx, err := ur.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, nil, err
	}
	err = newUser.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return 0, nil, err
	}
	return newUser.ID, tx, err

}

// GetUserInfo return data of user by id
func (ur *Repo) GetUserInfo(ctx context.Context, userID int) (*models.User, error) {
	return models.FindUser(ctx, ur.db, userID)
}

func (ur *Repo) GetUserInfoByUsername(ctx context.Context, username string) (*models.User, error) {
	return models.Users(models.UserWhere.Username.EQ(username)).One(ctx, ur.db)
}

// UpdateUsername chages user.
func (ur *Repo) UpdateUsername(ctx context.Context, userID int, newUsername string) error {
	user, err := models.FindUser(ctx, ur.db, userID)
	if err != nil {
		return err
	}

	user.Username = newUsername
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
