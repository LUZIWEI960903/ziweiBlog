package server

import (
	"log"
	"net/http"
	"ziweiBlog/router"
)

var App = MyServer{}

type MyServer struct {
}

func (*MyServer) Start(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}

	// router
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
