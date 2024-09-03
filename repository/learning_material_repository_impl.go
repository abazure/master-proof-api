package repository

import (
	"gorm.io/gorm"
	"master-proof-api/dto"
)

type LearningMaterialRepositoryImpl struct {
	DB *gorm.DB
}

func NewLearningMaterialRepository(db *gorm.DB) LearningMaterialRepository {
	return &LearningMaterialRepositoryImpl{
		DB: db,
	}
}

func (repository LearningMaterialRepositoryImpl) FindAll() ([]*dto.LearningMaterialResponse, error) {
	var learningMaterial []*dto.LearningMaterialResponse
	err := repository.DB.Table("learning_materials as lm").Select("lm.id ,lm.title, lm.description, f.url").Joins("join files as f on f.id = lm.file_id").Take(&learningMaterial).Error
	return learningMaterial, err
}
