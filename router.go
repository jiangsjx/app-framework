package main

import (
	"framework/handler"

	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/hello", handler.Hello)
	}

	return router
}
