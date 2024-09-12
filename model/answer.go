package model

import "time"

type Answer struct {
	ID         string `gorm:"primary_key;column:id"`
	QuestionId string `gorm:"column:question_id"`
	//Question   *Question `gorm:"foreignKey:QuestionId;references:Id"`
	Value     int8      `gorm:"column:value"`
	Text      string    `gorm:"column:text"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (a *Answer) TableName() string {
	return "answers"
}
