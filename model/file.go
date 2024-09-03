package model

import "time"

type File struct {
	ID        string    `gorm:"primary_key;column:id"`
	Url       string    `gorm:"column:url;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (f *File) TableName() string {
	return "files"
}
