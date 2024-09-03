package service

import (
	"context"
	"encoding/json"
	"errors"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"master-proof-api/dto"
	"master-proof-api/model"
	"master-proof-api/repository"
	"os"
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

	userRequest := model.User{
		ID:    userRecord.UID,
		NIM:   request.Nim,
		Name:  userRecord.DisplayName,
		Email: userRecord.Email,
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
	}
	return loginResponse, nil

}

func (service *UserServiceImpl) FindById(email string, nim string) (dto.UserResponse, error) {

	user, err := service.UserRepository.FindById(email, nim)
	if err != nil {
		return dto.UserResponse{}, err
	}
	userResponse := dto.UserResponse{
		Nim:   user.NIM,
		Name:  user.Name,
		Email: user.Email,
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
