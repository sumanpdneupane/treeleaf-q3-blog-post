package controllers

import (
	"q3-blog-app/middleware"
	"q3-blog-app/models"
	"q3-blog-app/services"
	"q3-blog-app/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginData); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Invalid input")
	}
	user, err := services.GetUserByUsername(loginData.Username)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Database error")
	}

	if user == nil {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Generate JWT using the utility function
	token, err := middleware.GenerateJWT(user.ID, user.Role, 0) // Default expiration of 72 hours
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Error generating token")
	}

	return utils.RespondWithJSON(c, fiber.StatusOK, fiber.Map{"token": token})
}
