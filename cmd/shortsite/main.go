package main

import (
	"log"
	"net"
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
		log.Fatal("REDISURL must be set")
	}

	site := short.Site{Host: os.Getenv("HOST"), RedisURL: os.Getenv("REDISURL")}

	http.HandleFunc("/", site.Redirect)
	http.HandleFunc("/post", site.Post)

	// If pool is full, connections will wait.
	// This is not a good pattern for high scale sites.
	// This only helps if http connection as a resource is cheaper
	// than underlying resources like db connetion,...
	maxServingClients := 2
	maxClientsPool := make(chan bool, maxServingClients)

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: nil,
		ConnState: func(conn net.Conn, state http.ConnState) {
			switch state {
			case http.StateNew:
				maxClientsPool <- true
			case http.StateClosed, http.StateHijacked:
				<-maxClientsPool

			}
		},
	}
	log.Fatal(server.ListenAndServe())
}
