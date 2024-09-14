package model

import "time"

type UserCompetenceReports struct {
	Id        string    `gorm:"column:id;primary_key"`
	UserId    string    `gorm:"column:user_id"`
	QuizName  string    `gorm:"column:quiz_name"`
	Score     int       `gorm:"column:score"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;"`
}

func (u *UserCompetenceReports) TableName() string {
	return "user_competence_reports"
}
