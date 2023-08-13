// 包，如果该目录下无main方法，则与目录保持一致；
// go 中的main同一个目录下仅允许存在一个，即使是不同的类中也不可以
package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Controller(engine *gin.Engine) {
	//路由分组，类似Java中的requestMapping，路由前缀
	user := engine.Group("/user")
	{
		user.GET("/get-user-info", GetUserInfo)
		//Get请求：变量位于URL路径中
		user.GET("/:id/:username", QueryByURL)
		//Get请求：变量以form表单形式提交
		user.GET("/query", QueryByParam)

		//Post请求
		user.POST("/insert", JsonPost)
		//Post请求
		user.POST("/form-insert", FormPost)
	}
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" form:"username" binding:"required"`
	Sex      int    `form:"sex"`
}

// GetUserInfo 自定义函数
func GetUserInfo(context *gin.Context) {
	log.Printf("获取用户信息，ID：")

	//初始化对象
	user := &User{
		Id:       1,
		Username: "ceshi",
	}
	context.JSON(http.StatusOK, user)
}

// QueryByURL 通过URL中的参数查询，例如：/user/2/测试
func QueryByURL(c *gin.Context) {
	//获取url中的参数,并将string转换为int类型
	var id, _ = strconv.Atoi(c.Param("id"))
	username := c.Param("username")
	log.Printf("id:%d,username:%s", id, username)
	//初始化对象
	user := &User{
		Id:       id,
		Username: username,
		Sex:      1,
	}

	c.JSON(http.StatusOK, user)
}

// QueryByParam 通过表单参数查询
func QueryByParam(c *gin.Context) {
	id1 := c.Query("id")
	id2 := c.Request.URL.Query().Get("id")
	log.Printf("%d, %d", id1, id2)

}

// JsonPost post通过json格式
func JsonPost(c *gin.Context) {
	//1、获取json数据
	//json, _ := c.GetRawData()
	//log.Println("post 请求json数据,", json)
	//// 定义map或结构体
	//var m map[string]interface{}
	//// 反序列化
	//_ = json.Unmarshal(json, &m)
	//
	//c.JSON(http.StatusOK, json)

	//2、直接将json数据与结构体绑定
	// 直接将结构体和提交的json参数作绑定
	//多次获取body会报错
	var user User
	err := c.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, user)
}

// FormPost post请求通过表单格式
func FormPost(c *gin.Context) {
	//username := c.PostForm("username")

	//user := &User{Username: username}

	//方式二
	user := &User{}
	if err := c.Bind(&user); err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, user)
}
