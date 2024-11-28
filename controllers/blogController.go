package controllers

import (
	"fmt"
	"q3-blog-app/models"
	"q3-blog-app/services"
	"q3-blog-app/utils"
	"strconv"
	"time"

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

func GetBlog(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Invalid blog ID")
	}

	blog, err := services.GetBlogByID(id)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Error fetching blog")
	}
	if blog == nil {
		return utils.RespondWithError(c, fiber.StatusNotFound, "Blog not found")
	}

	return utils.RespondWithJSON(c, fiber.StatusOK, blog)
}

func UpdateBlog(c *fiber.Ctx) error {
	// Get blog ID from URL
	id := c.Params("id")

	var updatedBlog models.Blog
	if err := c.BodyParser(&updatedBlog); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, fmt.Sprintf("Error parsing blog: %s", err.Error()))
	}

	updatedBlog.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := services.UpdateBlogByID(id, updatedBlog)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Failed to update the blog "+err.Error())
	}
	b_id, _ := strconv.Atoi(id)
	blog, _ := services.GetBlogByID(b_id)

	return utils.RespondWithJSON(c, fiber.StatusOK, blog)
}

func DeleteBlog(c *fiber.Ctx) error {
	// Get blog ID from URL
	blogID := c.Params("id")

	// Call service to delete the blog from the database
	err := services.DeleteBlogByID(blogID)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Failed to delete the blog "+err.Error())
	}

	return utils.RespondWithJSON(c, fiber.StatusOK, fiber.Map{
		"message": "Blog deleted successfully",
	})
}
