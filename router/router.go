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
	// blog首页
	http.HandleFunc("/", views.HTML.Index)
	// 根据category显示文章列表的页面
	http.HandleFunc("/c/", views.HTML.Category)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
