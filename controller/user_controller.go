package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Create(ctx *fiber.Ctx) error
	Find(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	ResetPassword(ctx *fiber.Ctx) error
	FindByRole(ctx *fiber.Ctx) error
	UpdatePhotoProfile(ctx *fiber.Ctx) error
}
