package repository

import (
	"master-proof-api/model"
)

type UserRepository interface {
	Save(user *model.User) error
	FindById(email string, nim string) (*model.User, error)
	FindByRole(role string) ([]*model.User, error)
	UpdatePhotoProfile(id, photoUrl string) error
}
