package views

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"ziweiBlog/comment"
	"ziweiBlog/service"
)

func (*HTMLApi) Detail(w http.ResponseWriter, r *http.Request) {
	detail := comment.Template.Detail

	path := r.URL.Path
	postIdStr := strings.TrimPrefix(path, "/p/")
	postIdStr = strings.Trim(postIdStr, ".html")
	postId, err := strconv.Atoi(postIdStr)
	if err != nil {
		detail.WriteError(w, errors.New("router is not match"))
		return
	}

	data, err := service.GetPostDetail(postId)
	if err != nil {
		detail.WriteError(w, err)
		return
	}
	detail.WriteData(w, data)
}
