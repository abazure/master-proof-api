package controller

import (
	"errors"
	"firebase.google.com/go/v4/auth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"master-proof-api/dto"
	"master-proof-api/service"
)

type UserControllerImpl struct {
	UserService service.UserService
	Validate    *validator.Validate
}

func NewUserController(userService service.UserService, validate *validator.Validate) UserController {
	return &UserControllerImpl{UserService: userService, Validate: validate}
}

func (controller *UserControllerImpl) Create(ctx *fiber.Ctx) error {

	var user dto.UserCreateRequest
	err := ctx.BodyParser(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	err = controller.Validate.Struct(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	err2 := controller.UserService.Create(user)
	if err2 != nil {
		var fiberErr *fiber.Error
		if errors.As(err2, &fiberErr) {
			if fiberErr.Code == 409 {
				return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
					"errors": err2.Error(),
				})
			}
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err2.Error(),
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err2.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message": "Register new user success",
	})
}

func (controller *UserControllerImpl) Login(ctx *fiber.Ctx) error {
	var user dto.UserLoginRequest
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	err = controller.Validate.Struct(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	response, err := controller.UserService.Login(user)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) && fiberErr.Code == 401 {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errors": err.Error(),
			})
		} else {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}
	}
	return ctx.Status(200).JSON(fiber.Map{
		"data": response,
	})
}

func (controller *UserControllerImpl) Find(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*auth.Token)
	email := token.Claims["email"].(string)
	response, err := controller.UserService.FindById(email, "")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Code:    500,
			Message: "Internal Server Error",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}

func (controller *UserControllerImpl) ResetPassword(ctx *fiber.Ctx) error {
	var userRequest dto.UserPasswordResetRequest
	_ = ctx.BodyParser(&userRequest)
	err := controller.Validate.Struct(userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	err = controller.UserService.ResetPassword(userRequest.Email)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": fiberErr.Error(),
			})
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Please click link in your email for reset password",
	})

}

func (controller *UserControllerImpl) FindByRole(ctx *fiber.Ctx) error {
	role := "STUDENT"
	result, err := controller.UserService.FindByRole(role)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": map[string]interface{}{
			"students": result,
		},
	})
}

func (controller *UserControllerImpl) UpdatePhotoProfile(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	token := ctx.Locals("user").(*auth.Token)
	id := token.Claims["user_id"].(string)

	request := dto.UpdateUserPhotoRequest{
		Id:    id,
		Photo: file,
	}
	err = controller.Validate.Struct(request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
	err = controller.UserService.UpdatePhotoProfile(&request)
	if err != nil {
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return ctx.Status(fiberErr.Code).JSON(fiber.Map{
				"errors": fiberErr.Error(),
			})
		}
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully update photo profile",
	})
}

func (controller *UserControllerImpl) FindAllTeacher(ctx *fiber.Ctx) error {
	role := "TEACHER"
	result, err := controller.UserService.FindAllTeacher(role)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": map[string]interface{}{
			"teachers": result,
		},
	})
}
