package main

import (
	"log"
	"net/http"

	"github.com/saravanastar/url-shortner/internal/router"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	router, err := router.NewRedirectRouter()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", router.HandleRedirects(mux)))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
