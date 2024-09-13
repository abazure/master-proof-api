package model

import "time"

type UserDiagnosticReport struct {
	Id                 string    `gorm:"column:id;primary_key"`
	UserId             string    `gorm:"column:user_id"`
	QuizId             string    `gorm:"column:quiz_id"`
	DiagnosticReportId string    `gorm:"column:diagnostic_report_id"`
	CreatedAt          time.Time `gorm:"column:created_at;autoCreateTime;"`
	//
	//User             User             `gorm:"foreignKey:UserId"`
	DiagnosticReport DiagnosticReport `gorm:"foreignKey:DiagnosticReportId;references:name"`
}

func (u *UserDiagnosticReport) TableName() string {
	return "user_diagnostic_reports"
}
