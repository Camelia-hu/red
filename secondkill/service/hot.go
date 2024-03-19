package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type Product struct {
	name   string
	id     int64
	stock  int64
	status int8
}

func Hot(c *gin.Context) {
	var products []Product
	DB.Find(&products)

	// 将数据存入 Redis 缓存
	ctx := context.Background()
	for _, p := range products {
		err := Rdb.HSet(ctx, p.name, map[string]interface{}{"id": p.id, "stock": p.stock, "status": p.status}, 24*time.Hour).Err() // 设置数据过期时间为 24 小时
		if err != nil {
			panic(err)
		}
	}
}
