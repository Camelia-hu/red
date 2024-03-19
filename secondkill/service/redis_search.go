package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Redis_search(c *gin.Context) {
	name := c.Param("name")
	result, err := Rdb.HExists(ctx, name, "id").Result()
	if err != nil {
		panic(err)
	}
	if result == false {
		fmt.Println("缓存中不存在该id")
		c.JSON(200, "缓存中不存在该id，可以去数据库进一步分布式锁查询")
	} else {
		fmt.Println("缓存中存在该id")
		c.JSON(200, "缓存中存在该id，可以去使用lua脚本预扣数据库库存")
	}
}
