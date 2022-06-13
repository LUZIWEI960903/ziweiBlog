package views

import (
	"net/http"
	"ziweiBlog/comment"
	"ziweiBlog/config"
)

func (a *HTMLApi) Login(w http.ResponseWriter, r *http.Request) {

	loginTemplate := comment.Template.Login

	loginTemplate.WriteData(w, config.Cfg.Viewer)
}
