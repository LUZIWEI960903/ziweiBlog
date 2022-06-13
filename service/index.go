package service

import (
	"html/template"
	"math"
	"ziweiBlog/config"
	"ziweiBlog/dao"
	"ziweiBlog/models"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeRes, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	posts, err := dao.GetPostPage(page, pageSize)
	if err != nil {
		return nil, err
	}

	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:100]
		}
		postMore := models.PostMore{
			Pid:          post.Pid,
			Title:        post.Title,
			Slug:         post.Slug,
			Content:      template.HTML(content),
			CategoryId:   post.CategoryId,
			CategoryName: categoryName,
			UserId:       post.UserId,
			UserName:     userName,
			ViewCount:    post.ViewCount,
			Type:         post.Type,
			CreateAt:     models.DateDay(post.CreateAt),
			UpdateAt:     models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}

	total := dao.GetPostCount()
	pageCount := math.Ceil(float64(total) / float64(page))
	var pages []int
	for i := 0; i < int(pageCount); i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeRes{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postMores,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != int(pageCount),
	}
	return hr, nil
}
