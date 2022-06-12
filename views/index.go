package views

import (
	"log"
	"net/http"
	"ziweiBlog/comment"
	"ziweiBlog/service"
)

func init() {
	// load template
	comment.LoadTemplate()
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {

	index := comment.Template.Index

	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Fatal("Index [service.GetAllIndexInfo] error:",err)
		index.
	}

	index.WriteData(w, hr)

}
