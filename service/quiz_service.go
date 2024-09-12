package service

import "master-proof-api/dto"

type QuizService interface {
	FindQuizWithCorrectAnswer(name string) ([]*dto.QuestionWithCorrectAnswer, error)
}
