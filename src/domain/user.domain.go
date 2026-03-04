package domain

import "goboilerplate.com/src/models"

type User struct {
	id          int
	firstName   string
	lastName    string
	username    string
	password    string
	role        string
	dateOfBirth string
}

func (u *User) GetFullName() string {
	return u.firstName + " " + u.lastName
}

func (u *User) IsAdmin() bool {
	return u.role == "admin"
}

func (u *User) IsAbleToLogin() bool {
	return u.role == "admin" || u.role == "user"
}

// FromModel converts models.User to domain.User
func FromModel(modelUser models.User) User {
	return User{
		id:          modelUser.ID,
		firstName:   modelUser.FirstName,
		lastName:    modelUser.LastName,
		username:    modelUser.Username,
		password:    modelUser.Password,
		role:        modelUser.Role,
		dateOfBirth: modelUser.DateOfBirth,
	}
}

// ToModel converts domain.User to models.User
func (u *User) ToModel() models.User {
	return models.User{
		ID:          u.id,
		FirstName:   u.firstName,
		LastName:    u.lastName,
		Username:    u.username,
		Password:    u.password,
		Role:        u.role,
		DateOfBirth: u.dateOfBirth,
	}
}