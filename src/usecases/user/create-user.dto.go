package user

type CreateUserRequest struct {
	FirstName   string `json:"firstName"   validate:"required"`
	LastName    string `json:"lastName"    validate:"required"`
	Email       string `json:"email"       validate:"required,email"`
	Password    string `json:"password"    validate:"required"`
	PhoneNumber string `json:"phoneNumber"`
}

type CreateUserResponse struct {
	ID string `json:"id"`
}
