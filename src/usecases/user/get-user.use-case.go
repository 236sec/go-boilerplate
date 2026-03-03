package user

import (
	"goboilerplate.com/src/repo"
	"goboilerplate.com/src/usecases"
)



type IGetUserUseCase interface {
	Apply(username string) (GetUserResponse, error)
}

type GetUserUseCase struct {
	userRepo repo.IUserRepo
}

func NewGetUserUseCase(userRepo repo.IUserRepo) *GetUserUseCase {
	return &GetUserUseCase{userRepo: userRepo}
}

func (u *GetUserUseCase) Apply(username string) (GetUserResponse, error) {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return GetUserResponse{}, usecases.ErrorUserNotFound
	}
	return GetUserResponse{
		ID:       user.ID,
		Username: user.Username,
		FirstName: user.FirstName,
		LastName: user.LastName,
		DateOfBirth: user.DateOfBirth,
	}, nil
}