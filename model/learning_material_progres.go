package model

type LearningMaterialProgress struct {
	ID                 string `gorm:"primary_key;column:id"`
	UserID             string `gorm:"column:user_id"`
	LearningMaterialId string `gorm:"column:learning_material_id"`
	IsFinished         bool   `gorm:"column:is_finished"`
}

func (l *LearningMaterialProgress) TableName() string {
	return "learning_material_progress"
}
