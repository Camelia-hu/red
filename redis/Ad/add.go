package Ad

import (
	"context"
	"github.com/redis/go-redis/v9"
)

// 向Zset中添加元素
func Add(c *redis.Client, score float64, name string) {
	ctx := context.Background()
	err := c.ZAdd(ctx, "score", redis.Z{Score: score, Member: name}).Err()
	if err != nil {
		panic(err)
	}
	println("add success.....")
}
