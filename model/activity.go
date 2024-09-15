package model

import "time"

type Activity struct {
	Id              string    `gorm:"primary_key:column:id"`
	FileId          string    `gorm:"column:file_id"`
	Name            string    `gorm:"column:name"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	File            File      `gorm:"foreignKey:file_id;references:id"`
	ActivityByUsers []User    `gorm:"many2many:user_activities;foreignKey:id;joinForeignKey:activity_id;references:id;JoinReferences:user_id"`
}

func (a *Activity) TableName() string {
	return "activities"
}
