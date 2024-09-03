package service

import "master-proof-api/dto"

type LearningMaterialService interface {
	FindAll() []*dto.LearningMaterialResponse
}
