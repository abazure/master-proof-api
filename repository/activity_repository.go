package repository

import "master-proof-api/model"

type ActivityRepository interface {
	CreateActivity(request *model.Activity) error
	FindAll() ([]*model.Activity, error)
	FindById(id string) (*model.Activity, error)
	CreateActivitySubmission(request *model.UserActivity) error
	FindByUserIdAndActivityId(userId string, activityId string) (*model.UserActivity, error)
	UpdateUserActivity(id string, comment string) error
	FindUserActivityByUserId(userId string) ([]*model.UserActivity, error)
	FindOneUserActivityByUserId(id string) (*model.UserActivity, error)
}
