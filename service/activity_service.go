package service

import "master-proof-api/dto"

type ActivityService interface {
	CreateActivity(request *dto.CreateActivityRequest) error
	UpdateActivity(request *dto.UpdateActivityRequest) error
	FindAll() ([]*dto.FindAllActivityResponse, error)
	FindById(id string) (*dto.FindAllActivityResponse, error)
	CreateActivitySubmission(request *dto.CreateActivitySubmissionRequest) error
	UpdateCommentUserActivity(request *dto.UpdateCommentRequest) error
	FindAllUserActivityById(userId string) ([]*dto.FindAllUserActivity, error)
	FindOneUserActivityById(id string) (*dto.FindAllUserActivity, error)
	DeleteActivityById(id string) error
}
