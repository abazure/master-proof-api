package repository

import (
	"gorm.io/gorm"
	"master-proof-api/model"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (repository *UserRepositoryImpl) Save(user *model.User) error {
	return repository.DB.Create(user).Error
}

func (repository *UserRepositoryImpl) FindById(email string, nim string) (*model.User, error) {
	var user model.User
	result := repository.DB.Where("email = ? OR nim = ?", email, nim).Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (repository *UserRepositoryImpl) FindByRole(role string) ([]*model.User, error) {
	var users []*model.User
	result := repository.DB.Where("role = ?", role).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (repository *UserRepositoryImpl) UpdatePhotoProfile(id, photoUrl string) error {
	result := repository.DB.Model(model.User{}).Where("id = ?", id).Update("photo_url", photoUrl)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
