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

type ActivityServiceImpl struct {
	ActivityRepository repository.ActivityRepository
}

func NewActivityService(activityRepository repository.ActivityRepository) ActivityService {
	return &ActivityServiceImpl{
		ActivityRepository: activityRepository,
	}
}

//func (service *ActivityServiceImpl) CreateActivity(request *dto.CreateActivityRequest) error {
//	err := godotenv.Load(".env")
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	file, err := request.File.Open()
//	if err != nil {
//		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
//	}
//	bytes, err := io.ReadAll(file)
//	//var buffer bytes2.Buffer
//	//writer := multipart.NewWriter(&buffer)
//	//field, err := writer.CreateFormField("file")
//	//if err != nil {
//	//	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
//	//}
//	//test, err := io.Copy(field, file)
//	args := fiber.AcquireArgs()
//	args.Add("fileName", request.Name+".pdf")
//	args.SetBytesV("file", bytes)
//	imagekitPrivateKey := os.Getenv("IMAGEKIT_PRIVATE_KEY")
//	fmt.Println("IMAGEKIT_PRIVATE_KEY:", imagekitPrivateKey)
//	url := "https://upload.imagekit.io/api/v1/files/upload"
//	fmt.Println(imagekitPrivateKey)
//	agent := fiber.Post(url)
//	agent.ContentType("multipart/form-data")
//	agent.BasicAuth(imagekitPrivateKey, "")
//	agent.MultipartForm(args)
//	_, responseBody, errors := agent.Bytes()
//	if len(errors) > 0 {
//		return errors[0]
//	}
//	fmt.Println(string(responseBody))
//	var result struct {
//		FileId string `json:"fileId"`
//		Name   string `json:"name"`
//		Size   int    `json:"size"`
//		Url    string `json:"url"`
//	}
//	err = json.Unmarshal(responseBody, &result)
//	if err != nil {
//		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
//	}
//	req := model.Activity{
//		Id:        uuid.New().String(),
//		FileId:    result.FileId,
//		Name:      request.Name,
//		CreatedAt: time.Time{},
//		File: model.File{
//			ID:  result.FileId,
//			Url: result.Url,
//		},
//	}
//	err = service.ActivityRepository.CreateActivity(&req)
//	if err != nil {
//		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
//	}
//	return nil}

func (service *ActivityServiceImpl) CreateActivity(request *dto.CreateActivityRequest) error {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
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
