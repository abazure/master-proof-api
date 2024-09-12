package dto

type LearningMaterialResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Url         string `json:"pdf_url"`
}
