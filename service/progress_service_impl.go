package service

import (
	"master-proof-api/dto"
	"master-proof-api/repository"
)

type ProgressServiceImpl struct {
	ProgressRepository repository.ProgressRepository
}

func NewProgressService(progressRepository repository.ProgressRepository) ProgressService {
	return &ProgressServiceImpl{ProgressRepository: progressRepository}
}

func (service *ProgressServiceImpl) GetDashboardMenu(userId string) ([]*dto.MenuDashboardResponse, error) {
	diagnostic, _ := service.ProgressRepository.GetDiagnosticTestData(userId)
	statusDiagnostic := true
	if diagnostic.FinishedMaterials == diagnostic.TotalMaterials {
		statusDiagnostic = false
	}
	// Learning Material
	learningMaterialData, _ := service.ProgressRepository.GetLearningMaterialData(userId)
	statusLearningMaterial := true
	if learningMaterialData.FinishedMaterials == learningMaterialData.TotalMaterials {
		statusLearningMaterial = false
	}
	//Activity
	activityData, _ := service.ProgressRepository.GetActivityData(userId)
	statusActivity := true
	if activityData.FinishedMaterials == activityData.TotalMaterials {
		statusActivity = false
	}
	//Competence
	competenceData, _ := service.ProgressRepository.GetCompetenceData(userId)
	//statusCompetence := true
	//if competenceData.FinishedMaterials == competenceData.TotalMaterials {
	//	statusCompetence = false
	//}

	result := []*dto.MenuDashboardResponse{
		{
			Category:          "pre-test",
			Endpoint:          "diagnostic",
			Title:             "Diagnostic Test",
			IcUrl:             "https://ik.imagekit.io/q1qexvvey/diagnostic_ic.png?updatedAt=1725337675997",
			Desc:              "lorem ipsum dolor sit amet",
			FinishedMaterials: diagnostic.FinishedMaterials,
			TotalMaterials:    diagnostic.TotalMaterials,
			IsLocked:          false,
		},
		{
			Category:          "material",
			Endpoint:          "introduction-proof",
			Title:             "Introduction to proof",
			IcUrl:             "https://ik.imagekit.io/q1qexvvey/introduction_proof_ic.png?updatedAt=1725337679695",
			Desc:              "Pengenalan pembuktian geometri",
			FinishedMaterials: learningMaterialData.FinishedMaterials,
			TotalMaterials:    learningMaterialData.TotalMaterials,
			IsLocked:          statusDiagnostic,
		},
		{
			Category:          "material",
			Endpoint:          "understanding-proof",
			Title:             "Understanding of Proof Structure",
			IcUrl:             "https://ik.imagekit.io/q1qexvvey/understanding_proof_ic.png",
			Desc:              "Memahami struktur bukti geometri",
			FinishedMaterials: activityData.FinishedMaterials,
			TotalMaterials:    activityData.TotalMaterials,
			IsLocked:          statusLearningMaterial,
		},
		{
			Category:          "material",
			Endpoint:          "proof-competence-test",
			Title:             "Proof Competence Test",
			IcUrl:             "https://ik.imagekit.io/q1qexvvey/proof_competence_ic.png",
			Desc:              "Tes kompetensi pembuktian",
			FinishedMaterials: competenceData.FinishedMaterials,
			TotalMaterials:    competenceData.TotalMaterials,
			IsLocked:          statusActivity,
		},
	}
	return result, nil
}
