package service

import (
	"master-proof-api/dto"
	"master-proof-api/repository"
)

type LearningMaterialServiceImpl struct {
	LearningMaterialRepository repository.LearningMaterialRepository
}

func NewLearningMaterialService(learningMaterialRepository repository.LearningMaterialRepository) LearningMaterialService {
	return &LearningMaterialServiceImpl{
		LearningMaterialRepository: learningMaterialRepository,
	}
}

func (service *LearningMaterialServiceImpl) FindAll() []*dto.LearningMaterialResponse {
	responses, err := service.LearningMaterialRepository.FindAll()
	if err != nil {
		return nil
	}
	return responses
}
