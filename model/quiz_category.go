package model

import "time"

type QuizCategory struct {
	ID        string    `gorm:"primary_key;column:id;not null"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Quizzes   []Quiz    `gorm:"foreignKey:quiz_category_id;references:id"`
}

func (q *QuizCategory) TableName() string {
	return "quiz_categories"
}
