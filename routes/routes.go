package routes

import (
	"q3-blog-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//Auth routes
	auth := app.Group("/auth")
	auth.Post("/register", controllers.RegisterUser)
	auth.Post("/login", controllers.LoginUser)

	// Blog routes (protected by auth middleware)
	blogGroup := app.Group("/blogs")
	blogGroup.Post("/", controllers.CreateBlog)
}
