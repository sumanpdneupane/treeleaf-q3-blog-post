package routes

import (
	"q3-blog-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//Auth routes
	app.Post("/register", controllers.RegisterUser)
	app.Post("/login", controllers.LoginUser)
}
