package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"master-proof-api/service"
)

type LearningMaterialControllerImpl struct {
	LearningMaterialService service.LearningMaterialService
	Validate                *validator.Validate
}

func NewLearningMaterialController(learningMaterialService service.LearningMaterialService, validate *validator.Validate) LearningMaterialController {
	return &LearningMaterialControllerImpl{
		LearningMaterialService: learningMaterialService,
		Validate:                validate,
	}
}

func (controller *LearningMaterialControllerImpl) FindAll(ctx *fiber.Ctx) error {
	responses := controller.LearningMaterialService.FindAll()
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": responses,
	})
}
