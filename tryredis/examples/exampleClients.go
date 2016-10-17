package examples

import (
	"gopkg.in/redis.v5"
	"fmt"
	"sync"
)

type RedisWriter interface {
	WriteToRedis()
}

type RedisReader interface {
	ReadFromRedis()
}

// made single object for a redis client
var instance *redis.Client
var once sync.Once

func GetRedisClient() *redis.Client {
	once.Do(func() {
		instance = ExampleNewClient()
	})
	return instance
}


func ExampleNewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	if pong == "PONG" {
		return client
	}

	return nil
}

