package main

import "q3-blog-app/config"

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	config.ConnectDB()

	// Create the database if they do not exist
	config.CreateDatabase()

	// Create the tables if they do not exist
	config.CreateTables()
}
