package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"master-proof-api/config"
	"master-proof-api/controller"
	"master-proof-api/database"
	"master-proof-api/repository"
	"master-proof-api/route"
	"master-proof-api/service"
)

func main() {
	db := database.OpenConnection()
	validate := validator.New()
	firebaseInitialize := config.InitializeFirebase()
	firebase := config.FirebaseAuthInitialize(firebaseInitialize)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, firebase)
	UserController := controller.NewUserController(userService, validate)

	app := fiber.New()
	//app.Use(middleware.ErrorHandler(app))
	route.SetupRoute(app, UserController)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
