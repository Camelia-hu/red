package Operate

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func Addscore(c *redis.Client, name string, addscore float64) {
	ctx := context.Background()
	err := c.ZIncrBy(ctx, "score", addscore, name).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("addscore success....")
}
