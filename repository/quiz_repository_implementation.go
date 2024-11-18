package repository

import (
	"master-proof-api/model"

	"gorm.io/gorm"
)

type QuizRepositoryImpl struct {
	DB *gorm.DB
}

func NewQuizRepository(db *gorm.DB) QuizRepository {
	return &QuizRepositoryImpl{
		DB: db,
	}
}

func (repository *QuizRepositoryImpl) FindQuizWithCorrectAnswer(name string) ([]*model.Quiz, error) {
	var quiz []*model.Quiz
	result := repository.DB.Model(&model.Quiz{}).Preload("Questions.Answers").Where("name = ?", name).Find(&quiz)
	if result.RowsAffected == 0 {
		return []*model.Quiz{}, gorm.ErrRecordNotFound
	}
	return quiz, nil
}

func (repository *QuizRepositoryImpl) FindQuizWithoutCorrectAnswer(name string) ([]*model.Quiz, error) {
	var quiz []*model.Quiz
	result := repository.DB.Model(&model.Quiz{}).Preload("Questions.Answers").Where("name = ?", name).Find(&quiz)
	if result.RowsAffected == 0 {
		return []*model.Quiz{}, gorm.ErrRecordNotFound
	}
	return quiz, nil
}

func (repository *QuizRepositoryImpl) FindByName(name string) (*model.Quiz, error) {
	var quiz model.Quiz
	result := repository.DB.Model(&model.Quiz{}).Where("name = ?", name).Find(&quiz)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &quiz, nil
}

func (repository *QuizRepositoryImpl) SaveDiagnosticReport(request *model.UserDiagnosticReport) error {
	return repository.DB.Save(request).Error
}

func (repository *QuizRepositoryImpl) FindUserDiagnosticReport(userId string, quizId string) (*model.UserDiagnosticReport, error) {
	var result model.UserDiagnosticReport
	err := repository.DB.Model(&model.UserDiagnosticReport{}).
		Preload("DiagnosticReport").
		Where("user_id = ? AND quiz_id = ?", userId, quizId).
		Order("created_at DESC").
		Take(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repository *QuizRepositoryImpl) SaveCompetenceReport(request *model.UserCompetenceReports) error {
	return repository.DB.Save(request).Error
}

func (repository *QuizRepositoryImpl) FindUserCompetenceReport(userId string, quizId string) (*model.UserCompetenceReports, error) {
	var result model.UserCompetenceReports
	err := repository.DB.Model(&model.UserCompetenceReports{}).
		Where("user_id = ? AND quiz_name = ?", userId, quizId).
		Order("created_at DESC").
		Take(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repository *QuizRepositoryImpl) GetDiagonosticAllQuizzes() ([]*model.Quiz, error) {

	var quizzes []*model.Quiz
	err := repository.DB.
		Joins("LEFT JOIN quiz_categories ON quizzes.quiz_category_id = quiz_categories.id").
		Where("quiz_categories.name = ?", "diagnostic").
		Find(&quizzes).
		Error

	if err != nil {
		return nil, err
	}

	return quizzes, nil
}

func (repository *QuizRepositoryImpl) GetCompetenceAllQuizzes() ([]*model.Quiz, error) {

	var quizzes []*model.Quiz
	err := repository.DB.
		Joins("LEFT JOIN quiz_categories ON quizzes.quiz_category_id = quiz_categories.id").
		Where("quiz_categories.name = ?", "competence").
		Find(&quizzes).
		Error

	if err != nil {
		return nil, err
	}

	return quizzes, nil
}
