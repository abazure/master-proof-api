package controller

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"master-proof-api/dto"
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
func (controller *LearningMaterialControllerImpl) Create(ctx *fiber.Ctx) error {

	pdfFile, _ := ctx.FormFile("file")
	icon, _ := ctx.FormFile("icon")
	title := ctx.FormValue("title")
	description := ctx.FormValue("description")

	request := dto.CreateLearningMaterialRequest{
		Title:       title,
		Description: description,
		File:        pdfFile,
		FileName:    uuid.New().String(),
		Icon:        icon,
		IconName:    uuid.New().String(),
	}
	err := controller.Validate.Struct(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	err = controller.LearningMaterialService.Create(&request)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		} else {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success upload learning material",
	})
}
