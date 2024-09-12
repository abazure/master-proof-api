package route

import (
	"github.com/gofiber/fiber/v2"
	"master-proof-api/controller"
	"master-proof-api/middleware"
)

func SetupRoute(app *fiber.App, userController controller.UserController, learningMaterialController controller.LearningMaterialController, quizController controller.QuizController) {
	api := app.Group("/api")

	//USERS
	api.Post("/users", userController.Create)
	api.Get("users/profile", middleware.FirebaseAuthMiddleware(), userController.Find)
	api.Post("users/login", userController.Login)
	api.Post("/users/reset-password", userController.ResetPassword)
	api.Get("/students", middleware.FirebaseAuthMiddleware(), userController.FindByRole)

	//LEARNING_MATERIAL
	api.Get("/learning-materials", middleware.FirebaseAuthMiddleware(), learningMaterialController.FindAll)

	//Quiz
	api.Get("/quizzes/:name", middleware.FirebaseAuthMiddleware(), quizController.FindQuizWithCorrectAnswer)
	api.Get("/quizzes/without/:name", middleware.FirebaseAuthMiddleware(), quizController.FindQuizWithoutCorrectAnswer)
}
