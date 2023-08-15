package cookie

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
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
		//todo :后期可以替换成全局sessionID
		cookieValue := RandStringBytes(6)
		//todo 设置cookie访问域名
		c.SetCookie(cookieKey, cookieValue, 3600, "/", "localhost", false, true)
	}

}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
