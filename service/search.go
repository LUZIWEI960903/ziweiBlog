package service

import (
	"ziweiBlog/dao"
	"ziweiBlog/models"
)

func Search(val string) []models.SearchRes {
	posts, _ := dao.Search(val)

	var srs []models.SearchRes
	for _, post := range posts {
		sr := models.SearchRes{
			Pid:   post.Pid,
			Title: post.Title,
		}
		srs = append(srs, sr)
	}
	return srs
}
