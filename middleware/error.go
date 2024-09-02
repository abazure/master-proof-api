package middleware

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Continue the middleware chain, catching any errors
		err := c.Next()
		if err != nil {
			var validationErrors validator.ValidationErrors
			if errors.As(err, &validationErrors) {
				// Handle validation errors
				validationErrorMessages := make(map[string]string)
				for _, validationError := range validationErrors {
					validationErrorMessages[validationError.Field()] = validationError.Error()
				}
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": validationErrorMessages,
				})
			}

			// For other types of errors, return a generic 500 error
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// If no error, proceed as usual
		return nil
	}
}
