package Operate

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func Show(c *redis.Client) {
	ctx := context.Background()
	result, err := c.ZRevRange(ctx, "score", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	for _, results := range result {
		fmt.Println(results)
	}
	fmt.Println("show success....")
}
