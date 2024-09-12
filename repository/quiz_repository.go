package repository

import "master-proof-api/model"

type QuizRepository interface {
	FindQuizWithCorrectAnswer(name string) ([]*model.Quiz, error)
	FindQuizWithoutCorrectAnswer(name string) ([]*model.Quiz, error)
}
