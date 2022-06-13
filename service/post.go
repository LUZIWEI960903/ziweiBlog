package service

import (
	"html/template"
	"ziweiBlog/config"
	"ziweiBlog/dao"
	"ziweiBlog/models"
)

func GetPostDetail(postId int) (*models.PostRes, error) {
	post, err := dao.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	username := dao.GetUserNameById(post.UserId)
	var postMore = models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     username,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     models.DateDay(post.CreateAt),
		UpdateAt:     models.DateDay(post.UpdateAt),
	}

	var postRes = &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article:      postMore,
	}
	return postRes, nil
}
