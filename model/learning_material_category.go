package model

import "time"

type LearningMaterialCategory struct {
	ID          string    `gorm:"primary_key;column:id"`
	Title       string    `gorm:"column:title;not null"`
	Description string    `gorm:"column:description;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (l *LearningMaterialCategory) TableName() string {
	return "learning_material_categories"
}
