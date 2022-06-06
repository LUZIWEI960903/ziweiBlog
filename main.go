package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "ziwei blog"
	indexData.Desc = "This is ziwei blog"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
