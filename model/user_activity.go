package model

import "time"

type UserActivity struct {
	Id         string    `gorm:"primary_key:column:id"`
	UserId     string    `gorm:"column:user_id"`
	ActivityId string    `gorm:"column:activity_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}

func (u *UserActivity) TableName() string {
	return "user_activities"
}
