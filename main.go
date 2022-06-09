package main

import (
	"log"
	"net/http"
	"ziweiBlog/router"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// router
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
