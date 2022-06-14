package service

import (
	"ziweiBlog/config"
	"ziweiBlog/dao"
	"ziweiBlog/models"
)

func Pigeonhole() *models.PigeonholeRes {

	posts, _ := dao.GetAllPost()

	var lines = map[string][]models.Post{}
	for _, post := range posts {
		month := post.CreateAt.Format("2006-01")
		lines[month] = append(lines[month], post)
	}

	category, _ := dao.GetAllCategory()
	return &models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    category,
		Lines:        lines,
	}
}
