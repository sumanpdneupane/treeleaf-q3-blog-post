package controllers

import (
	"q3-blog-app/models"
	"q3-blog-app/services"
	"q3-blog-app/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateBlog(c *fiber.Ctx) error {
	var blog models.Blog
	if err := c.BodyParser(&blog); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Invalid input"+err.Error())
	}

	if err := services.CreateBlog(&blog); err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Error saving blog")
	}

	return utils.RespondWithJSON(c, fiber.StatusCreated, fiber.Map{"message": "Blog created successfully"})
}
