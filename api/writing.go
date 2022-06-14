package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
	"ziweiBlog/comment"
	"ziweiBlog/models"
	"ziweiBlog/service"
	"ziweiBlog/utils"
)

func (a *Api) Writing(w http.ResponseWriter, r *http.Request) {
	// is user login
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		log.Println("Writing [utils.ParseToken] error:", err)
		comment.Error(w, errors.New("Login expired"))
		return
	}

	method := r.Method
	switch method {
	case http.MethodPost:
		// POST
		params := comment.GetRequestJsonParams(r)
		categoryIdStr := params["categoryId"].(string)
		categoryId, err := strconv.Atoi(categoryIdStr)
		if err != nil {
			log.Println("Writing [strconv.Atoi] error:", err)
			comment.Error(w, err)
			return
		}
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		post := &models.Post{
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     claim.Uid,
			ViewCount:  0,
			Type:       int(postType),
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		comment.Success(w, post)
	case http.MethodPut:
		// update
	}

}
