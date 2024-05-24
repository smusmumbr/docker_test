package main

import (
	"github.com/gin-gonic/gin"

	"smusmumbr.io/searcher/internal/api"
	"smusmumbr.io/searcher/internal/config"
)

var server *api.Server

func init() {
	server = api.NewServer()
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/sentences", server.SearchSentences)
		v1.POST("/sentences/upload", server.AddSentences)
	}
	router.Run(config.ServerURL())
}
