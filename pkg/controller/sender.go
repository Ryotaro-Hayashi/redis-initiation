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
		r := infrastructure.NewRedis()
		defer r.CloseRedis() // 遅延処理で接続解除

		requestInformation := model.User{}

		// requestをUser型で受け取り
		err := context.Bind(&requestInformation)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"Status": "BadRequest"})
		}

		// redisのキー
		key := requestInformation.ID + ":" + requestInformation.Name

		// requestを構造体からJSONにパース（payloadは[]byte型）
		payload, err := json.Marshal(requestInformation)
		if err != nil {
			fmt.Println("JSON Marshal Error :", err)
			return
		}

		// redisにデータ送信
		if err = r.Set(key, payload); err != nil {
			fmt.Println("Failed to store data in Redis. ", err)
		} else {
			context.JSON(http.StatusOK, gin.H{"Status": "Successfully added to redis"})
		}
	}
}
