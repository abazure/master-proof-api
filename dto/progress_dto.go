package dto

type MenuDashboardResponse struct {
	Category          string `json:"category"`
	Endpoint          string `json:"endpoint"`
	Title             string `json:"title"`
	IcUrl             string `json:"ic_url"`
	Desc              string `json:"desc"`
	FinishedMaterials int    `json:"finished_materials"`
	TotalMaterials    int    `json:"total_materials"`
	IsLocked          bool   `json:"is_locked"`
}

type CountData struct {
	FinishedMaterials int `json:"finished_materials"`
	TotalMaterials    int `json:"total_materials"`
}
