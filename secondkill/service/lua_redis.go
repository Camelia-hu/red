package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

var script string = `
		local value = redis.call("HGet", name, stock)
		print("当前值为 " .. value);
		if( value - dec >= 0 ) then
			local leftStock = redis.call("DecrBy" , stock,dec)
   			print("剩余值为" .. leftStock );
			return leftStock
		else
			print("数量不够，无法扣减");
			return stock - dec
		end
		return -1
	`
var luaHash string

func luainit() {
	luaHash, _ = Rdb.ScriptLoad(ctx, script).Result() //加载lua脚本到redis里面并且返回一个哈希值sha1
}

func useLuaHash(c *gin.Context) {
	name := c.PostForm("name")
	decrease := c.PostForm("dec")
	dec, err := strconv.Atoi(decrease)
	if err != nil {
		panic(err)
	}
	luainit()
	n, err1 := Rdb.EvalSha(ctx, luaHash, []string{name}, dec).Result()
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("结果", n, err)
}
