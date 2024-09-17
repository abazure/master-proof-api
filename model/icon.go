package model

import "time"

type Icon struct {
	Id        string    `gorm:"primary_key;column:id"`
	IcUrl     string    `gorm:"column:ic_url"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (i *Icon) TableName() string {
	return "icons"
}
