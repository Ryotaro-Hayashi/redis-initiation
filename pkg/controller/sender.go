package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"redis-initiation/pkg/infrastructure"
	"redis-initiation/pkg/model"

	"github.com/gin-gonic/gin"
)

func Send() gin.HandlerFunc {
	return func(context *gin.Context) {
		redis := infrastructure.NewRedis()
		defer redis.CloseRedis()

		requestInformation := model.User{}

		err := context.Bind(&requestInformation)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Status": "BadRequest"})
		}

		key := requestInformation.ID + ":" + requestInformation.Name

		fmt.Println("the request key is", key)
		fmt.Println("the request is", requestInformation)

		payload, err := json.Marshal(requestInformation)
		if err != nil {
			fmt.Println("JSON Marshal Error :", err)
			return
		}

		if err = redis.Set(key, payload); err != nil {
			fmt.Println("Failed to store data in Redis. ", err)
		} else {
			fmt.Println("the key is", key)
			fmt.Printf("the request is %s\n", payload)
			context.JSON(http.StatusOK, gin.H{"Status": "Successfully added to redis"})
		}
	}
}
