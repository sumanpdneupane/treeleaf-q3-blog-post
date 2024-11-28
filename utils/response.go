package utils

import "github.com/gofiber/fiber/v2"

func RespondWithError(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(fiber.Map{"error": message})
}

func RespondWithJSON(c *fiber.Ctx, code int, payload interface{}) error {
	return c.Status(code).JSON(payload)
}
