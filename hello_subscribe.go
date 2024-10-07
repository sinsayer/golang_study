package main

import (
	"context"
	"fmt"
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
		Password: "",
		DB:       0,
	})

	// get messages from redis
	pubsub := client.Subscribe(ctx, "room-100013")

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(msg.Channel, msg.Payload)
	}
}
