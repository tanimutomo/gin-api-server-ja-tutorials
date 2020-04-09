package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tanimutomo/go-samples/gorilla-jwt/auth"
)

type post struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
	URL   string `json:"url"`
}

func main() {
	r := mux.NewRouter()
	r.Handle("/public", public)
	r.Handle("/private", auth.JwtMiddleware.Handler(private))
	r.Handle("/auth", auth.GetTokenHandler)

	// Launch server
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", nil)
	}
}

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "gorilla-jwt",
		Tag:   "gorilla",
		URL:   "https://qiita.com/po3rin/items/740445d21487dfcb5d9f",
	}
	json.NewEncoder(w).Encode(post)
})

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "From argparse to hydra",
		Tag:   "Python",
		URL:   "https://qiita.com/tanimutomo/items/3c09cb34e47bed71e5f1",
	}
	json.NewEncoder(w).Encode(post)
})
