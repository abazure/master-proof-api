package service

import "master-proof-api/dto"

type ActivityService interface {
	CreateActivity(request *dto.CreateActivityRequest) error
	FindAll() ([]*dto.FindAllActivityResponse, error)
	FindById(id string) (*dto.FindAllActivityResponse, error)
}
