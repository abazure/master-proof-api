package controller

import "github.com/gofiber/fiber/v2"

type ProgressController interface {
	GetMenuDashboard(ctx *fiber.Ctx) error
	GetUserProgress(ctx *fiber.Ctx) error
}
