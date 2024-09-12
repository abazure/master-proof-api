package controller

import (
	"github.com/gofiber/fiber/v2"
	"master-proof-api/service"
)

type QuizControllerImpl struct {
	QuizService service.QuizService
}

func NewQuizController(quizService service.QuizService) QuizController {
	return &QuizControllerImpl{
		QuizService: quizService,
	}
}

func (controller *QuizControllerImpl) FindQuizWithCorrectAnswer(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	result, err := controller.QuizService.FindQuizWithCorrectAnswer(name)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}

func (controller *QuizControllerImpl) FindQuizWithoutCorrectAnswer(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	result, err := controller.QuizService.FindQuizWithoutCorrectAnswer(name)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": result,
	})
}
