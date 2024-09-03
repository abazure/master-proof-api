package dto

type UserCreateRequest struct {
	Nim      string `validate:"required" json:"nim"`
	Name     string `validate:"required,min=1" json:"name"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,max=32" json:"password"`
}
type UserLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,max=32" json:"password"`
}
type UserPasswordResetRequest struct {
	Email string `validate:"required,email" json:"email"`
}
type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserResponse struct {
	Nim   string `json:"nim"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
