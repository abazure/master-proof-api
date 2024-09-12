package model

import "time"

type Quiz struct {
	ID             string `gorm:"primary_key;column:id"`
	QuizCategoryId string `gorm:"column:quiz_category_id"`
	//QuizCategory   QuizCategory `gorm:"foreignKey:QuizCategoryId;references:id"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	CreatedAt   time.Time  `gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Questions   []Question `gorm:"foreignKey:quiz_id;references:id"`
}

func (q *Quiz) TableName() string {
	return "quizzes"
}
