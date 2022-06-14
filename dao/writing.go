package dao

import (
	"log"
	"ziweiBlog/models"
)

func SavePost(post *models.Post) {
	_, err := DB.Exec("insert into blog_post (title,slug,content,markdown,category_id,user_id,view_count,type,create_at,update_at) values (?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Slug,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.CreateAt,
		post.UpdateAt,
	)

	if err != nil {
		log.Println("SavePost error:", err)
	}

	//result.LastInsertId()
}
