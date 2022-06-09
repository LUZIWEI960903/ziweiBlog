package views

import (
	"net/http"
	"ziweiBlog/comment"
	"ziweiBlog/config"
	"ziweiBlog/models"
)

func init() {
	// load template
	comment.LoadTemplate()
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {

	index := comment.Template.Index

	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
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

	index.WriteData(w, hr)

}
