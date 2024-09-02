package route

import (
	"github.com/gofiber/fiber/v2"
	"master-proof-api/controller"
	"master-proof-api/middleware"
)

func SetupRoute(app *fiber.App, controller controller.UserController) {
	api := app.Group("/api")

	api.Post("/users", controller.Create)
	api.Get("users/profile", middleware.FirebaseAuthMiddleware(), controller.Find)
	api.Post("users/login", controller.Login)
}
