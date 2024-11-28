package services

import (
	"database/sql"
	"fmt"
	"q3-blog-app/config"
	"q3-blog-app/models"
)

func CreateBlog(blog *models.Blog) error {
	_, err := config.DB.Exec("INSERT INTO Blogs (title, content, thumbnail_url, user_id) VALUES (?, ?, ?, ?)",
		blog.Title, blog.Content, blog.Thumbnail, blog.UserID,
	)
	return err
}

func GetBlogByID(id int) (*models.Blog, error) {
	var blog models.Blog
	err := config.DB.QueryRow("SELECT id, title, content, thumbnail_url, user_id, created_at, updated_at FROM blogs WHERE id = ?", id).
		Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Thumbnail, &blog.UserID, &blog.CreatedAt, &blog.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &blog, err
}

func UpdateBlogByID(blogID string, updatedBlog models.Blog) error {
	query := `UPDATE Blogs SET title = ?, content = ?, thumbnail_url = ?, updated_at = ? WHERE id = ?`
	_, err := config.DB.Exec(query, updatedBlog.Title, updatedBlog.Content, updatedBlog.Thumbnail, updatedBlog.UpdatedAt, blogID)
	if err != nil {
		return fmt.Errorf("could not update blog: %v", err)
	}

	return nil
}

func DeleteBlogByID(blogID string) error {
	query := `DELETE FROM blogs WHERE id = ?`
	_, err := config.DB.Exec(query, blogID)
	if err != nil {
		return fmt.Errorf("could not delete blog: %v", err)
	}
	return nil
}
