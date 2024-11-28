package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

var DB *sql.DB

func ConnectDB() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	//username:password@tcp(localhost:3306)/dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	log.Println("Database connected successfully!")
}

// Create the 'blog_post' database
func CreateDatabase() {
	// SQL query to create the 'blog_post' database if it doesn't exist
	createDBQuery := "CREATE DATABASE IF NOT EXISTS blog_post;"

	// Execute the query
	_, err := DB.Exec(createDBQuery)
	if err != nil {
		log.Fatal("Error creating 'blog_post' database: ", err)
	}

	fmt.Println("Database 'blog_post' created or already exists.")
}

func CreateTables() {
	// Create User table if it does not exist
	userTableQuery := `
		CREATE TABLE IF NOT EXISTS User (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(50) NOT NULL UNIQUE,
			email VARCHAR(100) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(50) NOT NULL
		);
	`
	_, err := DB.Exec(userTableQuery)
	if err != nil {
		log.Fatal("Error creating User table: ", err)
	}

	fmt.Println("User table created or already exists.")

}
