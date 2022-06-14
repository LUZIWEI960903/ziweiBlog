package views

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"ziweiBlog/comment"
	"ziweiBlog/service"
)

func init() {
	// load template
	comment.LoadTemplate()
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {

	index := comment.Template.Index

	if err := r.ParseForm(); err != nil {
		log.Println("Index [r.ParseForm()] error:", err)
		index.WriteDefaultError(w, err)
		return
	}

	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	pageSize := 10

	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")

	hr, err := service.GetAllIndexInfo(slug, page, pageSize)
	if err != nil {
		log.Println("Index [service.GetAllIndexInfo] error:", err)
		index.WriteDefaultError(w, err)
		return
	}

	index.WriteData(w, hr)

}
