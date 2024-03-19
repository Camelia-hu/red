package main

import (
	"awesomeProject5/Ad"
	"awesomeProject5/Init"
	"awesomeProject5/Operate"
)

func main() {
	c := Init.Init()            //连接redis
	Ad.Add(c, 100, "测试用例1") //向Zset中添加新元素
	//如果想给某人加分的话调用Addscore函数
	//Operate.Addscore(c,"name")
	Operate.Show(c)
	Ad.Delete(c, "测试用例1")
	Operate.Show(c)

}
