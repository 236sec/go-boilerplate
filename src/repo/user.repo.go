package repo

import (
	"goboilerplate.com/src/models"
	"goboilerplate.com/src/utils/database"
)

type IUserRepo interface {
	CreateUser(opt models.User) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
}

type UserRepo struct {
	db database.Database
}

func NewUserRepo(db database.Database) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(opt models.User) (models.User, error) {
	if err := r.db.Create(&opt); err != nil {
		return models.User{}, err
	}
	return opt, nil
}

func (r *UserRepo) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user); err != nil {
		return models.User{}, err
	}
	return user, nil
}