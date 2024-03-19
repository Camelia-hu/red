package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

// 创建client
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
		wg.Add(1) //上锁阻塞协程
		go func() {
			lock(func() {
				cnt++
				fmt.Printf("after incr is %d\n", cnt)
			})
		}()
	}
	wg.Wait() //等待协程中wg.Done将waitgroup中的数字降为0才释放锁
	fmt.Printf("cnt = %d\n", cnt)
}

func lock(handler func()) {
	defer wg.Done() //延迟降低waitgroup里的数字

	lockSuccess, err := redisclient.SetNX(key, 1, time.Second*3).Result() //设置key键存放1，设置过期时间为三秒以防死锁
	if err != nil || lockSuccess != true {
		fmt.Println("get lock fail", err)
		return
	} else {
		fmt.Println("get lock success")
	}

	handler() //锁持有者进行操作

	//unlock
	_, err = redisclient.Del(key).Result() //使用del命令来将key键清空
	if err != nil {
		fmt.Println("unlock fail", err)
	} else {
		fmt.Println("unlock success")
	}
	//defer解锁
}

//有一个弊端，如果锁持有者处理时间超过过期时间锁会自动解锁，这个时候其他协程会来一起持有锁
//而且当锁持有者1完成操作后，锁会被打开，连同着锁持有者2也会失去这个锁的唯一权，后续协程会来一起持有锁
//所以为了防止此类事件产生，pro版加了一些其他操作
