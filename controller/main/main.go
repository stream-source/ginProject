package main

import (
	"ginProject/controller/user"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := setupRouter() //获取默认的路由引擎
	//db := database.InitDB()

	user.Controller(engine)
	engine.Run(":8083")
}

func setupRouter() *gin.Engine {
	engine := gin.Default() //获取默认的路由引擎
	return engine
}
