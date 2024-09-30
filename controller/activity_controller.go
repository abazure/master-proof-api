package controller

import "github.com/gofiber/fiber/v2"

type ActivityController interface {
	CreateActivity(ctx *fiber.Ctx) error
	UpdateActivity(ctx *fiber.Ctx) error
	FindAllActivity(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	CreateActivitySubmission(ctx *fiber.Ctx) error
	UpdateComment(ctx *fiber.Ctx) error
	FindAllUserActivity(ctx *fiber.Ctx) error
	FindOneAllUserActivity(ctx *fiber.Ctx) error
	FindAllUserActivityForStudent(ctx *fiber.Ctx) error
	DeleteActivity(ctx *fiber.Ctx) error
}
