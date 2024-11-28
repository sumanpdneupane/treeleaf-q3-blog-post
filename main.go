package main

import "q3-blog-app/config"

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	config.ConnectDB()
}
