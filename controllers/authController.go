package controllers

import (
	"q3-blog-app/models"
	"q3-blog-app/services"
	"q3-blog-app/utils"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Invalid input")
	}
	err := services.RegisterUser(&user)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Error saving user")
	}
	return utils.RespondWithJSON(c, fiber.StatusCreated, fiber.Map{"message": "User registered successfully"})

}

func LoginUser(c *fiber.Ctx) error {
	return nil
}
