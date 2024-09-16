package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"master-proof-api/dto"
	"master-proof-api/service"
)

type ActivityControllerImpl struct {
	ActivityService service.ActivityService
}

func NewActivityController(activityService service.ActivityService) ActivityController {
	return &ActivityControllerImpl{
		ActivityService: activityService,
	}
}
func (controller *ActivityControllerImpl) CreateActivity(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	name := ctx.FormValue("name")
	request := dto.CreateActivityRequest{
		File: file,
		Name: name,
	}

	err = controller.ActivityService.CreateActivity(&request)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			if fiberErr.Code == 409 {
				return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
					"errors": err.Error(),
				})
			}
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
		return err
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success upload file",
	})

	return nil
}
func (controller *ActivityControllerImpl) FindAllActivity(ctx *fiber.Ctx) error {
	responses, err := controller.ActivityService.FindAll()
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": map[string]interface{}{
			"activities": responses,
		},
	})
	return nil
}
