package dto

import "mime/multipart"

type CreateActivityRequest struct {
	File *multipart.FileHeader `json:"file;" validate:"required"`
	Name string                `json:"name;" validate:"required"`
}
type UpdateActivityRequest struct {
	Id   string                `json:"id;" validate:"required"`
	File *multipart.FileHeader `json:"file;"`
	Name string                `json:"name;"`
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
	Id      string
	Comment string `json:"comment" validate:"required"`
}
type FindAllUserActivity struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	PdfUrl  string `json:"pdf_url"`
}
