package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// HandlerToken gin中间件，类似Java拦截器AOP
func HandlerToken() gin.HandlerFunc {

	return func(context *gin.Context) {
		if auth := context.Request.Header.Get("Authorization"); auth == "nil" || auth == "" {
			//context.JSON(http.StatusForbidden, "token不能为空")
			panic("token不能为空")
		}
		log.Println("获取token成功，开始执行鉴权~")
		context.Next()
		context.JSON(http.StatusOK, "授权成功")
	}
}
