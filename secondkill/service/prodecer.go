package service

import (
	"github.com/gin-gonic/gin"
	"secondkill/producer"
)

func Producer(c *gin.Context) {
	producer.ProduceKafka()
	c.JSON(200, "produce success!")
}
