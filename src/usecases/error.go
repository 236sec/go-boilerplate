package usecases

import "errors"

var (
	ErrorUserNotFound       = errors.New("user not found")
	ErrorUserNotAuthorized = errors.New("user is not authorized to login")
	ErrorInvalidCredentials = errors.New("invalid credentials")
	ErrorUserNotAbleToLogin = errors.New("user is not able to login")
	ErrorCannotCreateUser      = errors.New("cannot create user")
	ErrorInternalServerError = errors.New("internal server error")
	ErrorUserAlreadyExists = errors.New("user already exists")
)