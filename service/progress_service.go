package service

import "master-proof-api/dto"

type ProgressService interface {
	GetDashboardMenu(userId string) ([]*dto.MenuDashboardResponse, error)
}
