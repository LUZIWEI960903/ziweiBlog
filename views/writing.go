package views

import (
	"net/http"
	"ziweiBlog/comment"
	"ziweiBlog/service"
)

func (a *HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {
	writing := comment.Template.Writing

	wr, err := service.Writing()
	if err != nil {
		writing.WriteDefaultError(w, err)
		return
	}

	writing.WriteData(w, wr)
}
