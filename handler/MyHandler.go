package handler

import (
	"ginProject/exception"
	"ginProject/result"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// TokenHandler gin中间件，类似Java拦截器AOP
func TokenHandler() gin.HandlerFunc {

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

// Recover 定义统一异常处理
func Recover(context *gin.Context) {
	//url := context.Request.URL.String()
	defer func() {
		if err := recover(); err != nil {
			// 异常日志
			log.Printf("异常信息: %v\n", err)
			// 打印错误堆栈信息
			//debug.PrintStack()
			//判断异常类型
			switch err.(type) {
			// 返回统一的Json风格
			case exception.BaseError:
				baseError := err.(exception.BaseError).GetBaseError()
				context.JSON(http.StatusBadRequest, result.ErrorResponse(http.StatusBadRequest, &baseError))
			}
			//终止后续操作
			context.Abort()
		}
	}()
	//继续操作
	context.Next()
}
