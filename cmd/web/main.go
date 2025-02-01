package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /video", videoView)
	mux.HandleFunc("POST /video", videoCreate)
	mux.HandleFunc("GET /video/{id}", videoDetail)

	log.Printf("Starting server on http://localhost:%s", *addr)

	port := fmt.Sprintf(":%s", *addr)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
