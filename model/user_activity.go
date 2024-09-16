package model

import "time"

type UserActivity struct {
	Id         string    `gorm:"primary_key:column:id"`
	UserId     string    `gorm:"column:user_id"`
	FileId     string    `gorm:"column:file_id"`
	ActivityId string    `gorm:"column:activity_id"`
	Comment    string    `gorm:"column:comment"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	File       File      `gorm:"foreignKey:file_id;references:id"`
	User       User      `gorm:"foreignKey:user_id;references:id"`
	Activity   Activity  `gorm:"foreignKey:activity_id;references:id"`
}

func (u *UserActivity) TableName() string {
	return "user_activities"
}
