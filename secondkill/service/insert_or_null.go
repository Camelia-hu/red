package service

import (
	"github.com/gin-gonic/gin"
	"time"
)

func insert(c *gin.Context, name string) {
	p := Product{}
	err1 := DB.Where("name =?", name).First(&p).Error
	if err1 != nil {
		panic(err1)
	}
	err := Rdb.HSet(ctx, p.name, map[string]interface{}{"id": p.id, "stock": p.stock, "status": p.status}, 24*time.Hour).Err()
	if err != nil {
		panic(err)
	}
	c.JSON(200, "已插入")
}
func SetNull(c *gin.Context, name string) {
	err := Rdb.HSet(ctx, name, map[string]interface{}{"id": 0, "stock": 0, "status": 0}, 24*time.Hour).Err()
	if err != nil {
		panic(err)
	}
	c.JSON(200, "已设空值")
}
