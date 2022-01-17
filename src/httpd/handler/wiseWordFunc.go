package handler

import (
	"net/http"

	wisesaying "github.com/daikimare/Kakugen/wiseSaying"
	"github.com/gin-gonic/gin"
)

func WordsGet(wisewords *wisesaying.Wisewordes) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := wisewords.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

type WordPostRequest struct {
	Id     int    `json:"id"`
	Word   string `json:"word"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func WordPost(post *wisesaying.Wisewordes) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := WordPostRequest{}
		c.Bind(&requestBody)

		item := wisesaying.Item{
			Id:     requestBody.Id,
			Word:   requestBody.Word,
			Name:   requestBody.Name,
			Detail: requestBody.Detail,
		}
		post.Add(item)

		c.Status(http.StatusNoContent)
	}
}
