package main

import (
	"fmt"
	"net/http"

	"github.com/daikimare/Kakugen/httpd/handler"
	wisesaying "github.com/daikimare/Kakugen/wiseSaying"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello")

	word := wisesaying.New()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/wiseword", handler.WordsGet(word))
	r.POST("/wiseword", handler.WordPost(word))

	r.Run()
}
