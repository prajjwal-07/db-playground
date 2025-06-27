package dao

import (
	"context"
	"db-playground/model"
)

type UserDao interface {
	Create(ctx context.Context, user *model.User) error
	BulkCreate(ctx context.Context, users []*model.User) error
	Get(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, ids []string) error
}
