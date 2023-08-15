package session

import (
	"ginProject/result"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Controller(engine *gin.Engine) {
	//定义session加密
	secret := []byte("secret")
	store := cookie.NewStore(secret)
	//绑定session中间件
	engine.Use(sessions.Sessions("xySession", store))
	session := engine.Group("/session")
	{
		session.GET("/get-session", SetXySession)
		session.GET("/global-session", globalSessionId)
	}
}

func SetXySession(c *gin.Context) {
	//初始化session对象
	session := sessions.Default(c)
	//key\value:随意指定，无协议约束
	sessionKey := "xy-session"
	sessionValue := "xySession"
	//如果浏览器第一次发送请求，则提示未授权403，否则授权成功200
	user := session.Get(sessionKey)
	log.Printf("get session: %s", user)
	if user == "" || user == nil {
		//设置session
		session.Set("session", sessionValue)
		session.Save()
		//第一次未携带session，禁止访问
		panic(result.FORBIDDEN)
	}
	if user != sessionValue {
		//携带session，但非法
		panic(result.UNAUTHORIZED)
	} else {
		c.JSON(http.StatusOK, result.OK)
	}

}

func globalSessionId(c *gin.Context) {
	request := c.Request
	responseWriter := c.Writer
	session := GlobalSessions.SessionStart(responseWriter, request)
	id := session.SessionID()
	log.Printf("session id:%s", id)
}
