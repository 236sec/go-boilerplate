package user

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"goboilerplate.com/src/models"
	"goboilerplate.com/src/pkg/database"
	"goboilerplate.com/src/repo/mocks"
	"goboilerplate.com/src/usecases"
)

type assertions struct {
	name           string
	request        *CreateUserRequest
	mockSetup      func(mockRepo *mocks.MockIUserRepo)
	expectedError  error
	expectedResult *CreateUserResponse
}

func TestCreateUserUseCase_Apply(t *testing.T) {
	// Define test cases
	testCases := []assertions{
		{
			name: "Error at query user - should not run domain logic",
			request: &CreateUserRequest{
				Email:     "testuser",
				Password:  "password123",
				FirstName: "Test",
				LastName:  "User",
			},
			mockSetup: func(mockRepo *mocks.MockIUserRepo) {
				// DB returns unexpected error → CreateUser should NOT be called
				mockRepo.EXPECT().GetUserByEmail(gomock.Any(), "testuser").Return(nil, errors.New("db connection error")).Times(1)
				mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(0)
			},
			expectedError:  usecases.ErrInternalServerError,
			expectedResult: &CreateUserResponse{},
		},
		{
			name: "Pass the db - run user domain logic once",
			request: &CreateUserRequest{
				Email:     "newuser",
				Password:  "password123",
				FirstName: "New",
				LastName:  "User",
			},
			mockSetup: func(mockRepo *mocks.MockIUserRepo) {
				// User not found → CreateUser should be called exactly once
				mockRepo.EXPECT().GetUserByEmail(gomock.Any(), "newuser").Return(nil, database.ErrRecordNotFound).Times(1)
				mockRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&models.User{
					ID:        uuid.Max,
					FirstName: "New",
					LastName:  "User",
					Password:  "password123",
					Role:      "user",
				}, nil).Times(1)
			},
			expectedError: nil,
			expectedResult: &CreateUserResponse{
				ID: uuid.Max.String(),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockIUserRepo(ctrl)
			tt.mockSetup(mockRepo)

			useCase := NewCreateUserUseCase(mockRepo)
			result, err := useCase.Apply(context.Background(), tt.request)

			assert.ErrorIs(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}
