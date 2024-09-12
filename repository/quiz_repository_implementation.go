package repository

import (
	"gorm.io/gorm"
	"master-proof-api/model"
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
