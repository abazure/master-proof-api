package controller

import (
	"errors"
	"firebase.google.com/go/v4/auth"
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
func (controller *ActivityControllerImpl) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := controller.ActivityService.FindById(id)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})

		}
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
	return nil
}
func (controller *ActivityControllerImpl) CreateActivitySubmission(ctx *fiber.Ctx) error {
	file, err2 := ctx.FormFile("file")
	if err2 != nil {
		return err2
	}
	token := ctx.Locals("user").(*auth.Token)
	UserId := token.Claims["user_id"].(string)
	activityId := ctx.Params("id")
	request := dto.CreateActivitySubmissionRequest{
		UserId:     UserId,
		ActivityId: activityId,
		File:       file,
	}
	err := controller.ActivityService.CreateActivitySubmission(&request)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success upload file",
	})
	return nil
}
func (controller *ActivityControllerImpl) UpdateComment(ctx *fiber.Ctx) error {
	var request *dto.UpdateCommentRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	id := ctx.Params("id")
	request.Id = id
	err = controller.ActivityService.UpdateCommentUserActivity(request)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success upload review",
	})
	return nil

}
func (controller *ActivityControllerImpl) FindAllUserActivity(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	response, err := controller.ActivityService.FindAllUserActivityById(userId)
	if response == nil {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": []dto.FindAllUserActivity{},
		})
	}
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
	return nil
}
func (controller *ActivityControllerImpl) FindOneAllUserActivity(ctx *fiber.Ctx) error {
	//userId := ctx.Params("userId")
	id := ctx.Params("id")
	response, err := controller.ActivityService.FindOneUserActivityById(id)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	if response == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "User Activity Not Found",
		})
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
	return nil
}

func (controller *ActivityControllerImpl) FindAllUserActivityForStudent(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*auth.Token)
	userId := token.Claims["user_id"].(string)
	response, err := controller.ActivityService.FindAllUserActivityById(userId)
	if response == nil {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": []dto.FindAllUserActivity{},
		})
	}
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
	return nil
}
