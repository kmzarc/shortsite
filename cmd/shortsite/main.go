package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kavehmz/short"
)

func main() {
	site := short.Site{Host: "https://short.kaveh.me/", RedisURL: os.Getenv("REDIS_URL")}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", site.Redirect)
	http.HandleFunc("/post", site.Post)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
