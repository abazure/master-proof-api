package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"io"
	"log"
	"master-proof-api/dto"
	"master-proof-api/model"
	"master-proof-api/repository"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	FirebaseAuth   *auth.Client
}

func NewUserService(userRepository repository.UserRepository, firebaseAuth *auth.Client) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		FirebaseAuth:   firebaseAuth,
	}
}
func (service *UserServiceImpl) Create(request dto.UserCreateRequest) error {

	existingUser, _ := service.UserRepository.FindById(request.Email, request.Nim)

	if existingUser != nil {
		return fiber.NewError(fiber.StatusConflict, "user with this email or NIM already exists")
	}

	params := (&auth.UserToCreate{}).Email(request.Email).Password(request.Password).DisplayName(request.Name)

	ctx := context.Background()
	userRecord, err := service.FirebaseAuth.CreateUser(ctx, params)
	if err != nil {
		return err
	}

	if request.Role == "" {
		request.Role = "STUDENT"
	}

	photoUrl := fmt.Sprintf("https://ui-avatars.com/api/?size=128&background=0D8ABC&color=fff&name=%s", userRecord.DisplayName)
	userRequest := model.User{
		ID:       userRecord.UID,
		Role:     request.Role,
		NIM:      request.Nim,
		Name:     userRecord.DisplayName,
		Email:    userRecord.Email,
		PhotoUrl: photoUrl,
	}
	// Attempt to save the user in the database
	err = service.UserRepository.Save(&userRequest)
	if err != nil {
		_ = service.FirebaseAuth.DeleteUser(ctx, userRecord.UID)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}

func (service *UserServiceImpl) Login(request dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return dto.UserLoginResponse{}, errors.New("API_KEY environment variable not set")
	}

	url := fmt.Sprintf("https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=%s", apiKey)
	user, _ := service.UserRepository.FindById(request.Email, "")
	if user == nil {
		return dto.UserLoginResponse{}, fiber.NewError(fiber.StatusUnauthorized, "Email or Password wrong")
	}
	requestBody, err := json.Marshal(map[string]interface{}{
		"email":             request.Email,
		"password":          request.Password,
		"returnSecureToken": true,
	})
	if err != nil {
		return dto.UserLoginResponse{}, err
	}
	agent := fiber.Post(url)
	agent.Body(requestBody)
	agent.ContentType("application/json")
	_, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return dto.UserLoginResponse{}, errs[0]
	}
	var token struct {
		IDToken string `json:"idToken"`
	}
	err = json.Unmarshal(body, &token)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}
	if token.IDToken == "" {
		return dto.UserLoginResponse{}, fiber.NewError(fiber.StatusUnauthorized, "Email or Password wrong")
	}
	loginResponse := dto.UserLoginResponse{
		Token: token.IDToken,
		Role:  user.Role,
	}
	return loginResponse, nil

}

func (service *UserServiceImpl) FindById(email string, nim string) (dto.UserResponse, error) {

	user, err := service.UserRepository.FindById(email, nim)
	if err != nil {
		return dto.UserResponse{}, err
	}
	userResponse := dto.UserResponse{
		Nim:      user.NIM,
		Name:     user.Name,
		Role:     string(user.Role),
		Email:    user.Email,
		PhotoUrl: user.PhotoUrl,
	}
	return userResponse, nil
}
func (service *UserServiceImpl) ResetPassword(email string) error {
	ctx := context.Background()
	userRecord, _ := service.FirebaseAuth.GetUserByEmail(ctx, email)
	if userRecord == nil {
		return fiber.NewError(fiber.StatusNotFound, "User with this email don't exist")
	}
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return fiber.NewError(fiber.StatusInternalServerError, "API key environment variable not set")
	}

	url := fmt.Sprintf("https://www.googleapis.com/identitytoolkit/v3/relyingparty/getOobConfirmationCode?key=%s", apiKey)

	requestBody, err := json.Marshal(map[string]interface{}{
		"requestType": "PASSWORD_RESET",
		"email":       email,
	})
	agent := fiber.Post(url)
	agent.Body(requestBody)
	agent.ContentType("application/json")
	_, _, errs := agent.Bytes()
	if len(errs) > 0 {
		return fiber.NewError(fiber.StatusInternalServerError, errs[0].Error())
	}
	return nil
}

func (service *UserServiceImpl) FindByRole(role string) ([]*dto.GetUserByRole, error) {
	role = "STUDENT"
	datas, err := service.UserRepository.FindByRole(role)
	if err != nil {
		return []*dto.GetUserByRole{}, err
	}

	var result []*dto.GetUserByRole
	for _, data := range datas {
		resultDto := dto.GetUserByRole{
			Id:       data.ID,
			Name:     data.Name,
			PhotoUrl: data.PhotoUrl,
		}
		result = append(result, &resultDto)
	}

	return result, nil

}

func (service *UserServiceImpl) UpdatePhotoProfile(request *dto.UpdateUserPhotoRequest) error {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	file, err := request.Photo.Open()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to open file: "+err.Error())
	}
	defer file.Close()

	// Determine file extension
	fileExt := filepath.Ext(request.Photo.Filename)
	if fileExt == "" {
		return fiber.NewError(fiber.StatusBadRequest, "File must have an extension")
	}

	// Generate timestamp
	timestamp := time.Now().Format("20060102_150405")

	// Create filename with timestamp
	filename := fmt.Sprintf("%s_%s%s", request.Id, timestamp, fileExt)

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	if err = writer.WriteField("fileName", filename); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to add fileName field: "+err.Error())
	}

	part, err := writer.CreateFormFile("file", filename)
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

	if err = json.Unmarshal(responseBody, &result); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to parse response: "+err.Error())
	}

	err = service.UserRepository.UpdatePhotoProfile(request.Id, result.Url)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return nil

}

func (service *UserServiceImpl) FindAllTeacher(role string) ([]*dto.GetUserByRole, error) {
	role = "TEACHER"
	datas, err := service.UserRepository.FindByRole(role)
	if err != nil {
		return []*dto.GetUserByRole{}, err
	}

	var result []*dto.GetUserByRole
	for _, data := range datas {
		resultDto := dto.GetUserByRole{
			Id:       data.ID,
			Name:     data.Name,
			PhotoUrl: data.PhotoUrl,
		}
		result = append(result, &resultDto)
	}

	return result, nil
}
