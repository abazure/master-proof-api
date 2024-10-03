package dto

import "mime/multipart"

type LearningMaterialResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"ic_url"`
	Url         string `json:"pdf_url"`
}

type CreateLearningMaterialRequest struct {
	Title       string                `json:"title" validate:"required"`
	Description string                `json:"description" validate:"required"`
	File        *multipart.FileHeader `json:"file;" validate:"required"`
	FileName    string                `json:"file_name;" validate:"required"`
	Icon        *multipart.FileHeader `json:"icon" validate:"required"`
	IconName    string                `json:"icon_name;" validate:"required"`
}
type UpdateLearningMaterialRequest struct {
	Id          string                `json:"id" validate:"required"`
	Title       string                `json:"title" validate:"required"`
	Description string                `json:"description" validate:"required"`
	File        *multipart.FileHeader `json:"file;" validate:"required"`
	FileName    string                `json:"file_name;" validate:"required"`
	Icon        *multipart.FileHeader `json:"icon" validate:"required"`
	IconName    string                `json:"icon_name;" validate:"required"`
}

type UserSaveProgressRequest struct {
	UserID             string `json:"user_id" validate:"required"`
	LearningMaterialId string `json:"learning_material_id" validate:"required"`
	IsFinished         bool   `json:"is_finished"`
}
