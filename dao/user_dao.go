package dao

import (
	"context"
	"db-playground/model"
	"gorm.io/gorm"
)

type UserDaoPGDB struct {
	db *gorm.DB
}

func NewUserDaoPGDB(db *gorm.DB) *UserDaoPGDB {
	return &UserDaoPGDB{
		db: db,
	}
}

var _ UserDao = &UserDaoPGDB{}

func (u *UserDaoPGDB) Create(ctx context.Context, user *model.User) error {
	return u.db.Create(user).Error
}

func (u *UserDaoPGDB) BulkCreate(ctx context.Context, users []*model.User) error {
	return u.db.Create(users).Error
}

func (u *UserDaoPGDB) Get(ctx context.Context, id string) (*model.User, error) {
	var user *model.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (u *UserDaoPGDB) Update(ctx context.Context, user *model.User) error {
	return u.db.Save(user).Error
}

func (u *UserDaoPGDB) Delete(ctx context.Context, ids []string) error {
	return u.db.Where("id IN ?", ids).Delete(&model.User{}).Error
}
