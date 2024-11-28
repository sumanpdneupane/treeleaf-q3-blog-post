package services

import (
	"q3-blog-app/config"
	"q3-blog-app/models"
)

func CreateBlog(blog *models.Blog) error {
	_, err := config.DB.Exec("INSERT INTO Blogs (title, content, thumbnail_url, user_id) VALUES (?, ?, ?, ?)",
		blog.Title, blog.Content, blog.Thumbnail, blog.UserID,
	)
	return err
}
