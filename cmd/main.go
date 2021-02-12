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

	// データ送信
	router.POST("/sender", controller.Send())

	// データ取得
	router.GET("/receiver", controller.Receive())

	router.Run(":8080")
}
