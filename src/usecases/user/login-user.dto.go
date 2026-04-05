package user

type LoginUserRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}
