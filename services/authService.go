package services

import (
	"database/sql"
	"q3-blog-app/config"
	"q3-blog-app/models"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	_, err = config.DB.Exec("INSERT INTO Users (username, password, email, role) VALUES (?, ?, ?, ?)", user.Username, user.Password, user.Email, user.Role)
	return err
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := config.DB.QueryRow("SELECT id, username, email, password, role FROM Users WHERE username = ?", username).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return &user, err
}
