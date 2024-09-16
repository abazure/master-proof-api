package repository

import (
	"gorm.io/gorm"
	"master-proof-api/model"
)

type ActivityRepositoryImpl struct {
	DB *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &ActivityRepositoryImpl{
		DB: db,
	}
}

func (repository *ActivityRepositoryImpl) CreateActivity(request *model.Activity) error {
	return repository.DB.Create(request).Error
}
func (repository *ActivityRepositoryImpl) FindAll() ([]*model.Activity, error) {
	var activities []*model.Activity
	err := repository.DB.Model(&model.Activity{}).Preload("File").Find(&activities).Error
	return activities, err
}
