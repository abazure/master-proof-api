package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
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

type ActivityServiceImpl struct {
	ActivityRepository repository.ActivityRepository
	Validate           *validator.Validate
}

func NewActivityService(activityRepository repository.ActivityRepository, validate *validator.Validate) ActivityService {
	return &ActivityServiceImpl{
		ActivityRepository: activityRepository,
		Validate:           validate,
	}
}

func (service *ActivityServiceImpl) CreateActivity(request *dto.CreateActivityRequest) error {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	err2 := service.Validate.Struct(request)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	file, err := request.File.Open()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to open file: "+err.Error())
	}
	defer file.Close()

	var buffer bytes.Buffer

	writer := multipart.NewWriter(&buffer)

	if err := writer.WriteField("fileName", request.Name+".pdf"); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to add fileName field: "+err.Error())
	}

	part, err := writer.CreateFormFile("file", request.Name+".pdf")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create form file: "+err.Error())
	}

	if _, err = io.Copy(part, file); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to copy file: "+err.Error())
	}

	writer.Close()

	imagekitPrivateKey := os.Getenv("IMAGEKIT_PRIVATE_KEY")
	url := "https://upload.imagekit.io/api/v1/files/upload"
	req, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create request: "+err.Error())
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(imagekitPrivateKey, "")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to send request: "+err.Error())
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to read response: "+err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return fiber.NewError(resp.StatusCode, "ImageKit API error: "+string(responseBody))
	}

	var result struct {
		FileId string `json:"fileId"`
		Name   string `json:"name"`
		Size   int    `json:"size"`
		Url    string `json:"url"`
	}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to parse response: "+err.Error())
	}

	activity := model.Activity{
		Id:     uuid.New().String(),
		FileId: result.FileId,
		Name:   request.Name,
		File: model.File{
			ID:  result.FileId,
			Url: result.Url,
		},
	}

	if err := service.ActivityRepository.CreateActivity(&activity); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create activity: "+err.Error())
	}

	return nil
}

func (service *ActivityServiceImpl) FindAll() ([]*dto.FindAllActivityResponse, error) {
	result, err := service.ActivityRepository.FindAll()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if result == nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "No result")
	}
	var response []*dto.FindAllActivityResponse
	for _, activity := range result {
		var data = dto.FindAllActivityResponse{
			Id:     activity.Id,
			Title:  activity.Name,
			PdfUrl: activity.File.Url,
		}
		response = append(response, &data)
	}
	return response, nil
}
func (service *ActivityServiceImpl) FindById(id string) (*dto.FindAllActivityResponse, error) {
	result, err := service.ActivityRepository.FindById(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	if result == nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "No result")
	}
	var data = dto.FindAllActivityResponse{
		Id:     result.Id,
		Title:  result.Name,
		PdfUrl: result.File.Url,
	}
	return &data, nil
}
func (service *ActivityServiceImpl) CreateActivitySubmission(request *dto.CreateActivitySubmissionRequest) error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	err2 := service.Validate.Struct(request)
	if err2 != nil {
		return fiber.NewError(fiber.StatusBadRequest, err2.Error())
	}
	activity, err := service.ActivityRepository.FindById(request.ActivityId)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	if activity == nil {
		return fiber.NewError(fiber.StatusNotFound, "No result")
	}

	file, err := request.File.Open()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to open file: "+err.Error())
	}
	defer file.Close()
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	if err := writer.WriteField("fileName", request.UserId+".pdf"); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to add fileName field: "+err.Error())
	}

	part, err := writer.CreateFormFile("file", request.UserId+".pdf")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create form file: "+err.Error())
	}

	if _, err = io.Copy(part, file); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to copy file: "+err.Error())
	}

	writer.Close()

	imagekitPrivateKey := os.Getenv("IMAGEKIT_PRIVATE_KEY")
	url := "https://upload.imagekit.io/api/v1/files/upload"
	req, err := http.NewRequest("POST", url, &buffer)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create request: "+err.Error())
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(imagekitPrivateKey, "")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to send request: "+err.Error())
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to read response: "+err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		return fiber.NewError(resp.StatusCode, "ImageKit API error: "+string(responseBody))
	}

	var result struct {
		FileId string `json:"fileId"`
		Name   string `json:"name"`
		Size   int    `json:"size"`
		Url    string `json:"url"`
	}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to parse response: "+err.Error())
	}

	userActivity := model.UserActivity{
		Id:         uuid.New().String(),
		UserId:     request.UserId,
		FileId:     result.FileId,
		ActivityId: request.ActivityId,
		File: model.File{
			ID:  result.FileId,
			Url: result.Url,
		},
	}
	err2 = service.ActivityRepository.CreateActivitySubmission(&userActivity)
	if err2 != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create activity submission: "+err.Error())
	}
	return nil

}
