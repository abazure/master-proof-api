package repository

import (
	"gorm.io/gorm"
	"master-proof-api/dto"
	"master-proof-api/model"
)

type ProgressRepositoryImpl struct {
	DB *gorm.DB
}

func NewProgressRepository(db *gorm.DB) ProgressRepository {
	return &ProgressRepositoryImpl{DB: db}

}

func (repository *ProgressRepositoryImpl) GetLearningMaterialData(userId string) (*dto.CountData, error) {
	var result *dto.CountData
	dataLearningMaterial := repository.DB.Model(&model.LearningMaterial{}).Select("COUNT(DISTINCT id) as total_materials").Take(&result)
	if dataLearningMaterial.Error != nil {
		return nil, dataLearningMaterial.Error
	}
	dataLearningMaterialProgress := repository.DB.Model(&model.LearningMaterialProgress{}).Select("COUNT(DISTINCT learning_material_id) as finished_materials").Where("user_id = ?", userId).Take(&result)
	if dataLearningMaterialProgress.Error != nil {
		return nil, dataLearningMaterialProgress.Error
	}
	return result, nil
}

func (repository *ProgressRepositoryImpl) GetDiagnosticTestData(userId string) (*dto.CountData, error) {
	var result *dto.CountData
	dataDiagnostic := repository.DB.Model(&model.QuizCategory{}).
		Joins("LEFT JOIN quizzes ON quizzes.quiz_category_id = quiz_categories.id").
		Select("count(quizzes.id) as total_materials").
		Where("quiz_categories.name = ?", "diagnostic").
		Take(&result)
	if dataDiagnostic.Error != nil {
		return nil, dataDiagnostic.Error
	}
	dataUserDiagnostic := repository.DB.
		Model(&model.UserDiagnosticReport{}).
		Select("COUNT(DISTINCT quiz_id) as finished_materials").
		Where("user_id = ?", userId).
		Take(&result)
	if dataUserDiagnostic.Error != nil {
		return nil, dataUserDiagnostic.Error
	}
	return result, nil
}
func (repository *ProgressRepositoryImpl) GetCompetenceData(userId string) (*dto.CountData, error) {
	var result *dto.CountData
	dataCompetence := repository.DB.Model(&model.QuizCategory{}).
		Joins("LEFT JOIN quizzes ON quizzes.quiz_category_id = quiz_categories.id").
		Select("count(quizzes.id) as total_materials").
		Where("quiz_categories.name = ?", "competence").
		Take(&result)
	if dataCompetence.Error != nil {
		return nil, dataCompetence.Error
	}
	dataUserCompetence := repository.DB.
		Model(&model.UserCompetenceReports{}).
		Select("COUNT(DISTINCT quiz_name) as finished_materials").
		Where("user_id = ?", userId).
		Take(&result)
	if dataUserCompetence.Error != nil {
		return nil, dataUserCompetence.Error
	}
	return result, nil
}
func (repository *ProgressRepositoryImpl) GetActivityData(userId string) (*dto.CountData, error) {
	var result *dto.CountData
	dataActivity := repository.DB.Model(&model.Activity{}).
		Select("count(id) as total_materials").
		Take(&result)
	if dataActivity.Error != nil {
		return nil, dataActivity.Error
	}
	dataUserActivity := repository.DB.
		Model(&model.UserActivity{}).
		Select("COUNT(DISTINCT activity_id) as finished_materials").
		Where("user_id = ?", userId).
		Take(&result)
	if dataUserActivity.Error != nil {
		return nil, dataUserActivity.Error
	}
	return result, nil
}
