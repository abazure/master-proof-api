package controller

import "github.com/gofiber/fiber/v2"

type QuizController interface {
	FindQuizWithCorrectAnswer(ctx *fiber.Ctx) error
	FindQuizWithoutCorrectAnswer(ctx *fiber.Ctx) error
	CreateUserDiagnosticReport(ctx *fiber.Ctx) error
	FindUserDiagnosticReport(ctx *fiber.Ctx) error
	CreateUserCompetenceReport(ctx *fiber.Ctx) error
}
