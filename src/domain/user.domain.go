package domain

import "goboilerplate.com/src/models"

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Username    string
	Password    string
	Role        string
	DateOfBirth string
}

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

func (u *User) IsAbleToLogin() bool {
	return u.Role == "admin" || u.Role == "user"
}

// FromModel converts models.User to domain.User
func FromModel(modelUser models.User) User {
	return User{
		ID:          modelUser.ID,
		FirstName:   modelUser.FirstName,
		LastName:    modelUser.LastName,
		Username:    modelUser.Username,
		Password:    modelUser.Password,
		Role:        modelUser.Role,
		DateOfBirth: modelUser.DateOfBirth,
	}
}

// ToModel converts domain.User to models.User
func (u *User) ToModel() models.User {
	return models.User{
		ID:          u.ID,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Username:    u.Username,
		Password:    u.Password,
		Role:        u.Role,
		DateOfBirth: u.DateOfBirth,
	}
}