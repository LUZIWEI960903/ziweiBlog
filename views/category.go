package views

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"ziweiBlog/comment"
	"ziweiBlog/service"
)

func (a *HTMLApi) Category(w http.ResponseWriter, r *http.Request) {

	categoryTemplate := comment.Template.Category

	path := r.URL.Path
	categoryIdStr := strings.TrimPrefix(path, "/c/")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("router is not match"))
		return
	}

	if err := r.ParseForm(); err != nil {
		log.Println("Category [r.ParseForm()] error:", err)
		categoryTemplate.WriteDefaultError(w, err)
		return
	}

	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	pageSize := 10
	categoryRes, err := service.GetPostsByCategoryId(page, pageSize, categoryId)
	if err != nil {
		log.Println("Category [service.GetPostsByCategoryId] error:", err)
		categoryTemplate.WriteDefaultError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryRes)
}
