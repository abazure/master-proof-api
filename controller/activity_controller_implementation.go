package controller

import (
	"fmt"
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
	fmt.Println(&request)

	err = controller.ActivityService.CreateActivity(&request)
	if err != nil {
		return err
	}
	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success upload file",
	})

	return nil
}
