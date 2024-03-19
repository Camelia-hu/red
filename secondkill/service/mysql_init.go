package service

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Mysqlinit() {
	dsn := "root:15926653820hhh@tcp(127.0.0.1:3306)/seckill?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB = db
	if err != nil {
		fmt.Println("mysql连接失败")
		panic(err)
	}
}
