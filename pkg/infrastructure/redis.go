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
	c, err := redis.Dial("tcp", ipPort)
	if err != nil {
		panic(err)
	}
	r := &Redis{
		connection: c,
	}

	return r
}

func (r *Redis) CloseRedis() {
	r.connection.Close()
}

func (r *Redis) Set(key string, payload []byte) error {
	if r.keyExist(key) == true {
		fmt.Println("Delete the key because it was already registered in redis.")
		fmt.Println("Update an existing key.")

		r.update(key, payload)
	} else {
		// キーをRedisに追加する
		if _, err := r.connection.Do("SET", key, payload); err != nil {
			fmt.Println("infrastructure/database/Set() : ", err)
			os.Exit(1)
			return err
		}
	}
	return nil
}

func (r *Redis) keyExist(key string) bool {
	result, err := redis.Bool(r.connection.Do("EXISTS", key))
	if err != nil {
		fmt.Println("infrastructure/database/keyExist() : ", err)
	}

	return result
}

func (r *Redis) update(key string, payload []byte) {
	_, err := r.connection.Do("GETSET", key, payload)
	if err != nil {
		fmt.Println("infrastructure/database/update() : ", err)
	}
}

func (r *Redis) Get(key string) ([]byte, error) {
	payload, err := redis.Bytes(r.connection.Do("GET", key))
	if err != nil {
		fmt.Println("infrastructure/database/Set() : ", err)
		return payload, err
	}

	return payload, err
}
