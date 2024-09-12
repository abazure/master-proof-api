package service

import (
	"master-proof-api/dto"
	"master-proof-api/repository"
)

type QuizServiceImpl struct {
	QuizRepository repository.QuizRepository
}

func NewQuizService(quizRepository repository.QuizRepository) QuizService {
	return &QuizServiceImpl{
		QuizRepository: quizRepository,
	}
}

func (service *QuizServiceImpl) FindQuizWithCorrectAnswer(name string) ([]*dto.QuestionWithCorrectAnswer, error) {
	quizzes, err := service.QuizRepository.FindQuizWithCorrectAnswer(name)
	if err != nil {
		return nil, err
	}

	var results []*dto.QuestionWithCorrectAnswer

	for _, quiz := range quizzes {
		for _, question := range quiz.Questions {
			questionDTO := &dto.QuestionWithCorrectAnswer{
				Id:            question.ID,
				Question:      question.Question,
				CorrectAnswer: question.CorrectAnswer,
			}
			for _, answer := range question.Answers {
				option := dto.Option{
					Value: int(answer.Value),
					Text:  answer.Text,
				}
				questionDTO.AnswerOptions = append(questionDTO.AnswerOptions, option)
			}

			results = append(results, questionDTO)
		}
	}

	return results, nil
}
func (service *QuizServiceImpl) FindQuizWithoutCorrectAnswer(name string) ([]*dto.QuestionWithoutCorrectAnswer, error) {
	quizzes, err := service.QuizRepository.FindQuizWithoutCorrectAnswer(name)
	if err != nil {
		return nil, err
	}

	var results []*dto.QuestionWithoutCorrectAnswer

	for _, quiz := range quizzes {
		for _, question := range quiz.Questions {
			questionDTO := &dto.QuestionWithoutCorrectAnswer{
				Id:       question.ID,
				Question: question.Question,
			}
			for _, answer := range question.Answers {
				option := dto.Option{
					Value: int(answer.Value),
					Text:  answer.Text,
				}
				questionDTO.AnswerOptions = append(questionDTO.AnswerOptions, option)
			}

			results = append(results, questionDTO)
		}
	}

	return results, nil
}
