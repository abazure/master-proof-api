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
