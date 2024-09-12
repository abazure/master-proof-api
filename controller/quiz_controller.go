package controller

import "github.com/gofiber/fiber/v2"

type QuizController interface {
	FindQuizWithCorrectAnswer(ctx *fiber.Ctx) error
	FindQuizWithoutCorrectAnswer(ctx *fiber.Ctx) error
}
