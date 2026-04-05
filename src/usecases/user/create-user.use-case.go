package user

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"goboilerplate.com/src/domain"
	"goboilerplate.com/src/models"
	"goboilerplate.com/src/pkg/database"
	"goboilerplate.com/src/repo"
	"goboilerplate.com/src/usecases"
)

var createUserTracer = otel.Tracer("usecase.createuser")

type ICreateUserUseCase interface {
	Apply(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error)
}

type CreateUserUseCase struct {
	userRepo repo.IUserRepo
}

func NewCreateUserUseCase(userRepo repo.IUserRepo) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo}
}

func (u *CreateUserUseCase) Apply(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, span := createUserTracer.Start(ctx, "CreateUserUseCase.Apply")
	defer span.End()

	existingUser, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return &CreateUserResponse{}, usecases.ErrUserAlreadyExists
	}
	if err != nil && !errors.Is(err, database.ErrRecordNotFound) {
		return &CreateUserResponse{}, usecases.ErrInternalServerError
	}

	newUser, err := u.userRepo.CreateUser(ctx, models.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Role:        string(domain.RoleCustomer),
	})
	if err != nil {
		return &CreateUserResponse{}, usecases.ErrCannotCreateUser
	}

	return &CreateUserResponse{
		ID: newUser.ID.String(),
	}, nil
}
