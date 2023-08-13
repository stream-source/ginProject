package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	/**
	使用gin框架启动程序
	*/
	//engine.GET("/index", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"code": "01",
	//		"name": "gin start success",
	//	})
	//})
	engine := setupRouter() //获取默认的路由引擎
	engine.GET("/index", Index)
	engine.Run(":8083")
}

func setupRouter() *gin.Engine {
	engine := gin.Default() //获取默认的路由引擎
	return engine
}

// 首页
func Index(c *gin.Context) {
	log.Println("开始执行首页方法")
	c.JSON(http.StatusOK, gin.H{
		"code": "01",
		"name": "gin start success",
	})
}
