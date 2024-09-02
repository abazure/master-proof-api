package model

import "time"

type User struct {
	ID        string    `gorm:"primary_key;column:id;not null" json:"id"`
	NIM       string    `gorm:"column:nim;not null;uniqueIndex" json:"nim"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}
