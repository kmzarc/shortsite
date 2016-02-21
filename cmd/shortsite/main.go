package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kavehmz/short"
)

func main() {
	if os.Getenv("PORT") == "" {
		log.Fatal("PORT must be set")
	}
	if os.Getenv("HOST") == "" {
		log.Fatal("HOST must be set")
	}
	if os.Getenv("REDISURL") == "" {
		log.Fatal("REDIS_URL must be set")
	}

	site := short.Site{Host: os.Getenv("HOST"), RedisURL: os.Getenv("REDISURL")}

	http.HandleFunc("/", site.Redirect)
	http.HandleFunc("/post", site.Post)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}
