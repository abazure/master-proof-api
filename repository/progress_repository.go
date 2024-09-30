package repository

import "master-proof-api/dto"

type ProgressRepository interface {
	GetLearningMaterialData(userId string) (*dto.CountData, error)
	GetDiagnosticTestData(userId string) (*dto.CountData, error)
	GetActivityData(userId string) (*dto.CountData, error)
	GetCompetenceData(userId string) (*dto.CountData, error)
}
