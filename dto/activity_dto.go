package dto

import "mime/multipart"

type CreateActivityRequest struct {
	File *multipart.FileHeader `json:"file;" validate:"required"`
	Name string                `json:"name;" validate:"required"`
}
type ActivityResponse struct {
}
