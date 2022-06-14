package views

import (
	"net/http"
	"ziweiBlog/comment"
	"ziweiBlog/service"
)

func (a *HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := comment.Template.Pigeonhole

	pr := service.Pigeonhole()

	pigeonhole.WriteData(w, pr)
}
