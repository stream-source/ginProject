package main

import (
	"ginProject/controller/person"
	"ginProject/controller/user"
	"ginProject/database"
	"github.com/gin-gonic/gin"
)

func init() {
	database.DB = database.InitDB()
}

func main() {
	engine := setupRouter() //获取默认的路由引擎
	//用户相关路由
	user.Controller(engine)
	//人相关路由
	person.Controller(engine)
	engine.Run(":8083")
}

func setupRouter() *gin.Engine {
	engine := gin.Default() //获取默认的路由引擎
	return engine
}
