package service

import (
	"ziweiBlog/config"
	"ziweiBlog/dao"
	"ziweiBlog/models"
)

func GetAllIndexInfo() (*models.HomeRes, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go blog",
			Content:      "content",
			UserName:     "ziwei",
			ViewCount:    123,
			CreateAt:     "2022-06-06",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeRes{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}
	return hr
}
