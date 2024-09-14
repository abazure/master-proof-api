package service

import "master-proof-api/dto"

type QuizService interface {
	FindQuizWithCorrectAnswer(name string) ([]*dto.QuestionWithCorrectAnswer, error)
	FindQuizWithoutCorrectAnswer(name string) ([]*dto.QuestionWithoutCorrectAnswer, error)
	CreateUserDiagnosticReport(request dto.DiagnosticReportRequest) error
	FindUserDiagnosticReport(request dto.RequestGetDiagnosticResult) (*dto.ResponseDiagnosticReport, error)
	CreateUserCompetenceReport(request dto.CompetenceReportRequest) error
	FindUserCompeteceReport(request dto.RequestGetDiagnosticResult) (*dto.ResponseDiagnosticReport, error)
}
