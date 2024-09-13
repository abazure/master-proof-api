package model

import "time"

type DiagnosticReport struct {
	Name        string    `gorm:"column:name;primary_key"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Users       []User    `gorm:"many2many:user_diagnostic_reports;foreignKey:name;joinForeignKey:diagnostic_report_id;references:id;joinReferences:user_id"`
}

func (d *DiagnosticReport) TableName() string {
	return "diagnostic_reports"
}
