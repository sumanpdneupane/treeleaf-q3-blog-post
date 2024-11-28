package main

import (
	"log"
	"q3-blog-app/config"
	"q3-blog-app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	config.ConnectDB()

	// Create the database if they do not exist
	config.CreateDatabase()

	// Create the tables if they do not exist
	config.CreateTables()

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
