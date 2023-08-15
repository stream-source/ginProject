package client

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Controller(engin *gin.Engine) {
	co := engin.Group("/cookie")
	{
		co.GET("/get-cookie", SetXyCookie)
	}
}

func SetXyCookie(c *gin.Context) {
	cookie, err := c.Cookie("xy-cookie")
	log.Printf("获取cookie：%s", cookie)
	if err != nil {
		cookieKey := "xy-cookie"
		cookieValue := "cookie-test"
		//todo 设置cookie访问域名
		c.SetCookie(cookieKey, cookieValue, 3600, "/", "localhost", false, true)
	}

}
