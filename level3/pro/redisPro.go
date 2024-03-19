package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
	"time"
)

var redisclient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

var cnt int64
var key = "i love you"
var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			lock(func() {
				cnt++
				fmt.Printf("after incr is %d\n", cnt)
			}, i)
		}()
	}
	wg.Wait()
	fmt.Printf("cnt = %d\n", cnt)
}

func lock(myfunc func(), i int) {
	//lock
	uuid := strconv.Itoa(i)
	lockSuccess, err := redisclient.SetNX(key, uuid, time.Second*3).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail", err)
		return
	} else {
		fmt.Println("get lock success")
	}
	//run func
	myfunc()
	//unlock
	value, _ := redisclient.Get(key).Result()
	if value == uuid { //compare value,if equal then del
		_, err := redisclient.Del(key).Result()
		if err != nil {
			fmt.Println("unlock fail")
		} else {
			fmt.Println("unlock success")
		}
	}
}
