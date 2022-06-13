package dao

import (
	"log"
	"ziweiBlog/models"
)

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post limit ?,?", pageSize*(page-1), pageSize)
	if err != nil {
		log.Println("GetPostPage [DB.Query] error:", err)
		return nil, err
	}

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Slug,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetPostPage [rows.Scan] error:", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostCount() int {
	row := DB.QueryRow("select count(1) from blog_post")
	var total int
	_ = row.Scan(&total)
	return total
}

func GetPostPageByCategoryId(page, pageSize, categoryId int) ([]models.Post, error) {
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", categoryId, pageSize*(page-1), pageSize)
	if err != nil {
		log.Println("GetPostPageByCategoryId [DB.Query] error:", err)
		return nil, err
	}

	var posts []models.Post

	for rows.Next() {
		var post models.Post
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Slug,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("GetPostPageByCategoryId [rows.Scan] error:", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetPostCountByCategoryId(categoryId int) int {
	row := DB.QueryRow("select count(1) from blog_post where category_id = ?", categoryId)
	var total int
	_ = row.Scan(&total)
	return total
}

func GetPostById(postId int) (*models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid = ?", postId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var post = new(models.Post)
	err := row.Scan(&post.Pid, &post.Title, &post.Slug, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId, &post.ViewCount, &post.Type, &post.CreateAt, &post.UpdateAt)
	if err != nil {
		return nil, err
	}
	return post, nil
}
