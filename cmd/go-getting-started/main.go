package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kavehmz/short"
)

func main() {
	site := short.Site{Host: "https://short.kaveh.me/"}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", site.Redirect)
	http.HandleFunc("/post", site.Post)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// router := gin.New()
	// router.Use(gin.Logger())
	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")
	//
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })
	//
	// router.Run(":" + port)
}
