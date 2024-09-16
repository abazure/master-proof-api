package main

import (
	"fmt"
	"master-proof-api/database"
	"master-proof-api/model"
	"testing"
)

func TestQuiz(t *testing.T) {
	db := database.OpenConnection()

	var quizz []*model.Quiz
	db.Model(&model.Quiz{}).Preload("Questions.Answers").Find(&quizz)
	for _, quiz := range quizz {
		fmt.Println(quiz)
	}

}
func TestReport(t *testing.T) {
	db := database.OpenConnection()
	createRequest := model.UserDiagnosticReport{ // Ensure this is a unique ID
		UserId:             "nQnrJsWDRAYvvOlcUDEYaivvTim1",         // Valid user ID
		QuizId:             "5875cd46-5d9e-416c-bb68-ae82501eaf7e", // Valid quiz ID
		DiagnosticReportId: "VISUAL",                               // Valid diagnostic report ID 		// Current timestamp
	}

	err := db.Create(&createRequest).Error
	if err != nil {
		fmt.Println(err)
	}

}

func TestUUID(t *testing.T) {
	db := database.OpenConnection()
	var result model.UserDiagnosticReport
	db.Model(model.UserDiagnosticReport{}).Preload("DiagnosticReport").Take(&result)
	fmt.Println(result)

}

func TestUserActivity(t *testing.T) {
	db := database.OpenConnection()
	userId := "PkPLzb6QeFTOJKrmnb1AqxxRPtH2"
	subQuery := db.Model(&model.UserActivity{}).
		Select("activity_id, MAX(created_at) as created_at").
		Where("user_id = ?", userId).
		Group("activity_id")

	var userActivities []model.UserActivity
	err := db.Joins("JOIN (?) AS subquery ON user_activities.activity_id = subquery.activity_id AND user_activities.created_at = subquery.created_at", subQuery).Preload("Activity").
		Where("user_id = ?", userId).
		Find(&userActivities).Error
	if err != nil {
		t.Fatal(err)
	}

	for _, activity := range userActivities {
		fmt.Println(activity)
	}
}

func TestName(t *testing.T) {
	db := database.OpenConnection()
	var userActivities []model.UserActivity
	db.Model(model.UserActivity{}).Preload("Activity").Find(&userActivities)
	for _, activity := range userActivities {
		fmt.Println(activity)
	}
}
