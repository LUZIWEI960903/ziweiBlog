package api

import (
	"net/http"
	"ziweiBlog/comment"
	"ziweiBlog/service"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {

	params := comment.GetRequestJsonParams(r)
	username := params["username"].(string)
	passwd := params["passwd"].(string)
	data, err := service.Login(username, passwd)
	if err != nil {
		comment.Error(w, err)
		return
	}
	comment.Success(w, data)
}
