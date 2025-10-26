// Package servies is layer between handler and reposirory.
// Resposbile for functions that bnot belong to handler and saving to db.
package service

import (
	"context"

	"github.com/yaufdd/project3/internal/models"
	"github.com/yaufdd/project3/internal/request"
)

func (us *Servise) CreateUser(ctx context.Context, user *models.User) error {
	return us.repo.CreateUser(ctx, user)
}

func (us *Servise) ReadUser(ctx context.Context, id int) (*models.User, error) {
	return us.repo.ReadUser(ctx, id)
}

func (us *Servise) UpdateUsername(ctx context.Context, newUserInfo *models.User) error {
	return us.repo.UpdateUsername(ctx, newUserInfo)
}

func (us *Servise) DeleteUser(ctx context.Context, id int) (int64, error) {
	return us.repo.DeleteUser(ctx, id)
}

func (us *Servise) GetUserID(ctx context.Context, request request.GetUserID) (int, error) {
	return us.repo.GetUserID(ctx, request)
}
