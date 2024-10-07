package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

const (
	REDIS_URL = "localhost:6379"
)

func main() {
	// redis client
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_URL,
		Password: "password",
		DB:       0,
	})

	err := client.Publish(ctx, "STATE", "A1").Err()
	if err != nil {
		log.Println("Error publishing message:", err)
	}
}
