package service

import (
	"master-proof-api/dto"
)

type UserService interface {
	Create(request dto.UserCreateRequest) error
	FindById(email string, nim string) (dto.UserResponse, error)
	Login(request dto.UserLoginRequest) (dto.UserLoginResponse, error)
}
