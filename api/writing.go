package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		params := comment.GetRequestJsonParams(r)
		categoryId := params["categoryId"].(float64)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		postId := params["pid"].(float64)

		post := &models.Post{
			Pid:        int(postId),
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: int(categoryId),
			UserId:     claim.Uid,
			ViewCount:  0,
			Type:       int(postType),
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.UpdatePost(post)
		comment.Success(w, post)
	}

}

func (a *Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	postIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		log.Println("GetPost [strconv.Atoi] error:", err)
		comment.Error(w, errors.New("Server error, please contact the administrator~"))
		return
	}

	post, err := service.GetPost(postId)
	if err != nil {
		log.Println("GetPost [service.GetPost] error:", err)
		comment.Error(w, err)
		return
	}

	comment.Success(w, post)
}
