package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:root@tcp(10.8.10.18:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{})
	return db

}
