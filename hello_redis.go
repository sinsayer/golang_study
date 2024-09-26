package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Message struct {
	key  string
	data RedisData
}

type RedisData struct {
	key   string
	value string
}

var ctx = context.Background()

func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "test1", "12348765", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "test1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "test2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("test2", val2)
	}
}

func main() {
	// message := Message{key: "1234", data: RedisData{key: "name", value: "John"}}
	// fmt.Println(message)
	ExampleClient()
}
