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
func (repository *ActivityRepositoryImpl) FindByUserIdAndActivityId(userId string, activityId string) (*model.UserActivity, error) {
	var user model.UserActivity
	result := repository.DB.Model(&model.UserActivity{}).Where("user_id = ? AND activity_id = ?", userId, activityId).Order("created_at DESC").Take(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repository *ActivityRepositoryImpl) UpdateUserActivity(id string, comment string) error {
	return repository.DB.Model(model.UserActivity{}).Where("id = ?", id).Update("comment", comment).Error
}
func (repository *ActivityRepositoryImpl) FindUserActivityByUserId(userId string) ([]*model.UserActivity, error) {
	subQuery := repository.DB.Model(&model.UserActivity{}).
		Select("activity_id, MAX(created_at) as created_at").
		Where("user_id = ?", userId).
		Group("activity_id")

	var userActivities []*model.UserActivity
	result := repository.DB.Joins("JOIN (?) AS subquery ON user_activities.activity_id = subquery.activity_id AND user_activities.created_at = subquery.created_at", subQuery).Preload("Activity").Preload("File").
		Where("user_id = ?", userId).
		Find(&userActivities)
	if result.Error != nil {
		return nil, result.Error
	}
	return userActivities, nil
}
func (repository *ActivityRepositoryImpl) FindOneUserActivityByUserId(userId string, id string) (*model.UserActivity, error) {
	subQuery := repository.DB.Model(&model.UserActivity{}).
		Select("activity_id, MAX(created_at) as created_at").
		Where("user_id = ?", userId).
		Group("activity_id")

	var userActivities *model.UserActivity
	result := repository.DB.Joins("JOIN (?) AS subquery ON user_activities.activity_id = subquery.activity_id AND user_activities.created_at = subquery.created_at", subQuery).Preload("Activity").Preload("File").
		Where("user_id = ? and id = ? ", userId, id).
		Find(&userActivities)
	if result.Error != nil {
		return nil, result.Error
	}
	return userActivities, nil
}
