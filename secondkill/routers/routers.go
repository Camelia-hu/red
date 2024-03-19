package routers

import "github.com/gin-gonic/gin"

func InitRouters() {
	r := gin.Default()
	r.POST("/hot")               //预热
	r.GET("/redis_search/:name") //在缓存中搜索
	r.POST("/producer")          //若存在该id，将秒杀成功信息输入kafka
	r.POST("/consumer")          //接受kafka内消息并且使用lua脚本预扣库存
	r.GET("mysql_search")        //若不存在该id，则使用分布式锁来查询数据库
	r.POST("redis_insert")       //若数据库中存在该id则将该数据存入缓存
	r.POST("redis_null")         //若不存在则将该id在缓存里面设为空值
	r.Run()
}
