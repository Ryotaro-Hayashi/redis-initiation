package main

import (
	"redis-initiation/pkg/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// hello world
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hello, World!!",
		})
	})

	router.POST("/sender", controller.Send())

	router.GET("/receiver", controller.Receive())

	router.Run(":8080")
}
