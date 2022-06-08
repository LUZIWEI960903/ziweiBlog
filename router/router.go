package router

import (
	"net/http"
	"ziweiBlog/views"
)

func Router() {
	/*
		页面：view.
		数据：api.
		静态资源：
	*/
	http.HandleFunc("/", views.HTML.Index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
