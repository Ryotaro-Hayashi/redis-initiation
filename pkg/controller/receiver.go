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
		r := infrastructure.NewRedis()
		defer r.CloseRedis()

		// クエリパラメーターからキーを取得
		key := context.Query("key")

		responseInformation := model.User{}

		// データ取得（payloadは[]byte型）
		if payload, err := r.Get(key); err != nil {
			fmt.Println("Failed to get data from Redis. :", err)
		} else {
			// []byte型を構造体に変換
			if err := json.Unmarshal(payload, &responseInformation); err != nil {
				fmt.Println("Could not Unmarshal the retrieved json. :", err)
			}
			// context.JSONの第2引数はinterface{}型
			context.JSON(http.StatusOK, responseInformation)
		}
	}
}
