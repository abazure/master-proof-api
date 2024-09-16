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

func (repository *ActivityRepositoryImpl) FindById(id string) (*model.Activity, error) {
	var activity *model.Activity
	err := repository.DB.Where("id = ?", id).Preload("File").Take(&activity).Error
	return activity, err
}
func (repository *ActivityRepositoryImpl) CreateActivitySubmission(request *model.UserActivity) error {
	return repository.DB.Create(request).Error
}
