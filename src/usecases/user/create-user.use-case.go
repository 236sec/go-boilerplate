package user

import (
	"errors"

	"goboilerplate.com/src/models"
	"goboilerplate.com/src/repo"
	"goboilerplate.com/src/usecases"
	"gorm.io/gorm"
)



type ICreateUserUseCase interface {
	Apply(req CreateUserRequest) (CreateUserResponse, error)
}

type CreateUserUseCase struct {
	userRepo repo.IUserRepo
}

func NewCreateUserUseCase(userRepo repo.IUserRepo) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo}
}

func (u *CreateUserUseCase) Apply(req CreateUserRequest) (CreateUserResponse, error) {
	existingUser, err := u.userRepo.GetUserByUsername(req.Username)
	if err == nil && existingUser.ID != 0 {
		return CreateUserResponse{}, usecases.ErrUserAlreadyExists
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return CreateUserResponse{}, usecases.ErrInternalServerError
	}
	
	newUser, err := u.userRepo.CreateUser(models.User{
		Username:    req.Username,
		Password:    req.Password,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		Role:        "user", // Default role
	})
	if err != nil {
		return CreateUserResponse{}, usecases.ErrCannotCreateUser
	}
	
	return CreateUserResponse{
		ID: newUser.ID,
	}, nil
}