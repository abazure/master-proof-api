package repository

import (
	"master-proof-api/model"
)

type LearningMaterialRepository interface {
	FindAll() ([]*model.LearningMaterial, error)
	Create(request *model.LearningMaterial) error
	FindById(id string) (*model.LearningMaterial, error)
	SaveProgress(progress *model.LearningMaterialProgress) error
	Update(request *model.LearningMaterial, id string) error
	CreateFile(file *model.File) error
	CreateIcon(request *model.Icon) error
}
