package controller

import "github.com/gofiber/fiber/v2"

type ActivityController interface {
	CreateActivity(ctx *fiber.Ctx) error
	FindAllActivity(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
}
