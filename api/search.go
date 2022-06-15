package api

import (
	"net/http"
	"ziweiBlog/comment"
	"ziweiBlog/service"
)

func (a *Api) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	val := query.Get("val")

	sr := service.Search(val)
	comment.Success(w, sr)
}
