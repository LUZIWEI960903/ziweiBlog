package dao

import (
	"log"
	"ziweiBlog/models"
)

func GetAllPost() ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post")
	if err != nil {
		log.Println("GetAllPost [DB.Query] error:", err)
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid, &post.Title, &post.Slug, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId, &post.ViewCount, &post.Type, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			log.Println("GetAllPost [rows.Scan] error:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
