package models

import "github.com/go-playground/validator/v10"

type Blog struct {
	ID        int    `json:"id"`
	Title     string `json:"title" validate:"required"`           // Title is required
	Content   string `json:"content" validate:"required"`         // Content is required
	Thumbnail string `json:"thumbnail" validate:"required,url"`   // Thumbnail is required and must be a valid URL
	UserID    int    `json:"user_id" validate:"required,numeric"` // UserID is required and must be numeric
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Validate performs validation on the Blog struct
func (b *Blog) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
