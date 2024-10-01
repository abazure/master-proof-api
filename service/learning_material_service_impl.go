package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"io"
	"master-proof-api/dto"
	"master-proof-api/model"
	"master-proof-api/repository"
	"mime/multipart"
	"net/http"
	"os"
)

type LearningMaterialServiceImpl struct {
	LearningMaterialRepository repository.LearningMaterialRepository
}

func NewLearningMaterialService(learningMaterialRepository repository.LearningMaterialRepository) LearningMaterialService {
	return &LearningMaterialServiceImpl{
		LearningMaterialRepository: learningMaterialRepository,
	}
}

func (service *LearningMaterialServiceImpl) FindAll() []*dto.LearningMaterialResponse {
	responses, err := service.LearningMaterialRepository.FindAll()
	if err != nil {
		return nil
	}
	var learningMaterials []*dto.LearningMaterialResponse
	for _, learningMaterial := range responses {
		learningMaterials = append(learningMaterials, &dto.LearningMaterialResponse{
			ID:          learningMaterial.ID,
			Title:       learningMaterial.Title,
			Description: learningMaterial.Description,
			Url:         learningMaterial.File.Url,
			Icon:        learningMaterial.Icon.IcUrl,
		})
	}
	return learningMaterials
}

func (service *LearningMaterialServiceImpl) Create(request *dto.CreateLearningMaterialRequest) error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	// Validate file types
	if err := validateFileType(request.File, "application/pdf"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid PDF file: "+err.Error())
	}
	if err := validateFileType(request.Icon, "image/png"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid icon file: "+err.Error())
	}

	// Upload files
	pdfResult, err := uploadFile(request.File, request.FileName+".pdf", "PDF")
	if err != nil {
		return err
	}

	iconResult, err := uploadFile(request.Icon, request.IconName+".png", "Icon")
	if err != nil {
		return err
	}

	// Create learning material
	learningMaterial := model.LearningMaterial{
		ID:          uuid.New().String(),
		FileId:      pdfResult.FileId,
		IconId:      iconResult.FileId,
		Title:       request.Title,
		Description: request.Description,
		File: model.File{
			ID:  pdfResult.FileId,
			Url: pdfResult.Url,
		},
		Icon: model.Icon{
			Id:    iconResult.FileId,
			IcUrl: iconResult.Url,
		},
	}

	if err := service.LearningMaterialRepository.Create(&learningMaterial); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create learning material: "+err.Error())
	}

	return nil
}

func validateFileType(file *multipart.FileHeader, expectedType string) error {
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	buff := make([]byte, 512)
	_, err = f.Read(buff)
	if err != nil {
		return err
	}

	fileType := http.DetectContentType(buff)
	if fileType != expectedType {
		return fmt.Errorf("expected %s, got %s", expectedType, fileType)
	}

	return nil
}

func uploadFile(file *multipart.FileHeader, filename string, fileType string) (*struct {
	FileId string `json:"fileId"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
	Url    string `json:"url"`
}, error) {
	imagekitPrivateKey := os.Getenv("IMAGEKIT_PRIVATE_KEY")
	url := "https://upload.imagekit.io/api/v1/files/upload"

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// Add fileName field
	if err := writer.WriteField("fileName", filename); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to add fileName field: "+err.Error())
	}

	// Add file field
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create form file: "+err.Error())
	}

	f, err := file.Open()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to open file: "+err.Error())
	}
	defer f.Close()

	if _, err = io.Copy(part, f); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to copy file: "+err.Error())
	}

	writer.Close()

	req, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create request: "+err.Error())
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(imagekitPrivateKey, "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to send request: "+err.Error())
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to read response: "+err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fiber.NewError(resp.StatusCode, fmt.Sprintf("ImageKit API error for %s: %s", fileType, string(responseBody)))
	}

	var result struct {
		FileId string `json:"fileId"`
		Name   string `json:"name"`
		Size   int    `json:"size"`
		Url    string `json:"url"`
	}

	if err = json.Unmarshal(responseBody, &result); err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to parse response: "+err.Error())
	}

	return &result, nil
}

func (service *LearningMaterialServiceImpl) FindById(learningMaterialId string) (*dto.LearningMaterialResponse, error) {
	if learningMaterialId == "" {
		return &dto.LearningMaterialResponse{}, fiber.NewError(fiber.StatusBadRequest, "Failed to find learningMaterialId")
	}
	result, _ := service.LearningMaterialRepository.FindById(learningMaterialId)
	if result == nil {
		return &dto.LearningMaterialResponse{}, fiber.NewError(fiber.StatusNotFound, "Learning material not found")
	}
	learningMaterial := dto.LearningMaterialResponse{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Icon:        result.Icon.IcUrl,
		Url:         result.File.Url,
	}

	return &learningMaterial, nil

}

func (service *LearningMaterialServiceImpl) UpdateProgress(request *dto.UserSaveProgressRequest) error {
	record, _ := service.LearningMaterialRepository.FindById(request.LearningMaterialId)
	if record == nil {
		return fiber.NewError(fiber.StatusNotFound, "Learning material not found")
	}
	userRequest := model.LearningMaterialProgress{
		ID:                 uuid.New().String(),
		UserID:             request.UserID,
		LearningMaterialId: request.LearningMaterialId,
		IsFinished:         true,
	}
	err := service.LearningMaterialRepository.SaveProgress(&userRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
func (service *LearningMaterialServiceImpl) UpdateLearningMaterial(request *dto.UpdateLearningMaterialRequest) error {

	record, err2 := service.LearningMaterialRepository.FindById(request.Id)
	if err2 != nil {
		return fiber.NewError(fiber.StatusNotFound, "Learning material not found")
	}

	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	// Validate file types
	if err := validateFileType(request.File, "application/pdf"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid PDF file: "+err.Error())
	}
	if err := validateFileType(request.Icon, "image/png"); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid icon file: "+err.Error())
	}

	// Upload files
	pdfResult, err := uploadFile(request.File, request.FileName+".pdf", "PDF")
	if err != nil {
		return err
	}

	iconResult, err := uploadFile(request.Icon, request.IconName+".png", "Icon")
	if err != nil {
		return err
	}
	if pdfResult.FileId != "" {
		record.FileId = pdfResult.FileId
		record.File.ID = pdfResult.FileId
		record.File.Url = pdfResult.Url

	}

	if iconResult.FileId != "" {
		record.IconId = iconResult.FileId
		record.Icon.Id = iconResult.FileId
		record.Icon.IcUrl = iconResult.Url

	}

	if request.Title != "" {
		record.Title = request.Title
	}
	if request.Description != "" {
		record.Description = request.Description
	}
	err = service.LearningMaterialRepository.CreateFile(&model.File{
		ID:  pdfResult.FileId,
		Url: pdfResult.Url,
	})
	err = service.LearningMaterialRepository.CreateIcon(&model.Icon{
		Id:    iconResult.FileId,
		IcUrl: iconResult.Url,
	})

	// Create learning material
	learningMaterial := model.LearningMaterial{
		FileId:      record.FileId,
		IconId:      record.IconId,
		Title:       record.Title,
		Description: record.Description,
		File: model.File{
			ID:  record.File.ID,
			Url: record.File.Url,
		},
		Icon: model.Icon{
			Id:    record.Icon.Id,
			IcUrl: record.Icon.IcUrl,
		},
	}
	err = service.LearningMaterialRepository.Update(&learningMaterial, request.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
