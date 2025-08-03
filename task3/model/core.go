package model

import (
	"fmt"

	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var SqlxDb *sqlx.DB
var err error
var err2 error
var err3 error

var Blogdb *gorm.DB

func init() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	dsn2 := "root:123456@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=True"
	SqlxDb, err2 = sqlx.Connect("mysql", dsn2)
	if err2 != nil {
		log.Fatalf("无法连接数据库: %v", err2)
	}
	// defer SqlxDb.Close()

	dsn3 := "root:123456@tcp(127.0.0.1:3306)/blogdb?charset=utf8mb4&parseTime=True&loc=Local"
	Blogdb, err3 = gorm.Open(mysql.Open(dsn3), &gorm.Config{})
	if err3 != nil {
		log.Fatalf("无法连接到数据库: %v", err3)
	}
	fmt.Println("数据库连接成功")

	// 2. 自动迁移：根据模型创建表
	err3 = Blogdb.AutoMigrate(&User{}, &Posts{}, &Comments{})
	if err3 != nil {
		log.Fatalf("自动迁移失败: %v", err3)
	}
	fmt.Println("数据库表创建成功（或已是最新）")

}
