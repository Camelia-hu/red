package Ad

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func Delete(c *redis.Client, name string) {
	ctx := context.Background()
	_, err := c.ZRem(ctx, "score", name).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("delete success.....")
}
