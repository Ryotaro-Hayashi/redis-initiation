package infrastructure

import (
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	connection redis.Conn
}

func NewRedis() *Redis {
	const ipPort = "redis:6379"
	// ネットワークとアドレスを指定してRedisサーバーに接続
	conn, err := redis.Dial("tcp", ipPort)
	if err != nil {
		panic(err)
	}
	r := &Redis{
		connection: conn,
	}

	return r
}

// 接続解除
func (r *Redis) CloseRedis() {
	r.connection.Close()
}

// redisにデータ送信
func (r *Redis) Set(key string, payload []byte) error {
	// キーが存在していたら更新処理
	if r.keyExist(key) == true {
		fmt.Println("Delete the key because it was already registered in redis.")
		fmt.Println("Update an existing key.")

		r.update(key, payload)
	} else { // データ保存
		if _, err := r.connection.Do("SET", key, payload); err != nil {
			fmt.Println("infrastructure/database/Set() : ", err)
			os.Exit(1)
			return err
		}
	}
	return nil
}

// キーの重複確認
func (r *Redis) keyExist(key string) bool {
	result, err := redis.Bool(r.connection.Do("EXISTS", key))
	if err != nil {
		fmt.Println("infrastructure/database/keyExist() : ", err)
	}

	return result
}

// データ更新
func (r *Redis) update(key string, payload []byte) {
	// 返り値は古いvalueを含むものなので使用しない
	_, err := r.connection.Do("GETSET", key, payload)
	if err != nil {
		fmt.Println("infrastructure/database/update() : ", err)
	}
}

// データ取得
func (r *Redis) Get(key string) ([]byte, error) {
	// Doメソッドの返り値はinterface{}型なので[]byteに変換
	payload, err := redis.Bytes(r.connection.Do("GET", key))
	if err != nil {
		fmt.Println("infrastructure/database/Set() : ", err)
		return payload, err
	}

	return payload, err
}
