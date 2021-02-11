package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redis-initiation/pkg/infrastructure"
	"redis-initiation/pkg/model"

	"github.com/gin-gonic/gin"
)

func Receive() gin.HandlerFunc {
	return func(context *gin.Context) {
		redis := infrastructure.NewRedis()
		defer redis.CloseRedis()

		key := context.Query("key")

		fmt.Println("the key is", key)

		responseInformation := model.User{}

		if payload, err := redis.Get(key); err != nil {
			fmt.Println("Failed to get data from Redis. :", err)
		} else {
			if err := json.Unmarshal(payload, &responseInformation); err != nil {
				fmt.Println("Could not Unmarshal the retrieved json. :", err)
			}
			context.JSON(http.StatusOK, responseInformation)
		}
	}
}
