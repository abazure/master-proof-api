package main

import (
	"log"
	"master-proof-api/config"
	"master-proof-api/controller"
	"master-proof-api/database"
	"master-proof-api/repository"
	"master-proof-api/route"
	"master-proof-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db := database.OpenConnection()
	validate := validator.New()
	firebaseInitialize := config.InitializeFirebase()
	firebase := config.FirebaseAuthInitialize(firebaseInitialize)

	//USER
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, firebase)
	UserController := controller.NewUserController(userService, validate)

	//LEARNING_MATERIAL
	learningMaterialRepository := repository.NewLearningMaterialRepository(db)
	learningMaterialService := service.NewLearningMaterialService(learningMaterialRepository)
	learningMaterialController := controller.NewLearningMaterialController(learningMaterialService, validate)

	//Quiz
	quizRepository := repository.NewQuizRepository(db)
	quizService := service.NewQuizService(quizRepository)
	quizController := controller.NewQuizController(quizService)

	//Activity
	activityRepository := repository.NewActivityRepository(db)
	activityService := service.NewActivityService(activityRepository, validate)
	activityController := controller.NewActivityController(activityService)

	//Progress
	progressRepository := repository.NewProgressRepository(db)
	progressService := service.NewProgressService(progressRepository)
	progressController := controller.NewProgressController(progressService)

	app := fiber.New()
	app.Use(cors.New())
	route.SetupRoute(app, UserController, learningMaterialController, quizController, activityController, progressController)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen("127.0.0.1:3000"))
}
