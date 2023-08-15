package main

import (
	"ginProject/controller/person"
	"ginProject/controller/user"
	"ginProject/cookie"
	"ginProject/database"
	"ginProject/handler"
	"ginProject/session"
	"github.com/gin-gonic/gin"
)

func init() {
	database.DB = database.InitDB()
}

func main() {
	engine := setupRouter() //获取默认的路由引擎
	//使用自定义中间件鉴权
	engine.Use(handler.Recover)
	//用户相关路由
	user.Controller(engine)
	//人相关路由
	person.Controller(engine)
	cookie.Controller(engine)  //客户端访问cookie
	session.Controller(engine) //访问session
	engine.Run(":8083")
}

func setupRouter() *gin.Engine {
	engine := gin.Default() //获取默认的路由引擎
	return engine
}
