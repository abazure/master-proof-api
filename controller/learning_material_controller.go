package controller

import "github.com/gofiber/fiber/v2"

type LearningMaterialController interface {
	FindAll(ctx *fiber.Ctx) error
}
