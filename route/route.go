package route

import (
	"github.com/gofiber/fiber/v2"
	"master-proof-api/controller"
	"master-proof-api/middleware"
)

func SetupRoute(app *fiber.App, userController controller.UserController, learningMaterialController controller.LearningMaterialController, quizController controller.QuizController, activityController controller.ActivityController, progressController controller.ProgressController) {
	api := app.Group("/api")

	//USERS
	api.Post("/users", userController.Create)
	api.Get("/users/profile", middleware.FirebaseAuthMiddleware(), userController.Find)
	api.Post("/users/login", userController.Login)
	api.Post("/users/reset-password", userController.ResetPassword)
	api.Get("/students", middleware.FirebaseAuthMiddleware(), userController.FindByRole)
	api.Get("/teachers", middleware.FirebaseAuthMiddleware(), userController.FindAllTeacher)
	api.Patch("/users/update-photo", middleware.FirebaseAuthMiddleware(), userController.UpdatePhotoProfile)
	api.Get("/users/activities", middleware.FirebaseAuthMiddleware(), activityController.FindAllUserActivityForStudent)

	//LEARNING_MATERIAL
	api.Get("/learning-materials", middleware.FirebaseAuthMiddleware(), learningMaterialController.FindAll)
	api.Post("/learning-materials/upload", middleware.FirebaseAuthMiddleware(), learningMaterialController.Create)
	api.Get("/learning-materials/:id", middleware.FirebaseAuthMiddleware(), learningMaterialController.FindByID)
	api.Post("/learning-materials/users/progress/:id", middleware.FirebaseAuthMiddleware(), learningMaterialController.SaveProgress)

	//Quiz
	api.Get("/quizzes/competences/:name", middleware.FirebaseAuthMiddleware(), quizController.FindQuizWithCorrectAnswer)
	api.Get("/quizzes/diagnostics/:name", middleware.FirebaseAuthMiddleware(), quizController.FindQuizWithoutCorrectAnswer)
	api.Post("/quizzes/diagnostics/:name", middleware.FirebaseAuthMiddleware(), quizController.CreateUserDiagnosticReport)
	api.Get("/reports/diagnostics/:name", middleware.FirebaseAuthMiddleware(), quizController.FindUserDiagnosticReport)
	api.Post("/quizzes/competences/:name", middleware.FirebaseAuthMiddleware(), quizController.CreateUserCompetenceReport)
	api.Get("/reports/competences/:name", middleware.FirebaseAuthMiddleware(), quizController.FindUserCompetenceReport)
	api.Get("/reports/diagnostics/:name/:userId", middleware.FirebaseAuthMiddleware(), quizController.FindUserDiagnosticReportForTeacher)

	//ACTIVITY
	api.Post("/activities/upload/", middleware.FirebaseAuthMiddleware(), activityController.CreateActivity)
	api.Get("/activities", middleware.FirebaseAuthMiddleware(), activityController.FindAllActivity)
	api.Get("/activities/:id", middleware.FirebaseAuthMiddleware(), activityController.FindById)
	api.Post("/activities/submission/:id", middleware.FirebaseAuthMiddleware(), activityController.CreateActivitySubmission)
	api.Post("/activities/review/:id", middleware.FirebaseAuthMiddleware(), activityController.UpdateComment)
	api.Get("/activities/students/:userId", middleware.FirebaseAuthMiddleware(), activityController.FindAllUserActivity)
	api.Get("/activities/students/answered/:id", middleware.FirebaseAuthMiddleware(), activityController.FindOneAllUserActivity)

	//Progress
	api.Get("/dashboard", middleware.FirebaseAuthMiddleware(), progressController.GetMenuDashboard)
}
