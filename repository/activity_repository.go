package repository

import "master-proof-api/model"

type ActivityRepository interface {
	CreateActivity(request *model.Activity) error
	FindAll() ([]*model.Activity, error)
	FindById(id string) (*model.Activity, error)
	CreateActivitySubmission(request *model.UserActivity) error
}
