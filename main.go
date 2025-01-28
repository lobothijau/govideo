package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func videoView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Video View!"))
}

func videoDetail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Video Detail!"))
}

func videoCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Video Create!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /video", videoView)
	mux.HandleFunc("POST /video", videoCreate)
	mux.HandleFunc("GET /video/{id}", videoDetail)

	log.Println("Starting server on http://localhost:4000")

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
