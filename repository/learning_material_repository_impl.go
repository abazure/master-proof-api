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
func (repository LearningMaterialRepositoryImpl) FindById(id string) (*model.LearningMaterial, error) {
	var learningMaterial model.LearningMaterial
	result := repository.DB.Model(&model.LearningMaterial{}).Preload("File").Preload("Icon").Where("id = ?", id).First(&learningMaterial)
	if result.Error != nil {
		return nil, result.Error
	}
	return &learningMaterial, nil
}

func (repository LearningMaterialRepositoryImpl) SaveProgress(progress *model.LearningMaterialProgress) error {
	return repository.DB.Model(model.LearningMaterialProgress{}).Create(progress).Error
}
func (repository LearningMaterialRepositoryImpl) Update(request *model.LearningMaterial, id string) error {
	return repository.DB.Model(&model.LearningMaterial{}).Where("id = ?", id).Updates(request).Error
}
func (repository LearningMaterialRepositoryImpl) CreateFile(file *model.File) error {
	return repository.DB.Model(&model.File{}).Create(file).Error
}

func (repository LearningMaterialRepositoryImpl) CreateIcon(request *model.Icon) error {
	return repository.DB.Model(&model.Icon{}).Create(request).Error

}
func (repository LearningMaterialRepositoryImpl) Delete(id string) error {
	return repository.DB.Model(&model.LearningMaterial{}).Where("id = ?", id).Delete(&model.LearningMaterial{}).Error
}
