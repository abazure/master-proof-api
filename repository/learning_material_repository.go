package repository

import (
	"master-proof-api/dto"
)

type LearningMaterialRepository interface {
	FindAll() ([]*dto.LearningMaterialResponse, error)
}
