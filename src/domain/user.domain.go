package domain

import (
	"time"
)

type UserRole string

const (
	RoleCustomer UserRole = "CUSTOMER"
	RoleAdmin    UserRole = "ADMIN"
)

type User struct {
	id          string
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	role        UserRole
	isActive    bool
	createdAt   time.Time
	updatedAt   time.Time
}

type NewUserParams struct {
	ID          string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Role        UserRole
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewUser reconstructs a User domain from persistence (e.g. database models).
func NewUser(params NewUserParams) *User {
	return &User{
		id:          params.ID,
		firstName:   params.FirstName,
		lastName:    params.LastName,
		email:       params.Email,
		phoneNumber: params.PhoneNumber,
		role:        params.Role,
		isActive:    params.IsActive,
		createdAt:   params.CreatedAt,
		updatedAt:   params.UpdatedAt,
	}
}

// --- Getters ---
func (u User) ID() string           { return u.id }
func (u User) FirstName() string    { return u.firstName }
func (u User) LastName() string     { return u.lastName }
func (u User) Email() string        { return u.email }
func (u User) PhoneNumber() string  { return u.phoneNumber }
func (u User) Role() UserRole       { return u.role }
func (u User) IsActive() bool       { return u.isActive }
func (u User) CreatedAt() time.Time { return u.createdAt }
func (u User) UpdatedAt() time.Time { return u.updatedAt }

func (u *User) GetFullName() string {
	return u.firstName + " " + u.lastName
}

func (u *User) IsAdmin() bool {
	return u.role == RoleAdmin
}

func (u *User) IsAbleToLogin() bool {
	return u.isActive
}
