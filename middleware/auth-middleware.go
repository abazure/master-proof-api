package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"master-proof-api/config"
	"strings"
)

var (
	app        = config.InitializeFirebase()
	authClient = config.FirebaseAuthInitialize(app)
)

func FirebaseAuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// Get the Authorization header
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is missing",
			})
		}

		// Split the header to extract the token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid Authorization header format",
			})
		}

		// Get the token part
		idToken := parts[1]

		// Verify the token
		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Store the token in the context, you can use this in your handler
		ctx.Locals("user", token)

		// Continue to the next handler
		return ctx.Next()
	}
}
