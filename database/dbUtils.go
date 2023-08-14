package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义全局变量DB
var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(10.8.10.18:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{})
	//解决默认gorm认为结构体与表命名存在复数关系

	return db
}
