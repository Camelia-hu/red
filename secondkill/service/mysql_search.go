package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func MysqlSearch(c *gin.Context) {
	name := c.PostForm("name")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			lock(func() {
				//查询数据库里面是否存在该id
				p := Product{}
				result := DB.Where("name = ?", name).First(&p)
				if result.Error != nil {
					if errors.Is(gorm.ErrRecordNotFound, result.Error) {
						fmt.Println("Record not found")
						c.JSON(200, "数据库中不存在该数据")
						SetNull(c, name)
					} else {
						fmt.Println("Error:", result.Error)
					}
					return
				}

				fmt.Println("Record found:", p)
				c.JSON(200, "数据库中找到了该数据")
				insert(c, name)
			}, i)
		}()
	}
}

func lock(myfunc func(), i int) {
	//lock
	key := "key"
	uuid := strconv.Itoa(i)
	lockSuccess, err := Rdb.SetNX(ctx, key, uuid, time.Second*3).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail", err)
		return
	} else {
		fmt.Println("get lock success")
	}
	//run func
	myfunc()
	//unlock
	value, _ := Rdb.Get(ctx, key).Result()
	if value == uuid { //compare value,if equal then del
		_, err := Rdb.Del(ctx, key).Result()
		if err != nil {
			fmt.Println("unlock fail")
		} else {
			fmt.Println("unlock success")
		}
	}
}
