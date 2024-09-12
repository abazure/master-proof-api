package model

import "time"

type LearningMaterial struct {
	ID          string                   `gorm:"primary_key;column:id"`
	FileId      string                   `gorm:"column:file_id;not null;unique"`
	CategoryId  string                   `gorm:"column:learning_material_category_id;not null;unique"`
	Title       string                   `gorm:"column:title;not null"`
	Description string                   `gorm:"column:description;not null"`
	CreatedAt   time.Time                `gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt   time.Time                `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	File        File                     `gorm:"foreignKey:file_id;references:id"`
	Category    LearningMaterialCategory `gorm:"foreignKey:learning_material_category_id;references:id"`
}

func (l *LearningMaterial) TableName() string {
	return "learning_materials"
}
