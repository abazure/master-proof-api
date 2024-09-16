package dto

import "mime/multipart"

type CreateActivityRequest struct {
	File *multipart.FileHeader `json:"file;" validate:"required"`
	Name string                `json:"name;" validate:"required"`
}
type FindAllActivityResponse struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	PdfUrl string `json:"pdf_url"`
}
type CreateActivitySubmissionRequest struct {
	UserId     string                `json:"user_id"`
	ActivityId string                `json:"activity_id"`
	File       *multipart.FileHeader `json:"file;" validate:"required"`
}
type UpdateCommentRequest struct {
	UserId     string `json:"user_id" validate:"required"`
	ActivityId string `json:"activity_id" validate:"required"`
	Comment    string `json:"comment" validate:"required"`
}
