package repo

import (
	"context"

	"goboilerplate.com/src/models"
	"goboilerplate.com/src/pkg/database"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, opt models.User) (models.User, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
}

type UserRepo struct {
	db database.Database
}

func NewUserRepo(db database.Database) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(ctx context.Context, opt models.User) (models.User, error) {
	if err := r.db.WithContext(ctx).Create(&opt); err != nil {
		return models.User{}, err
	}
	return opt, nil
}

func (r *UserRepo) GetUserByUsername(ctx context.Context, username string) (models.User, error) {
	var user models.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user); err != nil {
		return models.User{}, err
	}
	return user, nil
}