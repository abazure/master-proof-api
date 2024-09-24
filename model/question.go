package model

import "time"

type Question struct {
	ID     string `gorm:"primary_key;column:id"`
	QuizId string `gorm:"column:quiz_id"`
	//Quiz          *Quiz     `gorm:"foreignKey:QuizId;references:Id"`
	Question      string    `gorm:"column:question"`
	CorrectAnswer *int      `gorm:"column:correct_answer"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Answers       []Answer  `gorm:"foreignKey:question_id;references:id"`
}

func (q *Question) TableName() string {
	return "questions"
}
