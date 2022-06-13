package router

import (
	"net/http"
	"ziweiBlog/api"
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
	// 登录页面
	http.HandleFunc("/login", views.HTML.Login)
	// 登录接口
	http.HandleFunc("/api/v1/login", api.API.Login)

	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
