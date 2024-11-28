package models

type Blog struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Thumbnail string `json:"thumbnail"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}
