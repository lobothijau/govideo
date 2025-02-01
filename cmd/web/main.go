package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /video", videoView)
	mux.HandleFunc("POST /video", videoCreate)
	mux.HandleFunc("GET /video/{id}", videoDetail)

	log.Println("Starting server on http://localhost:4000")

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
