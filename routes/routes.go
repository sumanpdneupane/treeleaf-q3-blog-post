package routes

import (
	"q3-blog-app/controllers"
	"q3-blog-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//Auth routes
	auth := app.Group("/auth")
	auth.Post("/register", controllers.RegisterUser)
	auth.Post("/login", controllers.LoginUser)

	// Blog routes (protected by auth middleware)
	blogGroup := app.Group("/blogs", middleware.AuthMiddleware)
	blogGroup.Post("/", controllers.CreateBlog)
	blogGroup.Get("/:id", controllers.GetBlog)
	blogGroup.Put("/:id", controllers.UpdateBlog)
	blogGroup.Delete("/:id", controllers.DeleteBlog)
}
