package controller

import (
	"errors"
	"firebase.google.com/go/v4/auth"
	"fmt"
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
func (controller *LearningMaterialControllerImpl) FindByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := controller.LearningMaterialService.FindById(id)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}
func (controller *LearningMaterialControllerImpl) SaveProgress(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*auth.Token)
	userId := token.Claims["user_id"].(string)
	LearningMaterialId := ctx.Params("id")
	request := dto.UserSaveProgressRequest{
		UserID:             userId,
		LearningMaterialId: LearningMaterialId,
		IsFinished:         false,
	}
	err := controller.LearningMaterialService.UpdateProgress(&request)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success save progress",
	})
}
func (controller *LearningMaterialControllerImpl) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	pdfFile, _ := ctx.FormFile("file")
	icon, _ := ctx.FormFile("icon")
	title := ctx.FormValue("title")
	description := ctx.FormValue("description")

	request := dto.UpdateLearningMaterialRequest{
		Id:          id,
		Title:       title,
		Description: description,
		File:        pdfFile,
		FileName:    uuid.New().String(),
		Icon:        icon,
		IconName:    uuid.New().String(),
	}

	fmt.Println(request)
	err := controller.Validate.Struct(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	err = controller.LearningMaterialService.UpdateLearningMaterial(&request)
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
		"message": "Success update learning material",
	})
}
func (controller *LearningMaterialControllerImpl) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "id must not be empty",
		})
	}
	err := controller.LearningMaterialService.Delete(id)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
		return ctx.Status(fiberErr.Code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success delete learning material",
	})
}
