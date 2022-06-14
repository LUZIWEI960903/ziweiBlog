package service

import (
	"ziweiBlog/config"
	"ziweiBlog/dao"
	"ziweiBlog/models"
)

func Writing() (*models.WritingRes, error) {
	wr := new(models.WritingRes)
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL

	category, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	wr.Categorys = category
	return wr, nil
}

func SavePost(post *models.Post) {
	dao.SavePost(post)
}

func GetPost(postId int) (*models.Post, error) {
	post, err := dao.GetPostById(postId)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}
