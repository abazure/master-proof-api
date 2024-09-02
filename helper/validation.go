package helper

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ValidationCheck(ctx *fiber.Ctx, err error) error {
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		// Collect all validation error messages into an array
		errorMessages := make([]map[string]string, 0, len(errs))
		for _, validationError := range errs {
			errorMessages = append(errorMessages, map[string]string{
				"field":   validationError.Field(),
				"message": validationError.Error(),
			})
		}
		// Return all error messages as a JSON response
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": errorMessages,
		})
	}
	return nil
}
