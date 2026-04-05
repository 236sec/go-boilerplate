package user

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"goboilerplate.com/src/repo"
	"goboilerplate.com/src/usecases"
)

var getUserTracer = otel.Tracer("usecase.getuser")

type IGetUserUseCase interface {
	Apply(ctx context.Context, username string) (*GetUserResponse, error)
}

type GetUserUseCase struct {
	userRepo repo.IUserRepo
}

func NewGetUserUseCase(userRepo repo.IUserRepo) *GetUserUseCase {
	return &GetUserUseCase{userRepo: userRepo}
}

func (u *GetUserUseCase) Apply(ctx context.Context, email string) (*GetUserResponse, error) {
	ctx, span := getUserTracer.Start(ctx, "GetUserUseCase.Apply")
	defer span.End()

	user, err := u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return &GetUserResponse{}, usecases.ErrUserNotFound
	}
	return &GetUserResponse{
		ID:          user.ID.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		IsActive:    user.IsActive,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   user.UpdatedAt.Format(time.RFC3339),
	}, nil
}
