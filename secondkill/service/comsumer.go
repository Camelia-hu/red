package service

import (
	"github.com/gin-gonic/gin"
	"secondkill/comsumer"
)

func Comsumer(c *gin.Context) {
	//如果判定成功
	comsumer.Comsumekafka()
	c.JSON(200, "receive success!")
	useLuaHash(c)
}
