package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"master-proof-api/dto"
	"master-proof-api/model"
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
			var correctAnswer *int
			if question.CorrectAnswer != nil { // Check if CorrectAnswer is not nil
				correctAnswer = question.CorrectAnswer // Assign the correct answer if it's present
			}
			questionDTO := &dto.QuestionWithoutCorrectAnswer{
				Id:            question.ID,
				Question:      question.Question,
				CorrectAnswer: correctAnswer,
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

func (service *QuizServiceImpl) CreateUserDiagnosticReport(request dto.DiagnosticReportRequest) error {

	quiz, _ := service.QuizRepository.FindByName(request.QuizId)
	createRequest := model.UserDiagnosticReport{
		Id:                 uuid.New().String(),
		UserId:             request.UserId,
		QuizId:             quiz.ID,
		DiagnosticReportId: request.DiagnosticReportId,
	}

	fmt.Println(createRequest)
	err := service.QuizRepository.SaveDiagnosticReport(&createRequest)
	if err != nil {
		return err
	}
	return nil

}

func (service *QuizServiceImpl) FindUserDiagnosticReport(request dto.RequestGetDiagnosticResult) (*dto.ResponseDiagnosticReport, error) {
	quiz, _ := service.QuizRepository.FindByName(request.QuizName)
	if quiz == nil {
		return &dto.ResponseDiagnosticReport{}, nil
	}
	report, err := service.QuizRepository.FindUserDiagnosticReport(request.UserId, quiz.ID)
	if err != nil {
		return nil, err
	}
	result := &dto.ResponseDiagnosticReport{
		StudentId: request.UserId,
		Type:      report.DiagnosticReportId,
		Desc:      report.DiagnosticReport.Description,
		CreatedAt: report.CreatedAt,
	}

	return result, nil

}

func (service *QuizServiceImpl) CreateUserCompetenceReport(request dto.CompetenceReportRequest) error {
	quiz, _ := service.QuizRepository.FindByName(request.QuizId)
	if quiz == nil {
		return fiber.NewError(fiber.StatusNotFound, "quiz not found")
	}
	createRequest := model.UserCompetenceReports{
		Id:       uuid.New().String(),
		UserId:   request.UserId,
		QuizName: request.QuizId,
		Score:    request.Score,
	}
	err := service.QuizRepository.SaveCompetenceReport(&createRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
func (service *QuizServiceImpl) FindUserCompetenceReport(request dto.RequestGetCompetenceResult) (*dto.ResponseCompetenceReport, error) {
	quiz, _ := service.QuizRepository.FindByName(request.QuizName)
	if quiz == nil {
		return &dto.ResponseCompetenceReport{}, nil
	}

	report, err := service.QuizRepository.FindUserCompetenceReport(request.UserId, request.QuizName)
	if err != nil {
		return nil, err
	}
	result := &dto.ResponseCompetenceReport{
		StudentId: report.UserId,
		Score:     report.Score,
		CreatedAt: report.CreatedAt,
	}
	return result, nil
}
