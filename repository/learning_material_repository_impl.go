package repository

import (
	"gorm.io/gorm"
	"master-proof-api/model"
)

type LearningMaterialRepositoryImpl struct {
	DB *gorm.DB
}

func NewLearningMaterialRepository(db *gorm.DB) LearningMaterialRepository {
	return &LearningMaterialRepositoryImpl{
		DB: db,
	}
}

func (repository LearningMaterialRepositoryImpl) FindAll() ([]*model.LearningMaterial, error) {
	var learningMaterials []*model.LearningMaterial
	result := repository.DB.Model(&learningMaterials).Preload("File").Preload("Icon").Find(&learningMaterials)
	if result.Error != nil {
		return nil, result.Error
	}
	return learningMaterials, nil
}

func (repository LearningMaterialRepositoryImpl) Create(request *model.LearningMaterial) error {
	return repository.DB.Model(model.LearningMaterial{}).Create(request).Error
}
