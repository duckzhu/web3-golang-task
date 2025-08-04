package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
var DB *gorm.DB
var err error

func init() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败")
		fmt.Println(err)
	}

}
