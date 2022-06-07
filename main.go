package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
	"ziweiBlog/config"
	"ziweiBlog/models"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func IsODD(num int) bool {
	return num&1 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(format string) string {
	return time.Now().Format(format)
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	path := config.Cfg.System.CurrentDir
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println(err)
	}
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go blog",
			Content:      "content",
			UserName:     "ziwei",
			ViewCount:    123,
			CreateAt:     "2022-06-06",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeRes{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     posts,
		Total:     1,
		Page:      1,
		Pages:     []int{1},
		PageEnd:   true,
	}
	t.Execute(w, hr)
}

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
