package utils

import "github.com/gofiber/fiber/v2"

func HandleErrorResponse(context *fiber.Ctx, statusCode int, message string) error {
	return context.Status(statusCode).JSON(&fiber.Map{"message": message})
}
