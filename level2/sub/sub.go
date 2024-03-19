package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pubsub := c.Subscribe(ctx,"channel")
	_,err := pubsub.ReceiveMessage(ctx)
	if err != nil {
		panic(err)
	}
	ch := pubsub.Channel()
	for msg := range ch{
		fmt.Println(msg.Payload,msg.Channel)
	}
}
