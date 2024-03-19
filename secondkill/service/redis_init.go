package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var Pong string
var Rdb *redis.Client

func Redisinit() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.16.147.128:6379",
		Password: "",
		DB:       0,
	})
	Rdb = rdb
	pong, err := rdb.Ping(ctx).Result()
	Pong = pong
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	fmt.Println("成功连接redis")
}
