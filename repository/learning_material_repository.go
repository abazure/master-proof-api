package repository

import (
	"master-proof-api/model"
)

type LearningMaterialRepository interface {
	FindAll() ([]*model.LearningMaterial, error)
	Create(request *model.LearningMaterial) error
}
