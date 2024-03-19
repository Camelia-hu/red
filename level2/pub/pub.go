package main

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	err := c.Publish(ctx, "channel", "massage").Err()
	if err != nil {
		panic(err)
	}
}
