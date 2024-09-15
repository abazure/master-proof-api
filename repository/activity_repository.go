package repository

import "master-proof-api/model"

type ActivityRepository interface {
	CreateActivity(request *model.Activity) error
}
