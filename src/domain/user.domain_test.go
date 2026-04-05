package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_GetFullName(t *testing.T) {
	testCases := []struct {
		name      string
		firstName string
		lastName  string
		expected  string
	}{
		{"Should return full name when both first and last names are provided", "John", "Doe", "John Doe"},
		{"Should return first name when last name is empty", "John", "", "John "},
		{"Should return last name when first name is empty", "", "Doe", " Doe"},
		{"Should return empty string when both first and last names are empty", "", "", " "},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			user := User{
				firstName: tt.firstName,
				lastName:  tt.lastName,
			}
			assert.Equal(t, tt.expected, user.GetFullName())
		})
	}
}

func TestUser_IsAdmin(t *testing.T) {
	testCases := []struct {
		name     string
		role     UserRole
		expected bool
	}{
		{"Should be admin when role is admin", RoleAdmin, true},
		{"Should not be admin when role is user", RoleCustomer, false},
		{"Should not be admin when role is empty", "", false},
		{"Should not be admin when role is guest", "guest", false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			user := User{role: tt.role}
			assert.Equal(t, tt.expected, user.IsAdmin())
		})
	}
}

func TestUser_IsAbleToLogin(t *testing.T) {
	testCases := []struct {
		name     string
		isActive bool
		expected bool
	}{
		{"Should be able to login when user is active", true, true},
		{"Should not be able to login when user is inactive", false, false},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			user := User{isActive: tt.isActive}
			assert.Equal(t, tt.expected, user.IsAbleToLogin())
		})
	}
}
