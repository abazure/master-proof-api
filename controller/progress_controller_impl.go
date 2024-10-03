package controller

import (
	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
	"master-proof-api/service"
)

type ProgressControllerImpl struct {
	ProgressService service.ProgressService
}

func NewProgressController(progressService service.ProgressService) ProgressController {
	return &ProgressControllerImpl{
		ProgressService: progressService,
	}
}

func (controller *ProgressControllerImpl) GetMenuDashboard(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*auth.Token)
	userId := token.Claims["user_id"].(string)
	responses, err := controller.ProgressService.GetDashboardMenu(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"data": responses,
	})
}
func (controller *ProgressControllerImpl) GetUserProgress(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*auth.Token)
	userId := token.Claims["user_id"].(string)
	responses, err := controller.ProgressService.GetProgressPercentage(userId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": map[string]interface{}{
			"student_id": userId,
			"reports":    responses,
		},
	})

}
