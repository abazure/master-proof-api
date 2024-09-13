package repository

import (
	"master-proof-api/model"
)

type QuizRepository interface {
	FindQuizWithCorrectAnswer(name string) ([]*model.Quiz, error)
	FindQuizWithoutCorrectAnswer(name string) ([]*model.Quiz, error)
	FindByName(name string) (*model.Quiz, error)
	SaveDiagnosticReport(request *model.UserDiagnosticReport) error
	FindUserDiagnosticReport(userId string, quizId string) (*model.UserDiagnosticReport, error)
}
